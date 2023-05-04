package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
)

// @Router		/src [get]
// @Tags		src
// @Param		query	query	entities.SrcListParsSt	false	"query"
// @Produce	json
// @Success	200	{array}	entities.SrcSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hSrcList(c *gin.Context) {
	pars := &entities.SrcListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, _, err := o.ucs.SrcList(o.getRequestContext(c), pars)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/src [post]
// @Tags		src
// @Param		body	body		entities.SrcCUSt	false	"body"
// @Success	200		{object}	dopTypes.CreateRep{id=string}
// @Failure	400		{object}	dopTypes.ErrRep
func (o *St) hSrcCreate(c *gin.Context) {
	reqObj := &entities.SrcCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.SrcCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router		/src/:id [get]
// @Tags		src
// @Param		id	path	string	true	"id"
// @Produce	json
// @Success	200	{object}	entities.SrcSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hSrcGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.SrcGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/src/:id [put]
// @Tags		src
// @Param		id		path	string				true	"id"
// @Param		body	body	entities.SrcCUSt	false	"body"
// @Produce	json
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hSrcUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.SrcCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.SrcUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router		/src/:id [delete]
// @Tags		src
// @Param		id	path	string	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hSrcDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.SrcDelete(o.getRequestContext(c), id))
}
