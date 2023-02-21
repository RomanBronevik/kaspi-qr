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

func (s *St) FindOneCity(ctx *gin.Context, CityName string) (entities.City, error) {
	city, err := s.cr.FindOneCity(ctx, CityName)

	return city, err
}
