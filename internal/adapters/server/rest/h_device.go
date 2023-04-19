package rest

import (
	"kaspi-qr/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router   /device [get]
// @Tags     device
// @Param    query  query  entities.DeviceListParsSt  false  "query"
// @Produce  json
// @Success  200  {object}  dopTypes.PaginatedListRep{results=[]entities.DeviceSt}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hDeviceList(c *gin.Context) {
	pars := &entities.DeviceListParsSt{}
	if !BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.DeviceList(o.getRequestContext(c), pars)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /device [post]
// @Tags     device
// @Param    body  body  entities.DeviceCUSt  false  "body"
// @Success  200  {object}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hDeviceCreate(c *gin.Context) {
	reqObj := &entities.DeviceCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.DeviceCreate(o.getRequestContext(c), reqObj)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": result})
}

// @Router   /device/:id [get]
// @Tags     device
// @Param    id path string true "id"
// @Produce  json
// @Success  200  {object}  entities.DeviceSt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hDeviceGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.DeviceGet(o.getRequestContext(c), id)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /device/:id [put]
// @Tags     device
// @Param    id path string true "id"
// @Param    body  body  entities.DeviceCUSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hDeviceUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.DeviceCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	Error(c, o.ucs.DeviceUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router   /device/:id [delete]
// @Tags     device
// @Param    id path string true "id"
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hDeviceDelete(c *gin.Context) {
	id := c.Param("id")

	Error(c, o.ucs.DeviceDelete(o.getRequestContext(c), id))
}