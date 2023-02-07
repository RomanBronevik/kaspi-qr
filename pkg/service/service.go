package service

import "kaspi-qr/pkg/repository"

type Authorization interface {
}

type Device interface {
	CreateDevice()
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
