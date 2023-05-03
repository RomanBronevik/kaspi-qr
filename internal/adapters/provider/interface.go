package provider

type Provider interface {
	TradePointList(orgBin string) ([]*TradePointSt, error)

	DeviceCreate(reqObj DeviceCreateReqSt) (string, error)
	DeviceDelete(reqObj DeviceDeleteReqSt) error

	PaymentCreate(reqObj PaymentCreateReqSt) (*PaymentSt, error)
	PaymentLinkCreate(reqObj PaymentCreateReqSt) (*PaymentLinkSt, error)
	PaymentGetStatus(qrPaymentId int64) (string, error)
	PaymentGetDetails(paymentId int64, deviceToken string) (*PaymentDetailsSt, error)
	PaymentReturn(reqObj PaymentReturnReqSt) (int64, error)

	// for testing

	EmuPaymentScan(paymentId int64) error
	EmuPaymentScanError(paymentId int64) error
	EmuPaymentConfirm(paymentId int64) error
	EmuPaymentConfirmError(paymentId int64) error
}
