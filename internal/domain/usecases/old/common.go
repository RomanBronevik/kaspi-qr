package old

import (
	"kaspi-qr/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (u *St) SetMessageByStatusCode(statusCode int) string {
	output := u.cr.SetMessageByStatusCode(statusCode)

	return output
}

func (u *St) QrCreateOrderRecords(c *gin.Context, input entities.KaspiPaymentInput, output entities.QrTokenOutput) error {
	err := u.cr.QrCreateOrderRecords(c, input, output)

	return err
}

func (u *St) LinkCreateOrderRecords(c *gin.Context, input entities.KaspiPaymentInput, output entities.PaymentLinkRequestOutput) error {
	err := u.cr.LinkCreateOrderRecords(c, input, output)

	return err
}
