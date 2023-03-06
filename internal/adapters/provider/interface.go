package provider

import (
	"io"
	"kaspi-qr/internal/domain/entities"
)

type St interface {
	GetAllTradePoints(organizationBIN string) (entities.TradePointSt, error)
	DeviceRegistration(input entities.DeviceInputReg) (entities.DeviceOutputReg, error)
	DeviceDelete(input entities.DeviceInputDel) (entities.DeviceOutputDel, error)

	CreateQrToken(input entities.KaspiPaymentInput) (entities.QrTokenOutput, error)
	CreatePaymentLink(input entities.KaspiPaymentInput) (entities.PaymentLinkRequestOutput, error)
	OperationStatus(QrPaymentId string) (entities.OperationStatus, error)

	KaspiOperationDetails(requestBody io.ReadCloser) (entities.OperationDetails, error)
	KaspiReturnWithoutClient(input entities.ReturnRequestInput) (entities.ReturnSt, error)
}
