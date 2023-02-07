package entities

type TradePointSt struct {
	StatusCode int           `json:"StatusCode"`
	Message    string        `json:"Message"`
	Data       []*tradePoint `json:"Data"`
}

type tradePoint struct {
	TradePointId   int    `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}
