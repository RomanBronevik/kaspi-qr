package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router   /city [get]
// @Tags     city
// @Param    query  query  entities.CityListParsSt  false  "query"
// @Produce  json
// @Success  200  {object}  dopTypes.PaginatedListRep{results=[]entities.CitySt}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hCityList(c *gin.Context) {
	pars := &entities.CityListParsSt{}
	if !BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.CityList(o.getRequestContext(c), pars)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /city [post]
// @Tags     city
// @Param    body  body  entities.CityCUSt  false  "body"
// @Success  200  {object}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hCityCreate(c *gin.Context) {
	reqObj := &entities.CityCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.CityCreate(o.getRequestContext(c), reqObj)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": result})
}

// @Router   /city/:id [get]
// @Tags     city
// @Param    id path string true "id"
// @Produce  json
// @Success  200  {object}  entities.CitySt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hCityGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.CityGet(o.getRequestContext(c), id)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /city/:id [put]
// @Tags     city
// @Param    id path string true "id"
// @Param    body  body  entities.CityCUSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hCityUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.CityCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	Error(c, o.ucs.CityUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router   /city/:id [delete]
// @Tags     city
// @Param    id path string true "id"
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hCityDelete(c *gin.Context) {
	id := c.Param("id")

	Error(c, o.ucs.CityDelete(o.getRequestContext(c), id))
}
