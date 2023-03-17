package usecases

import (
	"kaspi-qr/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (u *St) CreatePayment(ctx *gin.Context, obj *entities.CreatePaymentDTO) error {
	err := u.cr.CreatePayment(ctx, obj)

	return err
}

//	func (s *St) UpdatePayment(ctx *gin.Context, obj *entities.Payment) error {
//		err := s.cr.UpdatePayment(ctx, obj)
//
//		return err
//	}
func (u *St) DeletePayment(ctx *gin.Context, orderNumber string) error {
	err := u.cr.DeletePayment(ctx, orderNumber)

	return err
}

func (u *St) FindAllPayments(ctx *gin.Context) ([]entities.Payment, error) {
	payments, err := u.cr.FindAllPayments(ctx)

	return payments, err
}

func (u *St) FindOnePaymentByOrderNumber(ctx *gin.Context, orderNumber string) (entities.Payment, error) {
	payment, err := u.cr.FindOnePaymentByOrderNumber(ctx, orderNumber)

	return payment, err
}

func (u *St) FindOnePayment(ctx *gin.Context, paymentId string) (entities.Payment, error) {
	payment, err := u.cr.FindOnePaymentByPaymentId(ctx, paymentId)

	return payment, err
}

func (u *St) CheckPaymentStatus(ctx *gin.Context) error {
	err := u.cr.CheckPaymentStatus(ctx)

	return err
}
