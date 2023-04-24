package core

import (
	"context"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/dopErrs"
)

type Src struct {
	r *St
}

func NewSrc(r *St) *Src {
	return &Src{r: r}
}

func (c *Src) ValidateCU(ctx context.Context, obj *entities.SrcCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Src) List(ctx context.Context, pars *entities.SrcListParsSt) ([]*entities.SrcSt, int64, error) {
	items, tCount, err := c.r.repo.SrcList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Src) Get(ctx context.Context, id string, errNE bool) (*entities.SrcSt, error) {
	result, err := c.r.repo.SrcGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, dopErrs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *Src) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.SrcIdExists(ctx, id)
}

func (c *Src) Create(ctx context.Context, obj *entities.SrcCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.SrcCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Src) Update(ctx context.Context, id string, obj *entities.SrcCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.SrcUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Src) Delete(ctx context.Context, id string) error {
	return c.r.repo.SrcDelete(ctx, id)
}
