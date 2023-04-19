package old

import (
	"kaspi-qr/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (u *St) CreateCity(ctx *gin.Context, obj *entities.CreateCityDTO) error {
	err := u.cr.CreateCity(ctx, obj)

	return err
}

func (u *St) DeleteCity(ctx *gin.Context, id string) error {
	err := u.cr.DeleteCity(ctx, id)

	return err
}

//
// func (s *St) UpdateDevice(ctx *gin.Context, obj *entities.Device) error {
//	err := s.cr.UpdateDevice(ctx, obj.Token)
//
//	return err
// }

func (u *St) FindAllCities(ctx *gin.Context) ([]entities.City, error) {
	cities, err := u.cr.FindAllCities(ctx)

	return cities, err
}

func (u *St) FindOneCityByCityCode(ctx *gin.Context, code string) (entities.City, error) {
	city, err := u.cr.FindOneCityByCityCode(ctx, code)

	return city, err
}

func (u *St) DeleteCities(ctx *gin.Context) error {
	err := u.cr.DeleteCities(ctx)

	return err
}

func (u *St) UpdateCities(ctx *gin.Context, output entities.CityUpdateReqOutput) error {
	err := u.cr.UpdateCities(ctx, output)

	return err
}

func (u *St) IsEmptyCity(city entities.City) bool {
	exist := u.cr.IsEmptyCity(city)

	return exist
}
