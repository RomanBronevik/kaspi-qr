package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

func (u *St) SrcList(ctx context.Context,
	pars *entities.SrcListParsSt) ([]*entities.SrcSt, int64, error) {
	//var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Src.List(ctx, pars)
}

func (u *St) SrcGet(ctx context.Context, id string) (*entities.SrcSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Src.Get(ctx, id, true)
}

func (u *St) SrcCreate(ctx context.Context,
	obj *entities.SrcCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Src.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) SrcUpdate(ctx context.Context,
	id string, obj *entities.SrcCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Src.Update(ctx, id, obj)
	})
}

func (u *St) SrcDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Src.Delete(ctx, id)
	})
}
