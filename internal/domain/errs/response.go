package errs

import (
	"github.com/gin-gonic/gin"
	"log"
)

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorSt{message})
}
