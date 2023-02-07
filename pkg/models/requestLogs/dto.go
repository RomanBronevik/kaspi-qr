package requestLogs

import (
	"github.com/golang-sql/civil"
	"kaspi-qr/pkg/models/orders"
	"kaspi-qr/pkg/models/organization"
)

type CreateRequestLogDTO struct {
	Source         string                    `json:"source"`
	OrderId        orders.Order              `json:"order_id"`
	OrganizationId organization.Organization `json:"organization_id"`
	RequestTime    civil.DateTime            `json:"request_time"`
	StatusCode     int                       `json:"status_code"`
	Success        bool                      `json:"success"`
}
