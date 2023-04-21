package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) DeviceGet(ctx context.Context, id string) (*entities.DeviceSt, error) {
	var result entities.DeviceSt

	err := d.db.QueryRow(ctx, `
		select
			t.id,
			t.created,
			t.token,
			t.trade_point_id,
			t.org_bin
		from device t
		where t.id = $1
	`, id).Scan(
		&result.Id,
		&result.Created,
		&result.Token,
		&result.TradePointId,
		&result.OrgBin,
	)
	if errors.Is(err, db.ErrNoRows) {
		return nil, nil
	}

	return &result, err
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

	rows, err := d.db.QueryM(ctx, `
		select
			t.id,
			t.created,
			t.token,
			t.trade_point_id,
			t.org_bin
		from device t
		`+d.tOptionalWhere(conds)+`
		order by t.created desc`,
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.DeviceSt, 0)

	for rows.Next() {
		item := &entities.DeviceSt{}

		err = rows.Scan(
			&item.Id,
			&item.Created,
			&item.Token,
			&item.TradePointId,
			&item.OrgBin,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *St) DeviceIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.db.QueryRow(ctx, `
		select count(*)
		from device
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) DeviceCreate(ctx context.Context, obj *entities.DeviceCUSt) (string, error) {
	fields := d.deviceGetCUFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId string

	err := d.db.QueryRowM(ctx, `
		insert into device (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) DeviceUpdate(ctx context.Context, id string, obj *entities.DeviceCUSt) error {
	fields := d.deviceGetCUFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.db.ExecM(ctx, `
		update device
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) deviceGetCUFields(obj *entities.DeviceCUSt) map[string]any {
	result := map[string]any{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.Token != nil {
		result["token"] = *obj.Token
	}

	if obj.TradePointId != nil {
		result["trade_point_id"] = *obj.TradePointId
	}

	if obj.OrgBin != nil {
		result["org_bin"] = *obj.OrgBin
	}

	return result
}

func (d *St) DeviceDelete(ctx context.Context, id string) error {
	return d.db.Exec(ctx, `
		delete from device
		where id = $1
	`, id)
}
