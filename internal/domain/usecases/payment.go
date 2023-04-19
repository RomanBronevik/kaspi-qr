package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

func (u *St) PaymentList(ctx context.Context,
	pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, error) {
	//var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Payment.List(ctx, pars)
}

func (u *St) PaymentGet(ctx context.Context, id string) (*entities.PaymentSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Payment.Get(ctx, id, true)
}

func (u *St) PaymentCreate(ctx context.Context,
	obj *entities.PaymentCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Payment.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) PaymentUpdate(ctx context.Context,
	id string, obj *entities.PaymentCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Payment.Update(ctx, id, obj)
	})
}

func (u *St) PaymentDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Payment.Delete(ctx, id)
	})
}
