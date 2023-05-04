package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/adapters/db"

	"github.com/rendau/dop/dopErrs"
)

func (d *St) PaymentGet(ctx context.Context, id int64) (*entities.PaymentSt, error) {
	result := &entities.PaymentSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"payment"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		err = nil
	}

	return result, err
}

func (d *St) PaymentGetLink(ctx context.Context, id int64) (string, error) {
	var result string

	err := d.DbQueryRow(ctx, `
		select link
		from payment
		where id = $1
	`, id).Scan(&result)
	if errors.Is(err, dopErrs.NoRows) {
		return "", nil
	}

	return result, err
}

func (d *St) PaymentList(ctx context.Context, pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: text[]))`)
		args["ids"] = *pars.Ids
	}
	if pars.OrdId != nil {
		conds = append(conds, "t.ord_id = ${ord_id}")
		args["ord_id"] = *pars.OrdId
	}
	if pars.Status != nil {
		conds = append(conds, "t.status = ${status}")
		args["status"] = *pars.Status
	}
	if pars.Statuses != nil {
		conds = append(conds, `t.status in (select * from unnest(${statuses} :: text[]))`)
		args["statuses"] = *pars.Statuses
	}

	result := make([]*entities.PaymentSt, 0)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`payment t`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.created",
		},
	})

	return result, tCount, err
}

func (d *St) PaymentIdExists(ctx context.Context, id int64) (bool, error) {
	var err error
	var cnt int

	err = d.DbQueryRow(ctx, `
		select count(*)
		from payment
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) PaymentCreate(ctx context.Context, obj *entities.PaymentCUSt) (int64, error) {
	var result int64

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "payment",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) PaymentUpdate(ctx context.Context, id int64, obj *entities.PaymentCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "payment",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) PaymentDelete(ctx context.Context, id int64) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "payment",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
