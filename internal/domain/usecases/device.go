package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (u *St) CreateDevice(ctx *gin.Context, obj *entities.CreateDeviceDTO) error {
	err := u.cr.CreateDevice(ctx, obj)

	return err
}

//	func (s *St) UpdateDevice(ctx *gin.Context, obj *entities.Device) error {
//		err := s.cr.UpdateDevice(ctx, obj)
//
//		return err
//	}
func (u *St) DeleteDevice(ctx *gin.Context, bin string, token string) error {
	err := u.cr.DeleteDevice(ctx, bin, token)

	return err
}

func (u *St) FindAllDevices(ctx *gin.Context) ([]entities.Device, error) {
	devices, err := u.cr.FindAllDevices(ctx)

	return devices, err
}

func (u *St) FindOneDevice(ctx *gin.Context, OrganizationBin string) (entities.Device, error) {
	device, err := u.cr.FindOneDevice(ctx, OrganizationBin)

	return device, err
}

func (u *St) CreateDeviceRecord(ctx context.Context, input entities.DeviceInputReg, output entities.DeviceOutputReg) error {
	err := u.cr.CreateDeviceRecord(ctx, input, output)

	return err
}
