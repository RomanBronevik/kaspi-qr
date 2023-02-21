package core

import (
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

func (s *St) FindOneCity(ctx *gin.Context, CityName string) (entities.City, error) {
	city, err := s.repo.FindOneCity(ctx, CityName)

	return city, err
}
