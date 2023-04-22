package repo

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

type Repo interface {
	// city
	CityGet(ctx context.Context, id string) (*entities.CitySt, error)
	CityList(ctx context.Context, pars *entities.CityListParsSt) ([]*entities.CitySt, error)
	CityIdExists(ctx context.Context, id string) (bool, error)
	CityCreate(ctx context.Context, obj *entities.CityCUSt) (string, error)
	CityUpdate(ctx context.Context, id string, obj *entities.CityCUSt) error
	CityDelete(ctx context.Context, id string) error

	// device
	DeviceGet(ctx context.Context, id string) (*entities.DeviceSt, error)
	DeviceGetIdForCityId(ctx context.Context, cityId string) (string, error)
	DeviceList(ctx context.Context, pars *entities.DeviceListParsSt) ([]*entities.DeviceSt, error)
	DeviceIdExists(ctx context.Context, id string) (bool, error)
	DeviceCreate(ctx context.Context, obj *entities.DeviceCUSt) (string, error)
	DeviceUpdate(ctx context.Context, id string, obj *entities.DeviceCUSt) error
	DeviceDelete(ctx context.Context, id string) error

	// ord
	OrdGet(ctx context.Context, id string) (*entities.OrdSt, error)
	OrdList(ctx context.Context, pars *entities.OrdListParsSt) ([]*entities.OrdSt, error)
	OrdIdExists(ctx context.Context, id string) (bool, error)
	OrdCreate(ctx context.Context, obj *entities.OrdCUSt) (string, error)
	OrdUpdate(ctx context.Context, id string, obj *entities.OrdCUSt) error
	OrdDelete(ctx context.Context, id string) error

	// payment
	PaymentGet(ctx context.Context, id int64) (*entities.PaymentSt, error)
	PaymentList(ctx context.Context, pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, error)
	PaymentIdExists(ctx context.Context, id int64) (bool, error)
	PaymentCreate(ctx context.Context, obj *entities.PaymentCUSt) (int64, error)
	PaymentUpdate(ctx context.Context, id int64, obj *entities.PaymentCUSt) error
	PaymentDelete(ctx context.Context, id int64) error
}
