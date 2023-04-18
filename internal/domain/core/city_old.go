package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateCity(ctx *gin.Context, obj *entities.CreateCityDTO) error {
	err := s.repo.CreateCity(ctx, obj)

	return err
}

func (s *St) DeleteCity(ctx *gin.Context, id string) error {
	err := s.repo.DeleteCity(ctx, id)

	return err
}

//
//func (s *St) UpdateDevice(ctx *gin.Context, obj *entities.Device) error {
//	err := s.repo.UpdateDevice(ctx, obj.Token)
//
//	return err
//}

func (s *St) FindAllCities(ctx *gin.Context) ([]entities.City, error) {
	cities, err := s.repo.FindAllCities(ctx)

	return cities, err
}

func (s *St) FindOneCityByCityCode(ctx *gin.Context, code string) (entities.City, error) {
	city, err := s.repo.FindOneCityByCityCode(ctx, code)

	return city, err
}

func (s *St) DeleteCities(ctx *gin.Context) error {
	err := s.repo.DeleteCities(ctx)

	return err
}

func (s *St) UpdateCities(ctx *gin.Context, output entities.CityUpdateReqOutput) error {
	err := s.DeleteCities(ctx)

	if err != nil {
		return err
	}

	if output.Data == nil {
		return errors.New("Data is null - cities update")
	}

	for _, val := range output.Data.Cities {
		newCity := entities.CreateCityDTO{
			Name:            val.Name,
			OrganizationBin: "160640004075",
			Code:            val.Code,
		}

		err = s.CreateCity(ctx, &newCity)
	}

	err = s.CreateTestCity(ctx)

	return err
}

func (s *St) CreateTestCity(ctx *gin.Context) error {
	newCity := entities.CreateCityDTO{
		Name:            "test",
		OrganizationBin: "160640004075",
		Code:            "test",
	}

	err := s.CreateCity(ctx, &newCity)
	return err
}

func (s *St) IsEmptyCity(city entities.City) bool {
	empty := entities.City{}

	if empty == city {
		return true
	}

	return false
}
