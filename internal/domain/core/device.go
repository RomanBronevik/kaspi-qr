package core

import (
	"context"
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
	// forCreate := id == ""

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
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

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
	return c.r.repo.DeviceDelete(ctx, id)
}
