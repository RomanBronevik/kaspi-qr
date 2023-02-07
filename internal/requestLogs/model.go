package requestLogs

import (
	"github.com/golang-sql/civil"
	"kaspi-qr/internal/orders"
	"kaspi-qr/internal/organization"
)

type RequestLog struct {
	Source         string                    `json:"source"`
	OrderId        orders.Order              `json:"order_id"`
	OrganizationId organization.Organization `json:"organization_id"`
	RequestTime    civil.DateTime            `json:"request_time"`
	StatusCode     int                       `json:"status_code"`
	Success        bool                      `json:"success"`
}
