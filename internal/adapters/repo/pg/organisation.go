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
		select id, name, bin
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
			t.name,
			t.bin
		from organisation t
		`+d.tOptionalWhere(conds),
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.OrganisationSt, 0, 100)

	return result, tCount, err
}

func (d *St) OrganisationIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from organisation
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) OrganisationCreate(ctx context.Context, obj *entities.OrganisationCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "organisation",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) OrganisationUpdate(ctx context.Context, id string, obj *entities.OrganisationCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "organisation",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) OrganisationDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "organisation",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
