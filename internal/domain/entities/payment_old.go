package entities

import (
	"time"
)

type Payment struct {
	Created                    time.Time `json:"created"`
	Modified                   time.Time `json:"modified"`
	Status                     string    `json:"status"`
	OrderNumber                string    `json:"order_number"`
	PaymentId                  string    `json:"payment_id"`
	PaymentMethod              string    `json:"payment_method"`
	WaitTimeout                time.Time `json:"wait_timeout"`
	PollingInterval            int       `json:"polling_interval"`
	PaymentConfirmationTimeout int       `json:"payment_confirmation_timeout"`
	Amount                     float64   `json:"amount"`
}

type CreatePaymentDTO struct {
	Created                    time.Time `json:"created"`
	Modified                   time.Time `json:"modified"`
	Status                     string    `json:"status"`
	OrderNumber                string    `json:"order_number"`
	PaymentId                  string    `json:"payment_id"`
	PaymentMethod              string    `json:"payment_method"`
	WaitTimeout                time.Time `json:"wait_timeout"`
	PollingInterval            int       `json:"polling_interval"`
	PaymentConfirmationTimeout int       `json:"payment_confirmation_timeout"`
	Amount                     float64   `json:"amount"`
}

type KaspiPaymentInput struct {
	OrganizationBin string  `json:"OrganizationBin"`
	DeviceToken     string  `json:"DeviceToken"`
	Amount          float64 `json:"Amount"`
	ExternalId      string  `json:"ExternalId"`
}

type QrTokenOutput struct {
	StatusCode int       `json:"StatusCode"`
	Message    string    `json:"Message"`
	Data       *QRStruct `json:"Data"`
}

type QrTokenRequestInput struct {
	OrderNumber string  `json:"OrderNumber"`
	Amount      float64 `json:"Amount"`
	Code        string  `json:"CityCode"`
}

type QRStruct struct {
	QRToken                  string                    `json:"QrToken"`
	ExpireDate               time.Time                 `json:"ExpireDate"`
	QrPaymentId              int                       `json:"QrPaymentId"`
	PaymentMethods           []string                  `json:"PaymentMethods"`
	QrPaymentBehaviorOptions *QrPaymentBehaviorOptions `json:"QrPaymentBehaviorOptions"`
}

type QrPaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	QrCodeScanWaitTimeout      int `json:"QrCodeScanWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type PaymentLinkRequestKaspiInput struct {
	OrganizationBin string  `json:"OrganizationBin"`
	DeviceToken     string  `json:"DeviceToken"`
	Amount          float64 `json:"Amount"`
	ExternalId      string  `json:"ExternalId"`
}

type PaymentLinkRequestInput struct {
	OrderNumber string  `json:"OrderNumber"`
	Amount      float64 `json:"Amount"`
	Code        string  `json:"Code"`
}

type PaymentLinkRequestOutput struct {
	StatusCode int            `json:"StatusCode"`
	Message    string         `json:"Message"`
	Data       *PaymentLinkSt `json:"Data"`
}

type PaymentLinkSt struct {
	PaymentLink            string                  `json:"PaymentLink"`
	ExpireDate             time.Time               `json:"ExpireDate"`
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
