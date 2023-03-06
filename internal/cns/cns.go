package cns

import "time"

const (
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
)

const UnInteger = 1000

const (
	CreatedStatus   = "Created"
	WaitStatus      = "Wait"
	ProcessedStatus = "Processed"
	ErrorStatus     = "Error"
	RefundStatus    = "Refunded"
)

const (
	QrPayment   = "Qr"
	LinkPayment = "PaymentLink"
)

const HoursQuantity = 72
