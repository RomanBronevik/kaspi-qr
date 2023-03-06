package usecases

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateOrder(ctx *gin.Context, obj *entities.CreateOrderDTO) error {
	err := s.cr.CreateOrder(ctx, obj)

	return err
}

//	func (s *St) UpdateOrder(ctx *gin.Context, obj *entities.Order) error {
//		err := s.cr.UpdateOrder(ctx, obj)
//
//		return err
//	}
func (s *St) DeleteOrder(ctx *gin.Context, orderNumber string) error {
	err := s.cr.DeleteOrder(ctx, orderNumber)

	return err
}

func (s *St) FindAllOrders(ctx *gin.Context) ([]entities.Order, error) {
	orders, err := s.cr.FindAllOrders(ctx)

	return orders, err
}

func (s *St) FindOneOrder(ctx *gin.Context, orderNumber string) (entities.Order, error) {
	order, err := s.cr.FindOneOrder(ctx, orderNumber)

	return order, err
}

func (s *St) ReturnOrder(c *gin.Context, paymentId int) error {
	err := s.cr.ReturnOrder(c, paymentId)

	return err
}
