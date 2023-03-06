package usecases

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) SetMessageByStatusCode(statusCode int) string {
	output := s.cr.SetMessageByStatusCode(statusCode)

	return output
}

func (s *St) QrCreateOrderRecords(c *gin.Context, input entities.KaspiPaymentInput, output entities.QrTokenOutput) error {
	err := s.cr.QrCreateOrderRecords(c, input, output)

	return err
}

func (s *St) LinkCreateOrderRecords(c *gin.Context, input entities.KaspiPaymentInput, output entities.PaymentLinkRequestOutput) error {
	err := s.cr.LinkCreateOrderRecords(c, input, output)

	return err
}
