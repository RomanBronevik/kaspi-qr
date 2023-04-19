package old

import (
	"kaspi-qr/internal/domain/entities"
)

func (u *St) CreateQrToken(input entities.KaspiPaymentInput) (entities.QrTokenOutput, error) {
	output, err := u.cr.CreateQrToken(input)

	return output, err
}

func (u *St) CreatePaymentLink(input entities.KaspiPaymentInput) (entities.PaymentLinkRequestOutput, error) {
	output, err := u.cr.CreatePaymentLink(input)

	return output, err
}

func (u *St) OperationStatus(QrPaymentId string) (entities.OperationStatus, error) {
	output, err := u.cr.OperationStatus(QrPaymentId)

	return output, err
}
