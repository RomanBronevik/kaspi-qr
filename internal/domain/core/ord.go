package core

import (
	"context"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"time"
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

	// SrcId
	if forCreate && obj.SrcId == nil {
		return errs.SrcRequired
	}
	if obj.SrcId != nil {
		if !cns.OrdSrcIsValid(*obj.SrcId) {
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

func (c *Ord) Create(ctx context.Context, obj *entities.OrdCUSt) (*entities.OrdCreateRepSt, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return nil, err
	}

	ordStatusCreated := cns.OrdStatusCreated

	// get ord
	ord, err := c.Get(ctx, *obj.Id, false)
	if err != nil {
		return nil, err
	}

	if ord == nil {
		obj.Status = &ordStatusCreated

		// find device
		device, err := c.r.Device.GetForCityId(ctx, *obj.CityId)
		if err != nil {
			return nil, err
		}
		if device == nil {
			return nil, errs.DeviceNotFound
		}

		obj.DeviceId = &device.Id

		// create ord
		_, err = c.create(ctx, obj)
		if err != nil {
			return nil, err
		}
	} else {
		if ord.Status != cns.OrdStatusCreated && ord.Status != cns.OrdStatusError {
			return nil, errs.OrderAlreadyPaid
		}

		uObj := &entities.OrdCUSt{}

		if ord.Status != cns.OrdStatusCreated {
			uObj.Status = &ordStatusCreated
		}

		if ord.SrcId != *obj.SrcId {
			uObj.SrcId = obj.SrcId
		}

		if ord.CityId != *obj.CityId {
			uObj.CityId = obj.CityId

			// find device
			device, err := c.r.Device.GetForCityId(ctx, *obj.CityId)
			if err != nil {
				return nil, err
			}
			if device == nil {
				return nil, errs.DeviceNotFound
			}

			obj.DeviceId = &device.Id
		}

		if ord.Amount != *obj.Amount {
			uObj.Amount = obj.Amount
		}

		if ord.Platform != *obj.Platform {
			uObj.Platform = obj.Platform
		}

		// update ord
		if *uObj != (entities.OrdCUSt{}) {
			err = c.Update(ctx, ord.Id, uObj)
			if err != nil {
				return nil, err
			}
		}
	}

	// create payment
	payment, err := c.r.Payment.CreateForOrd(ctx, *obj.Id)
	if err != nil {
		return nil, err
	}

	return &entities.OrdCreateRepSt{
		PaymentId: payment.Id,
		QrUrl:     c.r.Payment.createQrUrl(payment.Id),
		QrCode:    payment.Link,
	}, nil
}

func (c *Ord) create(ctx context.Context, obj *entities.OrdCUSt) (string, error) {
	return c.r.repo.OrdCreate(ctx, obj)
}

func (c *Ord) Update(ctx context.Context, id string, obj *entities.OrdCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	now := time.Now()

	obj.Modified = &now

	err = c.r.repo.OrdUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Ord) Delete(ctx context.Context, id string) error {
	return c.r.repo.OrdDelete(ctx, id)
}
