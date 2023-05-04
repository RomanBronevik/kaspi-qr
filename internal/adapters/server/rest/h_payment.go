package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
)

// @Router		/payment [get]
// @Tags		payment
// @Param		query	query	entities.PaymentListParsSt	false	"query"
// @Produce	json
// @Success	200	{object}	dopTypes.PaginatedListRep{results=[]entities.PaymentSt}
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentList(c *gin.Context) {
	pars := &entities.PaymentListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.PaymentList(o.getRequestContext(c), pars)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/payment/:id [get]
// @Tags		payment
// @Param		id	path	integer	true	"id"
// @Produce	json
// @Success	200	{object}	entities.PaymentSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentGet(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	result, err := o.ucs.PaymentGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/payment/:id [put]
// @Tags		payment
// @Param		id		path	integer					true	"id"
// @Param		body	body	entities.PaymentCUSt	false	"body"
// @Produce	json
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentUpdate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	reqObj := &entities.PaymentCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.PaymentUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router		/payment/:id [delete]
// @Tags		payment
// @Param		id	path	integer	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	dopHttps.Error(c, o.ucs.PaymentDelete(o.getRequestContext(c), id))
}

// @Router		/payment/:id/qr_picture [get]
// @Tags		payment
// @Param		id	path	integer	true	"id"
// @Produce	octet-stream
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentGetQrPicture(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	png, err := o.ucs.PaymentGetQrPicture(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}

// EMU

// @Router		/emu/payment/:id/scan [post]
// @Tags		emu
// @Param		id	path	integer	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentEmuPaymentScan(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	dopHttps.Error(c, o.ucs.PaymentEmuPaymentScan(o.getRequestContext(c), id))
}

// @Router		/emu/payment/:id/scan_error [post]
// @Tags		emu
// @Param		id	path	integer	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentEmuPaymentScanError(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	dopHttps.Error(c, o.ucs.PaymentEmuPaymentScanError(o.getRequestContext(c), id))
}

// @Router		/emu/payment/:id/confirm [post]
// @Tags		emu
// @Param		id	path	integer	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentEmuPaymentConfirm(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	dopHttps.Error(c, o.ucs.PaymentEmuPaymentConfirm(o.getRequestContext(c), id))
}

// @Router		/emu/payment/:id/confirm_error [post]
// @Tags		emu
// @Param		id	path	integer	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hPaymentEmuPaymentConfirmError(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	dopHttps.Error(c, o.ucs.PaymentEmuPaymentConfirmError(o.getRequestContext(c), id))
}
