package old

type ReturnSt struct {
	StatusCode            int                    `json:"StatusCode"`
	Message               string                 `json:"Message"`
	ReturnOperationDataSt *ReturnOperationDataSt `json:"ReturnOperationDataSt"`
}

type ReturnOperationDataSt struct {
	ReturnOperationId int `json:"ReturnOperationId"`
}

type ReturnRequestInput struct {
	DeviceToken     string  `json:"DeviceToken"`
	OrganizationBin string  `json:"OrganizationBin"`
	QrPaymentId     int     `json:"QrPaymentId"`
	Amount          float64 `json:"Amount"`
}

type ReturnInput struct {
	OrganizationBin string  `json:"OrganizationBin"`
	QrPaymentId     int     `json:"QrPaymentId"`
	Amount          float64 `json:"Amount"`
}
