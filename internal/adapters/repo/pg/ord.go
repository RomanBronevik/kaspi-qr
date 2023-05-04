package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/adapters/db"

	"github.com/rendau/dop/dopErrs"
)

func (d *St) OrdGet(ctx context.Context, id string) (*entities.OrdSt, error) {
	result := &entities.OrdSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"ord"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) OrdList(ctx context.Context, pars *entities.OrdListParsSt) ([]*entities.OrdSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: text[]))`)
		args["ids"] = *pars.Ids
	}
	if pars.SrcId != nil {
		conds = append(conds, "t.src_id = ${src_id}")
		args["src_id"] = *pars.SrcId
	}
	if pars.DeviceId != nil {
		conds = append(conds, "t.device_id = ${device_id}")
		args["device_id"] = *pars.DeviceId
	}
	if pars.PaymentId != nil {
		conds = append(conds, "t.id = (select ord_id from payment where id = ${payment_id})")
		args["payment_id"] = *pars.PaymentId
	}
	if pars.CityId != nil {
		conds = append(conds, "t.city_id = ${city_id}")
		args["city_id"] = *pars.CityId
	}
	if pars.Status != nil {
		conds = append(conds, "t.status = ${status}")
		args["status"] = *pars.Status
	}
	if pars.Platform != nil {
		conds = append(conds, "t.platform = ${platform}")
		args["platform"] = *pars.Platform
	}

	result := make([]*entities.OrdSt, 0)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`ord t`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.created",
		},
	})

	return result, tCount, err
}

func (d *St) OrdIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.DbQueryRow(ctx, `
		select count(*)
		from ord
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) OrdCreate(ctx context.Context, obj *entities.OrdCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "ord",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) OrdUpdate(ctx context.Context, id string, obj *entities.OrdCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "ord",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) OrdDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "ord",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
