package old

import "time"

type Order struct {
	Created         time.Time `json:"created"`
	Modified        time.Time `json:"modified"`
	OrderNumber     string    `json:"order_number"`
	OrganizationBin string    `json:"organization_bin"`
	Status          string    `json:"status"`
}

type CreateOrderDTO struct {
	Created         time.Time `json:"created"`
	Modified        time.Time `json:"modified"`
	OrderNumber     string    `json:"order_number"`
	OrganizationBin string    `json:"organization_id"`
	Status          string    `json:"status"`
}

type UnPaidOrder struct {
	Created         time.Time `json:"created"`
	OrderNumber     string    `json:"order_number"`
	OrganizationBin string    `json:"organization_bin"`
	PaymentId       string    `json:"payment_id"`
}

type PaidOrder struct {
	ID              string `json:"id"`
	OrderNumber     string `json:"order_number"`
	OrganizationBin string `json:"organization_bin"`
	PaymentId       string `json:"payment_id"`
}
