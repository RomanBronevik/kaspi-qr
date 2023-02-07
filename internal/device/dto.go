package device

type CreateDeviceDTO struct {
	Token          string `json:"token"`
	OrganizationId int    `json:"organization_id"`
	TradePointId   int    `json:"trade_point_id"`
}
