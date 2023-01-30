package handler

import "github.com/golang-sql/civil"

type tradePointSt struct {
	StatusCode int           `json:"StatusCode"`
	Message    string        `json:"Message"`
	Data       []*tradePoint `json:"Data"`
}

type tradePoint struct {
	TradePointId   int    `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}

type RegistrationOutputSt struct {
	Data       DeviceToken `json:"Data"`
	StatusCode int         `json:"StatusCode"`
}

type DeviceToken struct {
	DeviceToken string `json:"DeviceToken"`
}

type DeleteOutputSt struct {
	StatusCode int `json:"StatusCode"`
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

type ReturnSt struct {
	StatusCode            int                    `json:"StatusCode"`
	Message               string                 `json:"Message"`
	ReturnOperationDataSt *ReturnOperationDataSt `json:"ReturnOperationDataSt"`
}

type ReturnOperationDataSt struct {
	ReturnOperationId int `json:"ReturnOperationId"`
}
