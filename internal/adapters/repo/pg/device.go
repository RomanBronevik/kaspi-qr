package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/adapters/db"

	"github.com/rendau/dop/dopErrs"
)

func (d *St) DeviceGet(ctx context.Context, id string) (*entities.DeviceSt, error) {
	result := &entities.DeviceSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"device"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		err = nil
	}

	return result, err
}

func (d *St) DeviceGetIdForCityId(ctx context.Context, cityId string) (string, error) {
	var result string

	err := d.DbQueryRow(ctx, `
		select d.id
		from device d
			join city c on c.org_bin = d.org_bin
		where c.id = $1
		order by d.created desc
		limit 1
	`, cityId).Scan(&result)
	if errors.Is(err, dopErrs.NoRows) {
		result, err = "", nil
	}

	return result, err
}

func (d *St) DeviceList(ctx context.Context, pars *entities.DeviceListParsSt) ([]*entities.DeviceSt, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: text[]))`)
		args["ids"] = *pars.Ids
	}
	if pars.Token != nil {
		conds = append(conds, "t.token = ${token}")
		args["token"] = *pars.Token
	}
	if pars.TradePointId != nil {
		conds = append(conds, "t.trade_point_id = ${trade_point_id}")
		args["trade_point_id"] = *pars.TradePointId
	}
	if pars.OrgBin != nil {
		conds = append(conds, "t.org_bin = ${org_bin}")
		args["org_bin"] = *pars.OrgBin
	}
	if pars.CityId != nil {
		conds = append(conds, "t.org_bin = (select org_bin from city where id = ${city_id})")
		args["city_id"] = *pars.CityId
	}

	result := make([]*entities.DeviceSt, 0)

	_, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`device t`},
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.created desc",
		},
	})

	return result, err
}

func (d *St) DeviceIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.DbQueryRow(ctx, `
		select count(*)
		from device
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) DeviceCreate(ctx context.Context, obj *entities.DeviceCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "device",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) DeviceUpdate(ctx context.Context, id string, obj *entities.DeviceCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "device",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) DeviceDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "device",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
