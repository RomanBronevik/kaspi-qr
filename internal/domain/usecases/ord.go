package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

func (u *St) OrdList(ctx context.Context,
	pars *entities.OrdListParsSt) ([]*entities.OrdSt, error) {
	//var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Ord.List(ctx, pars)
}

func (u *St) OrdGet(ctx context.Context, id string) (*entities.OrdSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Ord.Get(ctx, id, true)
}

func (u *St) OrdCreate(ctx context.Context,
	obj *entities.OrdCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Ord.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) OrdUpdate(ctx context.Context,
	id string, obj *entities.OrdCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Ord.Update(ctx, id, obj)
	})
}

func (u *St) OrdDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Ord.Delete(ctx, id)
	})
}