package core

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreatePayment(ctx *gin.Context, obj *entities.CreatePaymentDTO) error {
	err := s.repo.CreatePayment(ctx, obj)

	return err
}

func (s *St) DeletePayment(ctx *gin.Context, orderNumber string) error {
	err := s.repo.DeletePayment(ctx, orderNumber)

	return err
}

//
//func (s *St) UpdatePayment(ctx *gin.Context, obj *entities.Payment) error {
//	err := s.repo.UpdatePayment(ctx, obj)
//
//	return err
//}

func (s *St) FindAllPayments(ctx *gin.Context) ([]entities.Payment, error) {
	payments, err := s.repo.FindAllPayments(ctx)

	return payments, err
}

func (s *St) FindOnePayment(ctx *gin.Context, obj *entities.Payment) (entities.Payment, error) {
	payment, err := s.repo.FindOnePayment(ctx, obj.OrderNumber)

	return payment, err
}
