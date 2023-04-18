package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
)

// @Router   /organisation [get]
// @Tags     organisation
// @Param    query  query  entities.OrganisationListParsSt  false  "query"
// @Produce  json
// @Success  200  {object}  dopTypes.PaginatedListRep{results=[]entities.OrganisationSt}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hOrganisationList(c *gin.Context) {
	pars := &entities.OrganisationListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, tCount, err := o.ucs.OrganisationList(o.getRequestContext(c), pars)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.PaginatedListRep{
		Page:       pars.Page,
		PageSize:   pars.PageSize,
		TotalCount: tCount,
		Results:    result,
	})
}

// @Router   /organisation [post]
// @Tags     organisation
// @Param    body  body  entities.OrganisationCUSt  false  "body"
// @Success  200  {object} dopTypes.CreateRep{id=string}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hOrganisationCreate(c *gin.Context) {
	reqObj := &entities.OrganisationCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.OrganisationCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router   /organisation/:id [get]
// @Tags     organisation
// @Param    id path string true "id"
// @Produce  json
// @Success  200  {object}  entities.OrganisationSt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hOrganisationGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.OrganisationGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /organisation/:id [put]
// @Tags     organisation
// @Param    id path string true "id"
// @Param    body  body  entities.OrganisationCUSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hOrganisationUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.OrganisationCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.OrganisationUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router   /organisation/:id [delete]
// @Tags     organisation
// @Param    id path string true "id"
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hOrganisationDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.OrganisationDelete(o.getRequestContext(c), id))
}
