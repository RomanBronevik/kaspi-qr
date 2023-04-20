package core

import (
	"context"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
)

type Ord struct {
	r *St
}

func NewOrd(r *St) *Ord {
	return &Ord{r: r}
}

func (c *Ord) ValidateCU(ctx context.Context, obj *entities.OrdCUSt, id string) error {
	forCreate := id == ""

	// Id
	if forCreate && obj.Id == nil {
		return errs.IdRequired
	}
	if obj.Id != nil {
		if *obj.Id == "" {
			return errs.IdRequired
		}
		if len(*obj.Id) > 200 {
			return errs.IdTooLong
		}
	}

	// Src
	if forCreate && obj.Src == nil {
		return errs.SrcRequired
	}
	if obj.Src != nil {
		if !cns.OrdSrcIsValid(*obj.Src) {
			return errs.BadSrc
		}
	}

	// CityCode
	if forCreate && obj.CityCode == nil {
		return errs.CityCodeRequired
	}
	if obj.CityCode != nil {
		if *obj.CityCode == "" {
			return errs.CityCodeRequired
		}
		cities, err := c.r.City.List(ctx, &entities.CityListParsSt{Code: obj.CityCode})
		if err != nil {
			return err
		}
		if len(cities) == 0 {
			return errs.CityNotFound
		}
		obj.CityId = &cities[0].Id
	}

	// Amount
	if forCreate && obj.Amount == nil {
		return errs.AmountRequired
	}
	if obj.Amount != nil {
		if *obj.Amount <= 0 {
			return errs.AmountMustBePositive
		}
	}

	// Platform
	if forCreate && obj.Platform == nil {
		return errs.PlatformRequired
	}
	if obj.Platform != nil {
		if !cns.PlatformIsValid(*obj.Platform) {
			return errs.BadPlatform
		}
	}

	return nil
}

func (c *Ord) List(ctx context.Context, pars *entities.OrdListParsSt) ([]*entities.OrdSt, error) {
	items, err := c.r.repo.OrdList(ctx, pars)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Ord) Get(ctx context.Context, id string, errNE bool) (*entities.OrdSt, error) {
	result, err := c.r.repo.OrdGet(ctx, id)
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

func (c *Ord) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.OrdIdExists(ctx, id)
}

func (c *Ord) Create(ctx context.Context, obj *entities.OrdCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.OrdCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Ord) Update(ctx context.Context, id string, obj *entities.OrdCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.OrdUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Ord) Delete(ctx context.Context, id string) error {
	return c.r.repo.OrdDelete(ctx, id)
}
