package core

import (
	"context"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
)

type Payment struct {
	r *St
}

func NewPayment(r *St) *Payment {
	return &Payment{r: r}
}

func (c *Payment) ValidateCU(ctx context.Context, obj *entities.PaymentCUSt, id int64) error {
	// forCreate := id == 0

	return nil
}

func (c *Payment) List(ctx context.Context, pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, error) {
	items, err := c.r.repo.PaymentList(ctx, pars)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Payment) Get(ctx context.Context, id int64, errNE bool) (*entities.PaymentSt, error) {
	result, err := c.r.repo.PaymentGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, errs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *Payment) IdExists(ctx context.Context, id int64) (bool, error) {
	return c.r.repo.PaymentIdExists(ctx, id)
}

func (c *Payment) Create(ctx context.Context, obj *entities.PaymentCUSt) (int64, error) {
	var err error

	err = c.ValidateCU(ctx, obj, 0)
	if err != nil {
		return 0, err
	}

	// create
	result, err := c.r.repo.PaymentCreate(ctx, obj)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *Payment) Update(ctx context.Context, id int64, obj *entities.PaymentCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.PaymentUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Payment) Delete(ctx context.Context, id int64) error {
	return c.r.repo.PaymentDelete(ctx, id)
}
