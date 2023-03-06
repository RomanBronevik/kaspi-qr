package errs

import (
	"github.com/gin-gonic/gin"
)

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorSt{message})
}
