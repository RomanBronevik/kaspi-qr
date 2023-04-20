package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router   /trade_point [get]
// @Tags     trade_point
// @Param    query  query  entities.TradePointListParsSt  false  "query"
// @Produce  json
// @Success  200  {array}  entities.TradePointSt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hTradePointList(c *gin.Context) {
	pars := &entities.TradePointListParsSt{}
	if !BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.TradePointList(o.getRequestContext(c), pars)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}
