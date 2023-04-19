package rest

import (
	"errors"
	"fmt"
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/adapters/server"
	"kaspi-qr/internal/domain/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func Error(c *gin.Context, err error) bool {
	if err != nil {
		_ = c.Error(err)
		return true
	}
	return false
}

func BindJSON(c *gin.Context, obj any) bool {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		Error(c, errs.ErrWithDesc{
			Err:  errs.BadJson,
			Desc: err.Error(),
		})

		return false
	}

	return true
}

func BindQuery(c *gin.Context, obj any) bool {
	err := c.ShouldBindQuery(obj)
	if err != nil {
		Error(c, errs.ErrWithDesc{
			Err:  errs.BadQueryParams,
			Desc: err.Error(),
		})

		return false
	}

	return true
}

func MwRecovery(lg logger.WarnAndError, handler func(*gin.Context, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var err error

			if gErr := c.Errors.Last(); gErr != nil { // gin error
				if gErr.IsType(gin.ErrorTypeBind) {
					err = errs.ErrWithDesc{
						Err:  errs.BadJson,
						Desc: err.Error(),
					}
				} else {
					err = gErr.Err
				}
			} else if recoverRep := recover(); recoverRep != nil { // recovery error
				var ok bool
				if err, ok = recoverRep.(error); !ok {
					err = errors.New(fmt.Sprint(recoverRep))
				}
			}

			if err == nil {
				return
			}

			if handler != nil {
				handler(c, err)
				return
			}

			switch cErr := err.(type) {
			case errs.Err:
				c.AbortWithStatusJSON(http.StatusBadRequest, server.ErrRep{
					ErrorCode: cErr.Error(),
				})
			case errs.ErrWithDesc:
				c.AbortWithStatusJSON(http.StatusBadRequest, server.ErrRep{
					ErrorCode: cErr.Err.Error(),
					Desc:      cErr.Desc,
				})
			default:
				lg.Errorw(
					"Error in httpc handler",
					err,
					"method", c.Request.Method,
					"path", c.Request.URL.String(),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}

func MwCors() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           604800,
	})
}
