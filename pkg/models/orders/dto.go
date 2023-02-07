package orders

type CreateOrderDTO struct {
	OrderNumber    string `json:"order_number"`
	OrganizationId int    `json:"organization_id"`
	Paid           bool   `json:"paid"`
	SentTo1C       bool   `json:"sent_to_1c"`
}
