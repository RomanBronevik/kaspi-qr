package old

import (
	"kaspi-qr/internal/domain/entities"
)

func (u *St) GetAllTradePoints(organizationBin string) (entities.TradePointSt, error) {
	tradePoints, err := u.cr.GetAllTradePoints(organizationBin)

	return tradePoints, err
}

func (u *St) DeviceRegistration(input entities.DeviceInputReg) (entities.DeviceOutputReg, error) {
	output, err := u.cr.DeviceRegistration(input)

	return output, err
}

func (u *St) DeleteOrOffDevice(input entities.DeviceInputDel) (entities.DeviceOutputDel, error) {
	output, err := u.cr.DeleteOrOffDevice(input)

	return output, err
}
