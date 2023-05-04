package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/dopErrs"
)

func (d *St) CityGet(ctx context.Context, id string) (*entities.CitySt, error) {
	var result entities.CitySt

	err := d.DbQueryRow(ctx, `
		select
			t.id,
			t.code,
			t.name,
			t.org_bin
		from city t
		where t.id = $1
	`, id).Scan(
		&result.Id,
		&result.Code,
		&result.Name,
		&result.OrgBin,
	)
	if errors.Is(err, dopErrs.NoRows) {
		return nil, nil
	}

	return &result, err
}

func (d *St) CityList(ctx context.Context, pars *entities.CityListParsSt) ([]*entities.CitySt, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: text[]))`)
		args["ids"] = *pars.Ids
	}
	if pars.Code != nil {
		conds = append(conds, "t.code = ${code}")
		args["code"] = *pars.Code
	}
	if pars.OrgBin != nil {
		conds = append(conds, "t.org_bin = ${org_bin}")
		args["org_bin"] = *pars.OrgBin
	}

	rows, err := d.DbQueryM(ctx, `
		select
			t.id,
			t.code,
			t.name,
			t.org_bin
		from city t
		`+d.tOptionalWhere(conds)+`
		order by t.name`,
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.CitySt, 0)

	for rows.Next() {
		item := &entities.CitySt{}

		err = rows.Scan(
			&item.Id,
			&item.Code,
			&item.Name,
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

func (d *St) CityIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.DbQueryRow(ctx, `
		select count(*)
		from city
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) CityCreate(ctx context.Context, obj *entities.CityCUSt) (string, error) {
	fields := d.cityGetCUFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId string

	err := d.DbQueryRowM(ctx, `
		insert into city (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) CityUpdate(ctx context.Context, id string, obj *entities.CityCUSt) error {
	fields := d.cityGetCUFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.DbExecM(ctx, `
		update city
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) cityGetCUFields(obj *entities.CityCUSt) map[string]any {
	result := map[string]any{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.Code != nil {
		result["code"] = *obj.Code
	}

	if obj.Name != nil {
		result["name"] = *obj.Name
	}

	if obj.OrgBin != nil {
		result["org_bin"] = *obj.OrgBin
	}

	return result
}

func (d *St) CityDelete(ctx context.Context, id string) error {
	return d.DbExec(ctx, `
		delete from city
		where id = $1
	`, id)
}
