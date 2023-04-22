package kaspi

import "time"

const (
	RequestTimeout = 15 * time.Second

	StatusSuccess                         = 0
	StatusNoCertificate                   = -10000
	StatusDeviceNotFound                  = -1501
	StatusDeviceDeactivated               = -1502
	StatusDeviceAlreadyExist              = -1503
	StatusPurchaseNotFound                = -1601
	StatusTradePointsDoesntExist          = -14000002
	tFoundTradePointNo                    = -99000002
	StatusRefundAmountGreater             = -99000005
	StatusRefundError                     = -99000006
	StatusTradePointDeactivated           = 990000018
	StatusTradePointDoesntAcceptQrPayment = 990000026
	StatusWrongAmount                     = 990000028
	StatusNoPaymentMethodsAvailable       = 99000033
	StatusPurchaseUuidNotFound            = -99000001
	StatusTradePointDoesntMatchDevice     = -99000003
	StatusWrongPurchase                   = -99000011
	StatusPartialRefundNotAvailable       = -99000020
	StatusServiceNotAvailable             = -999
)

const (
	PaymentStatusQrTokenCreated = "QrTokenCreated"
	PaymentStatusWait           = "Wait"
	PaymentStatusProcessed      = "Processed"
	PaymentStatusError          = "Error"
)
