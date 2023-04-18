package pg

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/adapters/logger"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib" // driver
)

type St struct {
	debug bool
	lg    logger.WarnAndError
	opts  OptionsSt

	Con *pgxpool.Pool
}

func New(debug bool, lg logger.WarnAndError, opts OptionsSt) (*St, error) {
	opts.mergeWithDefaults()

	cfg, err := pgxpool.ParseConfig(opts.Dsn)
	if err != nil {
		lg.Errorw("Fail to create config", err, "opts", opts)
		return nil, err
	}

	cfg.ConnConfig.RuntimeParams["timezone"] = opts.Timezone
	cfg.MaxConns = opts.MaxConns
	cfg.MinConns = opts.MinConns
	cfg.MaxConnLifetime = opts.MaxConnLifetime
	cfg.MaxConnIdleTime = opts.MaxConnIdleTime
	cfg.HealthCheckPeriod = opts.HealthCheckPeriod
	cfg.LazyConnect = true

	dbPool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		lg.Errorw(ErrPrefix+": Fail to connect to db", err)
		return nil, err
	}

	return &St{
		debug: debug,
		lg:    lg,
		opts:  opts,
		Con:   dbPool,
	}, nil
}

// transaction

func (d *St) RenewTransaction(ctx context.Context) (context.Context, error) {
	var err error

	err = d.commitContextTransaction(ctx)
	if err != nil {
		return ctx, err
	}

	return d.contextWithTransaction(ctx)
}

func (d *St) TransactionFn(ctx context.Context, f func(context.Context) error) error {
	var err error

	if ctx == nil {
		ctx = context.Background()
	}

	if ctx, err = d.contextWithTransaction(ctx); err != nil {
		return err
	}
	defer func() { d.rollbackContextTransaction(ctx) }()

	err = f(ctx)
	if err != nil {
		return err
	}

	return d.commitContextTransaction(ctx)
}

func (d *St) getContextTransaction(ctx context.Context) pgx.Tx {
	if v := ctx.Value(transactionCtxKey); v != nil {
		tr, ok := v.(pgx.Tx)
		if ok {
			return tr
		}
	}
	return nil
}

func (d *St) contextWithTransaction(ctx context.Context) (context.Context, error) {
	tx, err := d.Con.Begin(ctx)
	if err != nil {
		return ctx, d.HErr(err)
	}

	return context.WithValue(ctx, transactionCtxKey, tx), nil
}

func (d *St) commitContextTransaction(ctx context.Context) error {
	tx := d.getContextTransaction(ctx)
	if tx == nil {
		return nil
	}

	err := tx.Commit(ctx)
	if err != nil {
		if err != pgx.ErrTxClosed &&
			err != pgx.ErrTxCommitRollback {
			_ = tx.Rollback(ctx)

			return d.HErr(err)
		}
	}

	return nil
}

func (d *St) rollbackContextTransaction(ctx context.Context) {
	tx := d.getContextTransaction(ctx)
	if tx == nil {
		return
	}

	_ = tx.Rollback(ctx)
}

// query

func (d *St) Exec(ctx context.Context, sql string, args ...any) error {
	if tx := d.getContextTransaction(ctx); tx != nil {
		_, err := tx.Exec(ctx, sql, args...)
		return d.HErr(err)
	}

	_, err := d.Con.Exec(ctx, sql, args...)
	return d.HErr(err)
}

func (d *St) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	var err error
	var rows pgx.Rows

	if tx := d.getContextTransaction(ctx); tx != nil {
		rows, err = tx.Query(ctx, sql, args...)
	} else {
		rows, err = d.Con.Query(ctx, sql, args...)
	}

	return &rowsSt{Rows: rows, db: d}, d.HErr(err)
}

func (d *St) QueryRow(ctx context.Context, sql string, args ...any) db.Row {
	var row pgx.Row

	if tx := d.getContextTransaction(ctx); tx != nil {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = d.Con.QueryRow(ctx, sql, args...)
	}

	return &rowSt{Row: row, db: d}
}

func (d *St) queryRebindNamed(sql string, argMap map[string]any) (string, []any) {
	resultQuery := sql
	args := make([]any, 0, len(argMap))

	for k, v := range argMap {
		if strings.Contains(resultQuery, "${"+k+"}") {
			args = append(args, v)
			resultQuery = strings.ReplaceAll(resultQuery, "${"+k+"}", "$"+strconv.Itoa(len(args)))
		}
	}

	if d.debug {
		if strings.Index(resultQuery, "${") > -1 {
			for _, x := range queryParamRegexp.FindAllString(resultQuery, 1) {
				d.lg.Errorw(ErrPrefix+": missing param", nil, "param", x, "query", resultQuery)
			}
		}
	}

	return resultQuery, args
}

func (d *St) ExecM(ctx context.Context, sql string, argMap map[string]interface{}) error {
	rbSql, args := d.queryRebindNamed(sql, argMap)

	return d.Exec(ctx, rbSql, args...)
}

func (d *St) QueryM(ctx context.Context, sql string, argMap map[string]interface{}) (db.Rows, error) {
	rbSql, args := d.queryRebindNamed(sql, argMap)

	return d.Query(ctx, rbSql, args...)
}

func (d *St) QueryRowM(ctx context.Context, sql string, argMap map[string]interface{}) db.Row {
	rbSql, args := d.queryRebindNamed(sql, argMap)

	return d.QueryRow(ctx, rbSql, args...)
}

// handle error

func (d *St) HErr(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, pgx.ErrNoRows), errors.Is(err, sql.ErrNoRows):
		err = db.ErrNoRows
	default:
		d.lg.Errorw(ErrPrefix, err)
	}

	return err
}
