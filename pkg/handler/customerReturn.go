package handler

import "github.com/gin-gonic/gin"

func (h *Handler) details(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspiOperationDetails(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, output)
}

func (h *Handler) selfReturn(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspiReturnWithoutClient(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, output)
}
