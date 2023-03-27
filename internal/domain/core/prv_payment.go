package core

import (
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateQrToken(input entities.KaspiPaymentInput) (entities.QrTokenOutput, error) {
	output, err := s.kaspi.CreateQrToken(input)

	return output, err
}

func (s *St) CreatePaymentLink(input entities.KaspiPaymentInput) (entities.PaymentLinkRequestOutput, error) {
	output, err := s.kaspi.CreatePaymentLink(input)

	return output, err
}

func (s *St) OperationStatus(QrPaymentId string) (entities.OperationStatus, error) {
	output, err := s.kaspi.OperationStatus(QrPaymentId)

	return output, err
}
