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

func (c *Payment) ValidateCU(ctx context.Context, obj *entities.PaymentCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Payment) List(ctx context.Context, pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, error) {
	items, err := c.r.repo.PaymentList(ctx, pars)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Payment) Get(ctx context.Context, id string, errNE bool) (*entities.PaymentSt, error) {
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

func (c *Payment) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.PaymentIdExists(ctx, id)
}

func (c *Payment) Create(ctx context.Context, obj *entities.PaymentCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.PaymentCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Payment) Update(ctx context.Context, id string, obj *entities.PaymentCUSt) error {
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

func (c *Payment) Delete(ctx context.Context, id string) error {
	return c.r.repo.PaymentDelete(ctx, id)
}
