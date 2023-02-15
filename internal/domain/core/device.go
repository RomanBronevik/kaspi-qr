package core

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateDevice(ctx *gin.Context, obj *entities.CreateDeviceDTO) error {
	err := s.repo.CreateDevice(ctx, obj)

	return err
}

func (s *St) DeleteDevice(ctx *gin.Context, bin string, token string) error {
	err := s.repo.DeleteDevice(ctx, bin, token)

	return err
}

//
//func (s *St) UpdateDevice(ctx *gin.Context, obj *entities.Device) error {
//	err := s.repo.UpdateDevice(ctx, obj.Token)
//
//	return err
//}

func (s *St) FindAllDevices(ctx *gin.Context) ([]entities.Device, error) {
	err, devices := s.repo.FindAllDevices(ctx)

	return err, devices
}

func (s *St) FindOneDevice(ctx *gin.Context, obj *entities.Device) (entities.Device, error) {
	device, err := s.repo.FindOneDevice(ctx, obj.DeviceId)

	return device, err
}
