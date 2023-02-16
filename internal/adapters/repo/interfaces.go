package repo

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

type Repo interface {
	CreateDevice(ctx context.Context, device *entities.CreateDeviceDTO) error
	FindAllDevices(ctx context.Context) (u []entities.Device, err error)
	FindOneDevice(ctx context.Context, id string) (entities.Device, error)
	//UpdateDevice(ctx context.Context, token string) error
	DeleteDevice(ctx context.Context, bin string, token string) error

	CreateOrganization(ctx context.Context, organization *entities.CreateOrganizationDTO) error
	FindAllOrganizations(ctx context.Context) (u []entities.Organization, err error)
	FindOneOrganization(ctx context.Context, bin string) (entities.Organization, error)
	//UpdateOrganization(ctx context.Context, organization *entities.Organization) error
	DeleteOrganization(ctx context.Context, bin string) error

	CreateTradePoint(ctx context.Context, organization *entities.CreateTradePointDTO) error
	FindAllTradePoints(ctx context.Context) (u []entities.TradePoint, err error)
	FindOneTradePoint(ctx context.Context, id string) (entities.TradePoint, error)
	//UpdateTradePoint(ctx context.Context, tradePoint *entities.TradePoint) error
	DeleteTradePoint(ctx context.Context, bin string, tradePointId string) error

	CreatePayment(ctx context.Context, payment *entities.CreatePaymentDTO) error
	FindAllPayments(ctx context.Context) (u []entities.Payment, err error)
	FindOnePayment(ctx context.Context, orderNumber string) (entities.Payment, error)
	//UpdatePayment(ctx context.Context, payment *entities.Payment) error
	DeletePayment(ctx context.Context, orderNumber string) error

	CreateOrder(ctx context.Context, order *entities.CreateOrderDTO) error
	FindAllOrders(ctx context.Context) (u []entities.Order, err error)
	FindOneOrder(ctx context.Context, orderNumber string) (entities.Order, error)
	//UpdateOrder(ctx context.Context, order *entities.Order) error
	DeleteOrder(ctx context.Context, orderNumber string) error

	CreateCity(ctx context.Context, order *entities.CreateCityDTO) error
	FindAllCities(ctx context.Context) (u []entities.City, err error)
	FindOneCity(ctx context.Context, CityName string) (entities.City, error)
	//UpdateCity(ctx context.Context, order *entities.City) error
	DeleteCity(ctx context.Context, id string) error

	//CreateRequestLog(ctx context.Context, order *entities.CreateOrderDTO) error
	//FindAllRequestLogs(ctx context.Context) (u []entities.Order, err error)
	//FindOneRequestLog(ctx context.Context, orderNumber int) (entities.Order, error)
	//DeleteRequestLog(ctx context.Context, id string) error
}
