package payment

type CreatePaymentDTO struct {
	OrderId       int     `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	PaymentType   string  `json:"payment_type"`
	Amount        float64 `json:"amount"`
}
