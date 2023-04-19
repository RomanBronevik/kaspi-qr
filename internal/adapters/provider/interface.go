package provider

type St interface {
	GetAllTradePoints(orgBin string) ([]*TradePointSt, error)
	DeviceRegistration(input DeviceCreateReqSt) (string, error)
	DeviceDelete(input DeviceRemoveReqSt) error

	CreateQrToken(input PaymentCreateReqSt) (*PaymentSt, error)
	CreatePaymentLink(input PaymentLinkCreateReqSt) (*PaymentLinkSt, error)
	OperationStatus(QrPaymentId string) (string, error)

	KaspiOperationDetails(input OperationGetReqSt) (*OperationDetailsSt, error)
	KaspiReturnWithoutClient(input ReturnReqSt) (int64, error)
}
