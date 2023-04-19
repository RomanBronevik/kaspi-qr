package old

import "github.com/golang-sql/civil"

type OperationGetSt struct {
	QrPaymentId int    `json:"QrPaymentId"`
	DeviceToken string `json:"DeviceToken"`
}

type OperationDetails struct {
	Data       *OperationDetailsSt `json:"Data"`
	StatusCode int                 `json:"StatusCode"`
}

type OperationDetailsSt struct {
	QrPaymentId           int            `json:"QrPaymentId"`
	TotalAmount           float64        `json:"TotalAmount"`
	AvailableReturnAmount float64        `json:"AvailableReturnAmount"`
	TransactionDate       civil.DateTime `json:"TransactionDate"`
}

type OperationDetailsInput struct {
	QrPaymentId     int    `json:"QrPaymentId"`
	OrganizationBin string `json:"OrganizationBin"`
}
