package device

import (
	"kaspi-qr/internal/organization"
	"kaspi-qr/internal/tradePoint"
)

type Device struct {
	ID             string                    `json:"id"`
	Token          string                    `json:"token"`
	OrganizationId organization.Organization `json:"organization_id"`
	TradePointId   tradePoint.TradePoint     `json:"trade_point_id"`
}
