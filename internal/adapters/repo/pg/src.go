package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
)

func (d *St) SrcGet(ctx context.Context, id string) (*entities.SrcSt, error) {
	result := &entities.SrcSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"src"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) SrcList(ctx context.Context, pars *entities.SrcListParsSt) ([]*entities.SrcSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: text[]))`)
		args["ids"] = *pars.Ids
	}

	result := make([]*entities.SrcSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`src t`},
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.id",
		},
	})

	return result, tCount, err
}

func (d *St) SrcIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from src
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) SrcCreate(ctx context.Context, obj *entities.SrcCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "src",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) SrcUpdate(ctx context.Context, id string, obj *entities.SrcCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "src",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) SrcDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "src",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
