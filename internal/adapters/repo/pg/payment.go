package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) PaymentGet(ctx context.Context, id int64) (*entities.PaymentSt, error) {
	var result entities.PaymentSt

	err := d.db.QueryRow(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.ord_id,
			t.link,
			t.status,
			t.status_changed_at,
			t.amount,
			t.expire_dt,
			t.pbo
		from payment t
		where t.id = $1
	`, id).Scan(
		&result.Id,
		&result.Created,
		&result.Modified,
		&result.OrdId,
		&result.Link,
		&result.Status,
		&result.StatusChangedAt,
		&result.Amount,
		&result.ExpireDt,
		&result.Pbo,
	)
	if errors.Is(err, db.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (d *St) PaymentGetLink(ctx context.Context, id int64) (string, error) {
	var result string

	err := d.db.QueryRow(ctx, `
		select link
		from payment
		where id = $1
	`, id).Scan(&result)
	if errors.Is(err, db.ErrNoRows) {
		return "", nil
	}

	return result, err
}

func (d *St) PaymentList(ctx context.Context, pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Ids != nil {
		conds = append(conds, `t.id in (select * from unnest(${ids} :: text[]))`)
		args["ids"] = *pars.Ids
	}
	if pars.OrdId != nil {
		conds = append(conds, "t.ord_id = ${ord_id}")
		args["ord_id"] = *pars.OrdId
	}
	if pars.Status != nil {
		conds = append(conds, "t.status = ${status}")
		args["status"] = *pars.Status
	}
	if pars.Statuses != nil {
		conds = append(conds, `t.status in (select * from unnest(${statuses} :: text[]))`)
		args["statuses"] = *pars.Statuses
	}

	rows, err := d.db.QueryM(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.ord_id,
			t.link,
			t.status,
			t.status_changed_at,
			t.amount,
			t.expire_dt,
			t.pbo
		from payment t
		`+d.tOptionalWhere(conds)+`
		order by t.created`,
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.PaymentSt, 0)

	for rows.Next() {
		item := &entities.PaymentSt{}

		err = rows.Scan(
			&item.Id,
			&item.Created,
			&item.Modified,
			&item.OrdId,
			&item.Link,
			&item.Status,
			&item.StatusChangedAt,
			&item.Amount,
			&item.ExpireDt,
			&item.Pbo,
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

func (d *St) PaymentIdExists(ctx context.Context, id int64) (bool, error) {
	var err error
	var cnt int

	err = d.db.QueryRow(ctx, `
		select count(*)
		from payment
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) PaymentCreate(ctx context.Context, obj *entities.PaymentCUSt) (int64, error) {
	fields := d.paymentGetCUFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId int64

	err := d.db.QueryRowM(ctx, `
		insert into payment (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) PaymentUpdate(ctx context.Context, id int64, obj *entities.PaymentCUSt) error {
	fields := d.paymentGetCUFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.db.ExecM(ctx, `
		update payment
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) paymentGetCUFields(obj *entities.PaymentCUSt) map[string]any {
	result := map[string]any{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.Modified != nil {
		result["modified"] = *obj.Modified
	}

	if obj.OrdId != nil {
		result["ord_id"] = *obj.OrdId
	}

	if obj.Link != nil {
		result["link"] = *obj.Link
	}

	if obj.Status != nil {
		result["status"] = *obj.Status
	}

	if obj.StatusChangedAt != nil {
		result["status_changed_at"] = *obj.StatusChangedAt
	}

	if obj.Amount != nil {
		result["amount"] = *obj.Amount
	}

	if obj.ExpireDt != nil {
		result["expire_dt"] = *obj.ExpireDt
	}

	if obj.Pbo != nil {
		result["pbo"] = *obj.Pbo
	}

	return result
}

func (d *St) PaymentDelete(ctx context.Context, id int64) error {
	return d.db.Exec(ctx, `
		delete from payment
		where id = $1
	`, id)
}
