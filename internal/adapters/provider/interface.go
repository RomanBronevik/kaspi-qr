package provider

type Provider interface {
	TradePointList(orgBin string) ([]*TradePointSt, error)

	DeviceCreate(reqObj DeviceCreateReqSt) (string, error)
	DeviceDelete(reqObj DeviceRemoveReqSt) error

	PaymentCreate(reqObj PaymentCreateReqSt) (*PaymentSt, error)
	PaymentLinkCreate(reqObj PaymentLinkCreateReqSt) (*PaymentLinkSt, error)
	PaymentGetStatus(qrPaymentId string) (string, error)
	PaymentGetDetails(paymentId int64, deviceToken string) (*PaymentDetailsSt, error)
	PaymentReturn(reqObj PaymentReturnReqSt) (int64, error)
}
