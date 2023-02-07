package entities

type DeviceInputReg struct {
	DeviceId        string `json:"DeviceId"`
	OrganizationBin string `json:"OrganizationBin"`
	TradePointId    int    `json:"TradePointId"`
}

type DeviceOutputReg struct {
	Data       DeviceToken `json:"Data"`
	Message    string      `json:"Message"`
	StatusCode int         `json:"StatusCode"`
}

type DeviceToken struct {
	DeviceToken string `json:"DeviceToken"`
}

type DeviceOutputDel struct {
	StatusCode int `json:"StatusCode"`
}
