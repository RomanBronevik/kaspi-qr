package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router   /payment [get]
// @Tags     payment
// @Param    query  query  entities.PaymentListParsSt  false  "query"
// @Produce  json
// @Success  200  {object}  dopTypes.PaginatedListRep{results=[]entities.PaymentSt}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPaymentList(c *gin.Context) {
	pars := &entities.PaymentListParsSt{}
	if !BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.PaymentList(o.getRequestContext(c), pars)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /payment/:id [get]
// @Tags     payment
// @Param    id path integer true "id"
// @Produce  json
// @Success  200  {object}  entities.PaymentSt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPaymentGet(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	result, err := o.ucs.PaymentGet(o.getRequestContext(c), id)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /payment/:id [put]
// @Tags     payment
// @Param    id path integer true "id"
// @Param    body  body  entities.PaymentCUSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPaymentUpdate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	reqObj := &entities.PaymentCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	Error(c, o.ucs.PaymentUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router   /payment/:id [delete]
// @Tags     payment
// @Param    id path integer true "id"
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPaymentDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	Error(c, o.ucs.PaymentDelete(o.getRequestContext(c), id))
}
