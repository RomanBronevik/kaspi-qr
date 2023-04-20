package core

import (
	"context"
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
)

type Device struct {
	r *St
}

func NewDevice(r *St) *Device {
	return &Device{r: r}
}

func (c *Device) ValidateCU(ctx context.Context, obj *entities.DeviceCUSt, id string) error {
	forCreate := id == ""

	// Id
	if forCreate && obj.Id == nil {
		return errs.DeviceIdRequired
	}
	if obj.Id != nil {
		if *obj.Id == "" {
			return errs.DeviceIdRequired
		}
		if len([]rune(*obj.Id)) > 64 {
			return errs.DeviceIdTooLong
		}
	}

	// TradePointId
	if forCreate && obj.TradePointId == nil {
		return errs.TradePointIdRequired
	}
	if obj.TradePointId != nil {
		if *obj.TradePointId <= 0 {
			return errs.TradePointIdRequired
		}
	}

	// OrgBin
	if forCreate && obj.OrgBin == nil {
		return errs.OrgBinRequired
	}
	if obj.OrgBin != nil {
		if *obj.OrgBin == "" {
			return errs.OrgBinRequired
		}
		if len([]rune(*obj.Id)) > 30 {
			return errs.OrgBinTooLong
		}
	}

	return nil
}

func (c *Device) List(ctx context.Context, pars *entities.DeviceListParsSt) ([]*entities.DeviceSt, error) {
	items, err := c.r.repo.DeviceList(ctx, pars)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Device) Get(ctx context.Context, id string, errNE bool) (*entities.DeviceSt, error) {
	result, err := c.r.repo.DeviceGet(ctx, id)
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

func (c *Device) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.DeviceIdExists(ctx, id)
}

func (c *Device) Create(ctx context.Context, obj *entities.DeviceCUSt) (string, error) {
	err := c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	exists, err := c.IdExists(ctx, *obj.Id)
	if err != nil {
		return "", err
	}
	if exists {
		return *obj.Id, nil
	}

	token, err := c.r.prv.DeviceCreate(provider.DeviceCreateReqSt{
		OrganizationBin: *obj.OrgBin,
		DeviceId:        *obj.Id,
		TradePointId:    *obj.TradePointId,
	})
	if err != nil {
		return "", errs.Err(err.Error())
	}

	obj.Token = &token

	// create
	result, err := c.r.repo.DeviceCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Device) Update(ctx context.Context, id string, obj *entities.DeviceCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.DeviceUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Device) Delete(ctx context.Context, id string) error {
	item, err := c.Get(ctx, id, true)
	if err != nil {
		return err
	}

	err = c.r.prv.DeviceDelete(provider.DeviceDeleteReqSt{
		OrganizationBin: item.OrgBin,
		DeviceToken:     item.Token,
	})
	if err != nil {
		return errs.Err(err.Error())
	}

	return c.r.repo.DeviceDelete(ctx, id)
}
