package provider

// request

type ReturnReqSt struct {
	DeviceToken     string  `json:"DeviceToken"`
	OrganizationBin string  `json:"OrganizationBin"`
	QrPaymentId     int     `json:"QrPaymentId"`
	Amount          float64 `json:"Amount"`
}

// reply

type ReturnRepSt struct {
	BaseRepSt
	ReturnOperationDataSt struct {
		ReturnOperationId int64 `json:"ReturnOperationId"`
	} `json:"ReturnOperationDataSt"`
}
