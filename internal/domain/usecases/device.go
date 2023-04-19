package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

func (u *St) DeviceList(ctx context.Context,
	pars *entities.DeviceListParsSt) ([]*entities.DeviceSt, error) {
	//var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Device.List(ctx, pars)
}

func (u *St) DeviceGet(ctx context.Context, id string) (*entities.DeviceSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Device.Get(ctx, id, true)
}

func (u *St) DeviceCreate(ctx context.Context,
	obj *entities.DeviceCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Device.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) DeviceUpdate(ctx context.Context,
	id string, obj *entities.DeviceCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Device.Update(ctx, id, obj)
	})
}

func (u *St) DeviceDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Device.Delete(ctx, id)
	})
}
