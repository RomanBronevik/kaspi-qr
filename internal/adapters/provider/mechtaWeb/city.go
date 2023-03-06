package mechtaWeb

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func GetCitiesFromSite(c *gin.Context) entities.CityUpdateReqOutput {
	resp, err := http.Get(viper.GetString("cityWebPage"))

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return entities.CityUpdateReqOutput{}
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		}
	}()

	if resp.StatusCode != 200 {
		textError := fmt.Sprintf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		errs.NewErrorResponse(c, http.StatusBadRequest, textError)
		return entities.CityUpdateReqOutput{}
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return entities.CityUpdateReqOutput{}
	}

	var bodyRequest entities.CityUpdateReqOutput

	err = json.Unmarshal(bytes, &bodyRequest)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return entities.CityUpdateReqOutput{}
	}

	return bodyRequest
}
