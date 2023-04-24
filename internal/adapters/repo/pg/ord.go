package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) OrdGet(ctx context.Context, id string) (*entities.OrdSt, error) {
	var result entities.OrdSt

	err := d.db.QueryRow(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.src_id,
			t.device_id,
			t.city_id,
			t.amount,
			t.status,
			t.platform
		from ord t
		where t.id = $1
	`, id).Scan(
		&result.Id,
		&result.Created,
		&result.Modified,
		&result.SrcId,
		&result.DeviceId,
		&result.CityId,
		&result.Amount,
		&result.Status,
		&result.Platform,
	)
	if errors.Is(err, db.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (d *St) OrdList(ctx context.Context, pars *entities.OrdListParsSt) ([]*entities.OrdSt, error) {
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

	rows, err := d.db.QueryM(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.src_id,
			t.device_id,
			t.city_id,
			t.amount,
			t.status,
			t.platform
		from ord t
		`+d.tOptionalWhere(conds)+`
		order by t.created`,
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.OrdSt, 0)

	for rows.Next() {
		item := &entities.OrdSt{}

		err = rows.Scan(
			&item.Id,
			&item.Created,
			&item.Modified,
			&item.SrcId,
			&item.DeviceId,
			&item.CityId,
			&item.Amount,
			&item.Status,
			&item.Platform,
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

func (d *St) OrdIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.db.QueryRow(ctx, `
		select count(*)
		from ord
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) OrdCreate(ctx context.Context, obj *entities.OrdCUSt) (string, error) {
	fields := d.ordGetCUFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId string

	err := d.db.QueryRowM(ctx, `
		insert into ord (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) OrdUpdate(ctx context.Context, id string, obj *entities.OrdCUSt) error {
	fields := d.ordGetCUFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.db.ExecM(ctx, `
		update ord
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) ordGetCUFields(obj *entities.OrdCUSt) map[string]any {
	result := map[string]any{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.Modified != nil {
		result["modified"] = *obj.Modified
	}

	if obj.SrcId != nil {
		result["src_id"] = *obj.SrcId
	}

	if obj.DeviceId != nil {
		result["device_id"] = *obj.DeviceId
	}

	if obj.CityId != nil {
		result["city_id"] = *obj.CityId
	}

	if obj.Amount != nil {
		result["amount"] = *obj.Amount
	}

	if obj.Status != nil {
		result["status"] = *obj.Status
	}

	if obj.Platform != nil {
		result["platform"] = *obj.Platform
	}

	return result
}

func (d *St) OrdDelete(ctx context.Context, id string) error {
	return d.db.Exec(ctx, `
		delete from ord
		where id = $1
	`, id)
}
