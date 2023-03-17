package kaspi

import (
	"kaspi-qr/internal/adapters/logger"
)

type St struct {
	lg logger.Full
}

func New(lg logger.Full) *St {
	return &St{
		lg: lg,
	}
}
