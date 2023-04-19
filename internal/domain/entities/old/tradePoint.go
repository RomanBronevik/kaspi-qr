package old

type TradePoint struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	TradePointId    string `json:"trade_point_id"`
	OrganizationBin string `json:"organization_bin"`
}

type CreateTradePointDTO struct {
	Name            string `json:"name"`
	TradePointId    string `json:"trade_point_id"`
	OrganizationBin string `json:"organization_id"`
}

type TradePointSt struct {
	StatusCode int                  `json:"StatusCode"`
	Message    string               `json:"Message"`
	Data       []*KaspiTradePointSt `json:"Data"`
}

type KaspiTradePointSt struct {
	TradePointId   int    `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}
