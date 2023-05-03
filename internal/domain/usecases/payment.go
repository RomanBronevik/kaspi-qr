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

func (u *St) PaymentGet(ctx context.Context, id int64) (*entities.PaymentSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Payment.Get(ctx, id, true)
}

func (u *St) PaymentUpdate(ctx context.Context,
	id int64, obj *entities.PaymentCUSt) error {
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
	id int64) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Payment.Delete(ctx, id)
	})
}

func (u *St) PaymentGetQrPicture(ctx context.Context,
	id int64) ([]byte, error) {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.cr.Payment.GetQrPicture(ctx, id)
}

// for testing

func (u *St) PaymentEmuPaymentScan(ctx context.Context,
	id int64) error {
	return u.cr.Payment.EmuPaymentScan(ctx, id)
}

func (u *St) PaymentEmuPaymentScanError(ctx context.Context,
	id int64) error {
	return u.cr.Payment.EmuPaymentScanError(ctx, id)
}

func (u *St) PaymentEmuPaymentConfirm(ctx context.Context,
	id int64) error {
	return u.cr.Payment.EmuPaymentConfirm(ctx, id)
}

func (u *St) PaymentEmuPaymentConfirmError(ctx context.Context,
	id int64) error {
	return u.cr.Payment.EmuPaymentConfirmError(ctx, id)
}
