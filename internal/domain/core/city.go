package core

import (
	"context"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
)

type City struct {
	r *St
}

func NewCity(r *St) *City {
	return &City{r: r}
}

func (c *City) ValidateCU(ctx context.Context, obj *entities.CityCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *City) List(ctx context.Context, pars *entities.CityListParsSt) ([]*entities.CitySt, error) {
	items, err := c.r.repo.CityList(ctx, pars)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *City) Get(ctx context.Context, id string, errNE bool) (*entities.CitySt, error) {
	result, err := c.r.repo.CityGet(ctx, id)
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

func (c *City) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.CityIdExists(ctx, id)
}

func (c *City) Create(ctx context.Context, obj *entities.CityCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.CityCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *City) Update(ctx context.Context, id string, obj *entities.CityCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.CityUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *City) Delete(ctx context.Context, id string) error {
	return c.r.repo.CityDelete(ctx, id)
}
