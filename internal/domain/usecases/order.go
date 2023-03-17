package usecases

import (
	"kaspi-qr/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (u *St) CreateOrder(ctx *gin.Context, obj *entities.CreateOrderDTO) error {
	err := u.cr.CreateOrder(ctx, obj)

	return err
}

//	func (s *St) UpdateOrder(ctx *gin.Context, obj *entities.Order) error {
//		err := s.cr.UpdateOrder(ctx, obj)
//
//		return err
//	}
func (u *St) DeleteOrder(ctx *gin.Context, orderNumber string) error {
	err := u.cr.DeleteOrder(ctx, orderNumber)

	return err
}

func (u *St) FindAllOrders(ctx *gin.Context) ([]entities.Order, error) {
	orders, err := u.cr.FindAllOrders(ctx)

	return orders, err
}

func (u *St) FindOneOrder(ctx *gin.Context, orderNumber string) (entities.Order, error) {
	order, err := u.cr.FindOneOrder(ctx, orderNumber)

	return order, err
}

func (u *St) ReturnOrder(c *gin.Context, paymentId int) error {
	err := u.cr.ReturnOrder(c, paymentId)

	return err
}
