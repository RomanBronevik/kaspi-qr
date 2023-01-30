package handler

type tradePointSt struct {
	StatusCode int           `json:"StatusCode"`
	Message    string        `json:"Message"`
	Data       []*tradePoint `json:"Data"`
}

type tradePoint struct {
	TradePointId   int    `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}

type RegistrationOutputSt struct {
	Data       DeviceToken `json:"Data"`
	StatusCode int         `json:"StatusCode"`
}

type DeviceToken struct {
	DeviceToken string `json:"DeviceToken"`
}

type DeleteOutputSt struct {
	StatusCode int `json:"StatusCode"`
}

type QRToken struct {
}
