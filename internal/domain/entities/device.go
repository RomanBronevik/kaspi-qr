package entities

type Device struct {
	DeviceId        string `json:"device_id"`
	Token           string `json:"token"`
	OrganizationBin string `json:"organization_bin"`
}

type CreateDeviceDTO struct {
	Token           string `json:"token"`
	DeviceId        string `json:"device_id"`
	OrganizationBin string `json:"organization_id"`
}

type DeviceInputReg struct {
	DeviceId        string `json:"DeviceId"`
	OrganizationBin string `json:"OrganizationBin"`
	TradePointId    string `json:"TradePointId"`
}

type DeviceOutputReg struct {
	Data       *DeviceToken `json:"Data"`
	Message    string       `json:"Message"`
	StatusCode int          `json:"StatusCode"`
}

type DeviceToken struct {
	DeviceToken string `json:"DeviceToken"`
}

type DeviceInputDel struct {
	OrganizationBin string `json:"OrganizationBin"`
	DeviceToken     string `json:"DeviceToken"`
}

type DeviceOutputDel struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
}
