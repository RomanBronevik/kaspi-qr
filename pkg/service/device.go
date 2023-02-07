package service

import "kaspi-qr/pkg/repository"

type DeviceService struct {
	repo repository.Repository
}

func newDeviceService(repo repository.Repository) *DeviceService {
	return &DeviceService{repo: repo}
}
