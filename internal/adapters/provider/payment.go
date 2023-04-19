package provider

import (
	"time"
)

// request

type PaymentCreateReqSt struct {
	OrganizationBin string  `json:"OrganizationBin"`
	DeviceToken     string  `json:"DeviceToken"`
	Amount          float64 `json:"Amount"`
	ExternalId      string  `json:"ExternalId"`
}

// reply

type PaymentCreateRepSt struct {
	BaseRepSt
	Data PaymentSt `json:"Data"`
}

// common

type PaymentSt struct {
	QRToken                  string    `json:"QrToken"`
	ExpireDate               time.Time `json:"ExpireDate"`
	QrPaymentId              int64     `json:"QrPaymentId"`
	PaymentMethods           []string  `json:"PaymentMethods"`
	QrPaymentBehaviorOptions struct {
		StatusPollingInterval      int `json:"StatusPollingInterval"`
		QrCodeScanWaitTimeout      int `json:"QrCodeScanWaitTimeout"`
		PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
	} `json:"QrPaymentBehaviorOptions"`
}
