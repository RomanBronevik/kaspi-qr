package orders

import (
	"kaspi-qr/internal/organization"
)

type Order struct {
	ID             string                    `json:"id"`
	OrderNumber    string                    `json:"order_number"`
	OrganizationId organization.Organization `json:"organization_id"`
	Paid           bool                      `json:"paid"`
	SentTo1C       bool                      `json:"sent_to_1c"`
}
