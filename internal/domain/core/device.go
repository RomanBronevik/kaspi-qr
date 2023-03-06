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

func (s *St) DeviceAlredyExist(ctx *gin.Context, token string) (bool, error) {
	device, err := s.repo.FindOneDevice(ctx, token)
	if err != nil {
		return false, err
	}

	empty := entities.Device{}

	if device == empty {
		return false, nil
	}

	return true, nil
}

func (s *St) CreateDeviceRecord(ctx *gin.Context, input entities.DeviceInputReg, output entities.DeviceOutputReg) error {
	exits, err := s.DeviceAlredyExist(ctx, output.Data.DeviceToken)

	if err != nil {
		return err
	}

	if exits {
		return nil
	}

	dtoSt := entities.CreateDeviceDTO{
		Token:           output.Data.DeviceToken,
		DeviceId:        input.DeviceId,
		OrganizationBin: input.OrganizationBin,
	}

	err = s.CreateDevice(ctx, &dtoSt)
	if err != nil {
		return err
	}

	return nil
}

//
//func (s *St) UpdateDevice(ctx *gin.Context, obj *entities.Device) error {
//	err := s.repo.UpdateDevice(ctx, obj.Token)
//
//	return err
//}

func (s *St) FindAllDevices(ctx *gin.Context) ([]entities.Device, error) {
	devices, err := s.repo.FindAllDevices(ctx)

	return devices, err
}

func (s *St) FindOneDevice(ctx *gin.Context, OrganizationBin string) (entities.Device, error) {
	device, err := s.repo.FindOneDevice(ctx, OrganizationBin)

	return device, err
}
