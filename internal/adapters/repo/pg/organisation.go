package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) OrganisationGet(ctx context.Context, id string) (*entities.OrganisationSt, error) {
	result := &entities.OrganisationSt{}

	err := d.db.QueryRow(ctx, `
		select 
		    id,
		    name,
		    bin
		from organisation
		where id = $1
	`, id).Scan(
		&result.Id,
		&result.Name,
		&result.Bin,
	)
	if !errors.Is(err, db.ErrNoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) OrganisationList(ctx context.Context, pars *entities.OrganisationListParsSt) ([]*entities.OrganisationSt, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: bigint[]))`)
		args["ids"] = *pars.Ids
	}
	if pars.Name != nil {
		conds = append(conds, `t.name = ${name}`)
		args["name"] = *pars.Name
	}

	rows, err := d.db.QueryM(ctx, `
		select
			t.id,
			t.bin,
			t.name
		from organisation t
		`+d.tOptionalWhere(conds),
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.OrganisationSt, 0, 50)

	for rows.Next() {
		rec := &entities.OrganisationSt{}

		err = rows.Scan(
			&rec.Id,
			&rec.Bin,
			&rec.Name,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *St) OrganisationIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.db.QueryRow(ctx, `
		select count(*)
		from organisation
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) OrganisationCreate(ctx context.Context, obj *entities.OrganisationCUSt) (string, error) {
	fields := d.organisationGetCUFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId string

	err := d.db.QueryRowM(ctx, `
		insert into organisation (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) OrganisationUpdate(ctx context.Context, id string, obj *entities.OrganisationCUSt) error {
	fields := d.organisationGetCUFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.db.ExecM(ctx, `
		update organisation
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) organisationGetCUFields(obj *entities.OrganisationCUSt) map[string]any {
	result := map[string]any{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.Name != nil {
		result["name"] = *obj.Name
	}

	if obj.Bin != nil {
		result["bin"] = *obj.Bin
	}

	return result
}

func (d *St) OrganisationDelete(ctx context.Context, id string) error {
	return d.db.Exec(ctx, `
		delete from organisation
		where id = $1
	`, id)
}
