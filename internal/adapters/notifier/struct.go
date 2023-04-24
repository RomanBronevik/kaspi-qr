package notifier

type OrderStatusChangeReqSt struct {
	OrdId     string `json:"ord_id"`
	PaymentId int64  `json:"payment_id"`
	Status    string `json:"status"`
}
