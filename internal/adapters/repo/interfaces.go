package repo

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

type Repo interface {
	CreateDevice(ctx context.Context, device *entities.CreateDeviceDTO) error
	FindAllDevices(ctx context.Context) (u []entities.Device, err error)
	FindOneDevice(ctx context.Context, OrganizationBin string) (entities.Device, error)
	DeleteDevice(ctx context.Context, bin string, token string) error

	CreateOrganization(ctx context.Context, organization *entities.CreateOrganizationDTO) error
	FindAllOrganizations(ctx context.Context) (u []entities.Organization, err error)
	FindOneOrganization(ctx context.Context, bin string) (entities.Organization, error)
	DeleteOrganization(ctx context.Context, bin string) error

	CreatePayment(ctx context.Context, payment *entities.CreatePaymentDTO) error
	FindAllPayments(ctx context.Context) (u []entities.Payment, err error)
	FindOnePaymentByPaymentId(ctx context.Context, paymentId string) (entities.Payment, error)
	FindOnePaymentByOrderNumber(ctx context.Context, orderNumber string) (entities.Payment, error)
	DeletePayment(ctx context.Context, orderNumber string) error
	UpdatePaymentRecordsToFail(ctx context.Context, orderNumber string) error
	FindLastPaymentByDesc(ctx context.Context, orderNumber string) (entities.Payment, error)
	UpdatePaymentStatus(ctx context.Context, paymentId string, status string) error

	CreateOrder(ctx context.Context, order *entities.CreateOrderDTO) error
	FindAllOrders(ctx context.Context) (u []entities.Order, err error)
	FindAllUnpaidOrders(ctx context.Context) (u []entities.UnPaidOrder, err error)
	FindOneOrder(ctx context.Context, orderNumber string) (entities.Order, error)
	DeleteOrder(ctx context.Context, orderNumber string) error
	FindAllPaidOrders(ctx context.Context) (u []entities.PaidOrder, err error)
	UpdateOrderStatus(ctx context.Context, orderNumber string, status string) error

	CreateCity(ctx context.Context, city *entities.CreateCityDTO) error
	FindAllCities(ctx context.Context) (u []entities.City, err error)
	FindOneCityByCityCode(ctx context.Context, code string) (entities.City, error)
	DeleteCity(ctx context.Context, id string) error
	DeleteCities(ctx context.Context) error
}
