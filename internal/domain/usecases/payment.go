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

func (s *St) FindOnePayment(ctx *gin.Context, obj *entities.Payment) (entities.Payment, error) {
	payment, err := s.cr.FindOnePayment(ctx, obj)

	return payment, err
}
