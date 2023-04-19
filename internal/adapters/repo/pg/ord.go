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
			t.org_bin,
			t.status
		from ord t
		where t.id = $1
	`, id).Scan(
		&result.Id,
		&result.Created,
		&result.Modified,
		&result.OrgBin,
		&result.Status,
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
	if pars.OrgBin != nil {
		conds = append(conds, "t.org_bin = ${org_bin}")
		args["org_bin"] = *pars.OrgBin
	}
	if pars.Status != nil {
		conds = append(conds, "t.status = ${status}")
		args["status"] = *pars.Status
	}

	rows, err := d.db.Query(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.org_bin,
			t.status
		from ord t
		`+d.tOptionalWhere(conds)+`
		order by t.created`,
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entities.OrdSt

	for rows.Next() {
		item := &entities.OrdSt{}

		err = rows.Scan(
			&item.Id,
			&item.Created,
			&item.Modified,
			&item.OrgBin,
			&item.Status,
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

	if obj.OrgBin != nil {
		result["org_bin"] = *obj.OrgBin
	}

	if obj.Status != nil {
		result["status"] = *obj.Status
	}

	return result
}

func (d *St) OrdDelete(ctx context.Context, id string) error {
	return d.db.Exec(ctx, `
		delete from ord
		where id = $1
	`, id)
}
