package kaspi

const (
	SuccessStatus                         = 0
	NoCertificateStatus                   = -10000
	DeviceNotFoundStatus                  = -1501
	DeviceDeactivatedStatus               = -1502
	DeviceAlreadyExistStatus              = -1503
	PurchaseNotFoundStatus                = -1601
	TradePointsDoesntExistStatus          = -14000002
	TradePointNotFound                    = -99000002
	RefundAmountGreaterStatus             = -99000005
	RefundErrorStatus                     = -99000006
	TradePointDeactivatedStatus           = 990000018
	TradePointDoesntAcceptQrPaymentStatus = 990000026
	WrongAmountStatus                     = 990000028
	NoPaymentMethodsAvailableStatus       = 99000033
	PurchaseUuidNotFoundStatus            = -99000001
	TradePointDoesntMatchDeviceStatus     = -99000003
	WrongPurchaseStatus                   = -99000011
	PartialRefundNotAvailableStatus       = -99000020
	ServiceNotAvailableStatus             = -999
)
