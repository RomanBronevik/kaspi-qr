package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

func (u *St) CityList(ctx context.Context,
	pars *entities.CityListParsSt) ([]*entities.CitySt, error) {
	//var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.City.List(ctx, pars)
}

func (u *St) CityGet(ctx context.Context, id string) (*entities.CitySt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.City.Get(ctx, id, true)
}

func (u *St) CityCreate(ctx context.Context,
	obj *entities.CityCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.City.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) CityUpdate(ctx context.Context,
	id string, obj *entities.CityCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.City.Update(ctx, id, obj)
	})
}

func (u *St) CityDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.City.Delete(ctx, id)
	})
}
