package entities

import (
	"github.com/golang-sql/civil"
)

type Payment struct {
	ID            string  `json:"id"`
	OrderNumber   string  `json:"order_number"`
	PaymentMethod string  `json:"payment_method"`
	PaymentType   string  `json:"payment_type"`
	Amount        float64 `json:"amount"`
}

type CreatePaymentDTO struct {
	OrderNumber   string  `json:"order_number"`
	PaymentMethod string  `json:"payment_method"`
	PaymentType   string  `json:"payment_type"`
	Amount        float64 `json:"amount"`
}

type QRToken struct {
	StatusCode int       `json:"StatusCode"`
	Message    string    `json:"Message"`
	Data       *QRStruct `json:"Data"`
}

type QRStruct struct {
	QRToken                  string                    `json:"QrToken"`
	ExpireDate               civil.DateTime            `json:"ExpireDate"`
	QrPaymentId              int                       `json:"QrPaymentId"`
	PaymentMethods           []string                  `json:"PaymentMethods"`
	QrPaymentBehaviorOptions *QrPaymentBehaviorOptions `json:"QrPaymentBehaviorOptions"`
}

type QrPaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	QrCodeScanWaitTimeout      int `json:"QrCodeScanWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type PaymentLink struct {
	StatusCode int            `json:"StatusCode"`
	Message    string         `json:"Message"`
	Data       *PaymentLinkSt `json:"Data"`
}

type PaymentLinkSt struct {
	PaymentLink            string                  `json:"PaymentLink"`
	ExpireDate             civil.DateTime          `json:"ExpireDate"`
	PaymentId              int                     `json:"PaymentId"`
	PaymentMethods         []string                `json:"PaymentMethods"`
	PaymentBehaviorOptions *PaymentBehaviorOptions `json:"PaymentBehaviorOptions"`
}

type PaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	LinkActivationWaitTimeout  int `json:"LinkActivationWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type OperationStatus struct {
	StatusCode int       `json:"StatusCode"`
	Message    string    `json:"Message"`
	Data       *StatusSt `json:"Data"`
}

type StatusSt struct {
	Status string `json:"Status"`
}