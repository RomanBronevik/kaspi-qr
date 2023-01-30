package handler

import "github.com/gin-gonic/gin"

func (h *Handler) QR(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspiQR(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, output)
}

func (h *Handler) paymentLink(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspiPaymentLink(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, output)
}

func (h *Handler) operationStatus(c *gin.Context) {
	QrPaymentId := c.Param("QrPaymentId")

	req, err := operationStatus(QrPaymentId)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	c.JSON(200, req)
}
