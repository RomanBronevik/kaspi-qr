package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
)

// @Router		/ord [get]
// @Tags		ord
// @Param		query	query	entities.OrdListParsSt	false	"query"
// @Produce	json
// @Success	200	{object}	dopTypes.PaginatedListRep{results=[]entities.OrdSt}
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hOrdList(c *gin.Context) {
	pars := &entities.OrdListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.OrdList(o.getRequestContext(c), pars)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/ord [post]
// @Tags		ord
// @Param		body	body		entities.OrdCUSt	false	"body"
// @Success	200		{object}	entities.OrdCreateRepSt
// @Failure	400		{object}	dopTypes.ErrRep
func (o *St) hOrdCreate(c *gin.Context) {
	reqObj := &entities.OrdCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.OrdCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/ord/:id [get]
// @Tags		ord
// @Param		id	path	string	true	"id"
// @Produce	json
// @Success	200	{object}	entities.OrdSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hOrdGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.OrdGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/ord/:id [put]
// @Tags		ord
// @Param		id		path	string				true	"id"
// @Param		body	body	entities.OrdCUSt	false	"body"
// @Produce	json
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hOrdUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.OrdCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.OrdUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router		/ord/:id [delete]
// @Tags		ord
// @Param		id	path	string	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hOrdDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.OrdDelete(o.getRequestContext(c), id))
}
