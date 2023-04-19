package provider

import (
	"time"
)

// request

type PaymentLinkCreateReqSt struct {
	PaymentCreateReqSt
}

// reply

type PaymentLinkCreateRepSt struct {
	BaseRepSt
	Data *PaymentLinkSt `json:"Data"`
}

// common

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
