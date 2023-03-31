package core

import (
	"kaspi-qr/internal/domain/entities"
)

func (s *St) GetAllTradePoints(organizationBin string) (entities.TradePointSt, error) {
	tradePoints, err := s.kaspi.GetAllTradePoints(organizationBin)

	return tradePoints, err
}

func (s *St) DeviceRegistration(input entities.DeviceInputReg) (entities.DeviceOutputReg, error) {
	output, err := s.kaspi.DeviceRegistration(input)

	return output, err
}

func (s *St) DeleteOrOffDevice(input entities.DeviceInputDel) (entities.DeviceOutputDel, error) {
	output, err := s.kaspi.DeviceDelete(input)

	return output, err
}
