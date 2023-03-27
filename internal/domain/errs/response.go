package errs

import (
	"github.com/gin-gonic/gin"
)

func NewErrorResponse(c *gin.Context, statusCode int, err Err, desc string) {
	c.AbortWithStatusJSON(statusCode, ErrWithDesc{err, desc})
}
