package provider

// request

type DeviceCreateReqSt struct {
	OrganizationBin string `json:"OrganizationBin"`
	DeviceId        string `json:"DeviceId"`
	TradePointId    string `json:"TradePointId"`
}

type DeviceRemoveReqSt struct {
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
