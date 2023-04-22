package old

import (
	"time"
)

type OperationGetSt struct {
	QrPaymentId int    `json:"QrPaymentId"`
	DeviceToken string `json:"DeviceToken"`
}

type OperationDetails struct {
	Data       *OperationDetailsSt `json:"Data"`
	StatusCode int                 `json:"StatusCode"`
}

type OperationDetailsSt struct {
	QrPaymentId           int       `json:"QrPaymentId"`
	TotalAmount           float64   `json:"TotalAmount"`
	AvailableReturnAmount float64   `json:"AvailableReturnAmount"`
	TransactionDate       time.Time `json:"TransactionDate"`
}

type OperationDetailsInput struct {
	QrPaymentId     int    `json:"QrPaymentId"`
	OrganizationBin string `json:"OrganizationBin"`
}
