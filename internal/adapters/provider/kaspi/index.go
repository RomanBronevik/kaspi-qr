package kaspi

import (
	"kaspi-qr/internal/adapters/logger"
)

type St struct {
	lg       logger.Full
	kaspiUrl string
}

func New(lg logger.Full, kaspiUrl string) *St {
	return &St{
		lg:       lg,
		kaspiUrl: kaspiUrl,
	}
}
