package payment

import "kaspi-qr/pkg/models/orders"

type Payment struct {
	ID            string       `json:"id"`
	OrderId       orders.Order `json:"order_id"`
	PaymentMethod string       `json:"payment_method"`
	PaymentType   string       `json:"payment_type"`
	Amount        float64      `json:"amount"`
}
