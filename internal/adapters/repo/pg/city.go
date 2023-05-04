package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
)

func (d *St) CityGet(ctx context.Context, id string) (*entities.CitySt, error) {
	result := &entities.CitySt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"city"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
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

	result := make([]*entities.CitySt, 0, 100)

	_, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`city t`},
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.name",
		},
	})

	return result, err
}

func (d *St) CityIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
		select count(*)
		from city
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) CityCreate(ctx context.Context, obj *entities.CityCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "city",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) CityUpdate(ctx context.Context, id string, obj *entities.CityCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "city",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) CityDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "city",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
