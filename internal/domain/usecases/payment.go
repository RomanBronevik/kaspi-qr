package usecases

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreatePayment(ctx *gin.Context, obj *entities.CreatePaymentDTO) error {
	err := s.cr.CreatePayment(ctx, obj)

	return err
}

//	func (s *St) UpdatePayment(ctx *gin.Context, obj *entities.Payment) error {
//		err := s.cr.UpdatePayment(ctx, obj)
//
//		return err
//	}
func (s *St) DeletePayment(ctx *gin.Context, orderNumber string) error {
	err := s.cr.DeletePayment(ctx, orderNumber)

	return err
}

func (s *St) FindAllPayments(ctx *gin.Context) ([]entities.Payment, error) {
	payments, err := s.cr.FindAllPayments(ctx)

	return payments, err
}

func (s *St) FindOnePaymentByOrderNumber(ctx *gin.Context, orderNumber string) (entities.Payment, error) {
	payment, err := s.cr.FindOnePaymentByOrderNumber(ctx, orderNumber)

	return payment, err
}

func (s *St) FindOnePayment(ctx *gin.Context, paymentId string) (entities.Payment, error) {
	payment, err := s.cr.FindOnePaymentByPaymentId(ctx, paymentId)

	return payment, err
}

func (s *St) CheckPaymentStatus(ctx *gin.Context) error {
	err := s.cr.CheckPaymentStatus(ctx)

	return err
}
