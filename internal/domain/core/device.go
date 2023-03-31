package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateDevice(ctx context.Context, obj *entities.CreateDeviceDTO) error {
	err := s.repo.CreateDevice(ctx, obj)

	return err
}

func (s *St) DeleteDevice(ctx *gin.Context, bin string, token string) error {
	err := s.repo.DeleteDevice(ctx, bin, token)

	return err
}

func (s *St) DeviceAlredyExist(ctx context.Context, token string) (bool, error) {
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

func (s *St) CreateDeviceRecord(ctx context.Context, input entities.DeviceInputReg, output entities.DeviceOutputReg) error {
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

func (s *St) CreateDeviceTwoSystems(input entities.DeviceInputReg) (*entities.DeviceOutputReg, error) {
	output, err := s.DeviceRegistration(input)

	if err != nil {
		return nil, err
	}

	output.Message = s.SetMessageByStatusCode(output.StatusCode)

	if output.StatusCode == 0 {
		err = s.CreateDeviceRecord(context.Background(), input, output)
		if err != nil {
			return nil, err
		}
	}

	return &output, err
}
