package tradePoint

type CreateTradePointDTO struct {
	Name           string `json:"name"`
	OrganizationId int    `json:"organization_id"`
}
