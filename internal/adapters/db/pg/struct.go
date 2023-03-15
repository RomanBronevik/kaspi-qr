package pg

import (
	"context"
	"time"

	"kaspi-qr/internal/adapters/db"

	"github.com/jackc/pgx/v4"
)

// Options

type OptionsSt struct {
	Dsn               string
	Timezone          string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func (o *OptionsSt) mergeWithDefaults() {
	if o.Timezone == "" {
		o.Timezone = defaultOptions.Timezone
	}
	if o.MaxConns == 0 {
		o.MaxConns = defaultOptions.MaxConns
	}
	if o.MinConns == 0 {
		o.MinConns = defaultOptions.MinConns
	}
	if o.MaxConnLifetime == 0 {
		o.MaxConnLifetime = defaultOptions.MaxConnLifetime
	}
	if o.MaxConnIdleTime == 0 {
		o.MaxConnIdleTime = defaultOptions.MaxConnIdleTime
	}
	if o.HealthCheckPeriod == 0 {
		o.HealthCheckPeriod = defaultOptions.HealthCheckPeriod
	}
}

// TransactionWrapper

type TransactionWrapper struct {
	tx pgx.Tx
}

func (o *TransactionWrapper) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := o.tx.Exec(ctx, sql, args...)
	return err
}

func (o *TransactionWrapper) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	return o.tx.Query(ctx, sql, args...)
}

func (o *TransactionWrapper) QueryRow(ctx context.Context, sql string, args ...any) db.Row {
	return o.tx.QueryRow(ctx, sql, args...)
}

// rows

type rowsSt struct {
	pgx.Rows

	db db.DB
}

func (o rowsSt) Err() error {
	return o.db.HErr(o.Rows.Err())
}

func (o rowsSt) Scan(dest ...any) error {
	return o.db.HErr(o.Rows.Scan(dest...))
}

// row

type rowSt struct {
	pgx.Row

	db db.DB
}

func (o rowSt) Scan(dest ...any) error {
	return o.db.HErr(o.Row.Scan(dest...))
}
