package old

import (
	"errors"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *St) CreateOrder(ctx *gin.Context, obj *entities.CreateOrderDTO) error {
	err := s.repo.CreateOrder(ctx, obj)

	return err
}

func (s *St) DeleteOrder(ctx *gin.Context, orderNumber string) error {
	err := s.repo.DeleteOrder(ctx, orderNumber)

	return err
}

func (s *St) FindAllOrders(ctx *gin.Context) ([]entities.Order, error) {
	orders, err := s.repo.FindAllOrders(ctx)

	return orders, err
}

func (s *St) FindAllUnpaidOrders(ctx *gin.Context) ([]entities.UnPaidOrder, error) {
	orders, err := s.repo.FindAllUnpaidOrders(ctx)

	return orders, err
}

func (s *St) FindOneOrder(ctx *gin.Context, orderNumber string) (entities.Order, error) {
	order, err := s.repo.FindOneOrder(ctx, orderNumber)

	return order, err
}

func (s *St) UpdateOrderStatus(ctx *gin.Context, orderNumber string, status string) error {
	err := s.repo.UpdateOrderStatus(ctx, orderNumber, status)
	return err
}

func (s *St) createOrderRecord(c *gin.Context, input entities.KaspiPaymentInput) error {
	orderExist, err := s.orderAlreadyExist(c, input.ExternalId)

	if orderExist {
		return nil
	}

	curTime := time.Now().Local()

	dtoSt := entities.CreateOrderDTO{
		Created:         curTime,
		Modified:        curTime,
		OrderNumber:     input.ExternalId,
		OrganizationBin: input.OrganizationBin,
		Status:          cns.StatusCreated,
	}

	err = s.CreateOrder(c, &dtoSt)
	if err != nil {
		return err
	}
	return nil
}

func (s *St) orderAlreadyExist(c *gin.Context, orderNumber string) (bool, error) {
	order, err := s.FindOneOrder(c, orderNumber)
	if err != nil {
		return false, err
	}

	emptyOrder := entities.Order{}

	if order == emptyOrder {
		return false, nil
	}
	return true, nil
}

func isEmptyPayment(payment entities.Payment) bool {
	empty := entities.Payment{}

	if empty == payment {
		return true
	}

	return false
}

func (s *St) ReturnOrder(c *gin.Context, paymentId int) error {
	strPaymentId := strconv.Itoa(paymentId)

	payment, err := s.FindOnePaymentByPaymentId(c, strPaymentId)
	if err != nil {
		return err
	}

	if isEmptyPayment(payment) {
		return errors.New("Payment doesn't exist")
	}

	orderNumber := payment.OrderNumber

	order, err := s.FindOneOrder(c, orderNumber)

	if err != nil {
		return err
	}

	if !s.StatusPaid(order.Status) {
		return errors.New("Order not paid or already refunded")
	}

	err = s.UpdateOrderStatus(c, orderNumber, cns.StatusRefund)

	if err != nil {
		return err
	}

	err = s.UpdatePaymentStatus(c, strPaymentId, cns.StatusRefund)

	if err != nil {
		return err
	}

	return nil
}
