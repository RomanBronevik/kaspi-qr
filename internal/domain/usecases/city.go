package usecases

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateCity(ctx *gin.Context, obj *entities.CreateCityDTO) error {
	err := s.cr.CreateCity(ctx, obj)

	return err
}

func (s *St) DeleteCity(ctx *gin.Context, id string) error {
	err := s.cr.DeleteCity(ctx, id)

	return err
}

//
//func (s *St) UpdateDevice(ctx *gin.Context, obj *entities.Device) error {
//	err := s.cr.UpdateDevice(ctx, obj.Token)
//
//	return err
//}

func (s *St) FindAllCities(ctx *gin.Context) ([]entities.City, error) {
	cities, err := s.cr.FindAllCities(ctx)

	return cities, err
}

func (s *St) FindOneCityByCityCode(ctx *gin.Context, code string) (entities.City, error) {
	city, err := s.cr.FindOneCityByCityCode(ctx, code)

	return city, err
}

func (s *St) DeleteCities(ctx *gin.Context) error {
	err := s.cr.DeleteCities(ctx)

	return err
}

func (s *St) UpdateCities(ctx *gin.Context, output entities.CityUpdateReqOutput) error {
	err := s.cr.UpdateCities(ctx, output)

	return err
}

func (s *St) IsEmptyCity(city entities.City) bool {
	exist := s.cr.IsEmptyCity(city)

	return exist
}
