package provider

// request

type DeviceCreateReqSt struct {
	OrganizationBin string `json:"OrganizationBin"`
	DeviceId        string `json:"DeviceId"`
	TradePointId    int64  `json:"TradePointId"`
}

type DeviceDeleteReqSt struct {
	OrganizationBin string `json:"OrganizationBin"`
	DeviceToken     string `json:"DeviceToken"`
}

// reply

type DeviceCreateRepSt struct {
	BaseRepSt
	Data struct {
		DeviceToken string `json:"DeviceToken"`
	} `json:"Data"`
}
