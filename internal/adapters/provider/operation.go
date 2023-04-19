package provider

import (
	"time"
)

// request

type OperationGetReqSt struct {
	QrPaymentId int64  `json:"QrPaymentId"`
	DeviceToken string `json:"DeviceToken"`
}

// reply

type OperationStatusRepSt struct {
	BaseRepSt
	Data struct {
		Status string `json:"Status"`
	} `json:"Data"`
}

type OperationGetRepSt struct {
	BaseRepSt
	Data OperationDetailsSt `json:"Data"`
}

// common

type OperationDetailsSt struct {
	QrPaymentId           int64     `json:"QrPaymentId"`
	TotalAmount           float64   `json:"TotalAmount"`
	AvailableReturnAmount float64   `json:"AvailableReturnAmount"`
	TransactionDate       time.Time `json:"TransactionDate"`
}
