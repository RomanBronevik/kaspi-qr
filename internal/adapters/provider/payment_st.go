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

type PaymentLinkCreateReqSt struct {
	PaymentCreateReqSt
}

type PaymentReturnReqSt struct {
	DeviceToken     string  `json:"DeviceToken"`
	OrganizationBin string  `json:"OrganizationBin"`
	QrPaymentId     int     `json:"QrPaymentId"`
	Amount          float64 `json:"Amount"`
}

// reply

type PaymentCreateRepSt struct {
	BaseRepSt
	Data PaymentSt `json:"Data"`
}

type PaymentLinkCreateRepSt struct {
	BaseRepSt
	Data PaymentLinkSt `json:"Data"`
}

type PaymentStatusRepSt struct {
	BaseRepSt
	Data struct {
		Status string `json:"Status"`
	} `json:"Data"`
}

type PaymentDetailsRepSt struct {
	BaseRepSt
	Data PaymentDetailsSt `json:"Data"`
}

type PaymentReturnRepSt struct {
	BaseRepSt
	ReturnOperationDataSt struct {
		ReturnOperationId int64 `json:"ReturnOperationId"`
	} `json:"ReturnOperationDataSt"`
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

type PaymentLinkSt struct {
	PaymentLink            string    `json:"PaymentLink"`
	ExpireDate             time.Time `json:"ExpireDate"`
	PaymentId              int64     `json:"PaymentId"`
	PaymentMethods         []string  `json:"PaymentMethods"`
	PaymentBehaviorOptions struct {
		StatusPollingInterval      int `json:"StatusPollingInterval"`
		LinkActivationWaitTimeout  int `json:"LinkActivationWaitTimeout"`
		PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
	} `json:"PaymentBehaviorOptions"`
}

type PaymentDetailsSt struct {
	QrPaymentId           int64     `json:"QrPaymentId"`
	TotalAmount           float64   `json:"TotalAmount"`
	AvailableReturnAmount float64   `json:"AvailableReturnAmount"`
	TransactionDate       time.Time `json:"TransactionDate"`
}
