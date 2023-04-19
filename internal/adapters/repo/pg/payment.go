package pg

import (
	"context"
	"errors"
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) PaymentGet(ctx context.Context, id string) (*entities.PaymentSt, error) {
	var result entities.PaymentSt

	err := d.db.QueryRow(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.ord_id,
			t.status,
			t.payment_method,
			t.amount
		from payment t
		where t.id = $1
	`, id).Scan(
		&result.Id,
		&result.Created,
		&result.Modified,
		&result.OrdId,
		&result.Status,
		&result.PaymentMethod,
		&result.Amount,
	)
	if errors.Is(err, db.ErrNoRows) {
		return nil, nil
	}

	return &result, err
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
	if pars.PaymentMethod != nil {
		conds = append(conds, "t.payment_method = ${payment_method}")
		args["payment_method"] = *pars.PaymentMethod
	}

	rows, err := d.db.Query(ctx, `
		select
			t.id,
			t.created,
			t.modified,
			t.ord_id,
			t.status,
			t.payment_method,
			t.amount
		from payment t
		`+d.tOptionalWhere(conds)+`
		order by t.name`,
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entities.PaymentSt

	for rows.Next() {
		item := &entities.PaymentSt{}

		err = rows.Scan(
			&item.Id,
			&item.Created,
			&item.Modified,
			&item.OrdId,
			&item.Status,
			&item.PaymentMethod,
			&item.Amount,
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

func (d *St) PaymentIdExists(ctx context.Context, id string) (bool, error) {
	var err error
	var cnt int

	err = d.db.QueryRow(ctx, `
		select count(*)
		from payment
		where id = $1
	`, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) PaymentCreate(ctx context.Context, obj *entities.PaymentCUSt) (string, error) {
	fields := d.paymentGetCUFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId string

	err := d.db.QueryRowM(ctx, `
		insert into payment (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) PaymentUpdate(ctx context.Context, id string, obj *entities.PaymentCUSt) error {
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

	if obj.Status != nil {
		result["status"] = *obj.Status
	}

	if obj.PaymentMethod != nil {
		result["payment_method"] = *obj.PaymentMethod
	}

	if obj.Amount != nil {
		result["amount"] = *obj.Amount
	}

	return result
}

func (d *St) PaymentDelete(ctx context.Context, id string) error {
	return d.db.Exec(ctx, `
		delete from payment
		where id = $1
	`, id)
}
