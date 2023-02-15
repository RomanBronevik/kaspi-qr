package service

import (
	"kaspi-qr/internal/adapters/repo/pg"
	//"kaspi-qr/pkg/repository"
)

type Authorization interface {
}

type Device interface {
	CreateDevice()
}

type Service struct {
	Authorization
}

func NewService(repos *pg.St) *Service {
	return &Service{}
}
