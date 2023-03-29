package kaspi

import (
	"kaspi-qr/internal/adapters/logger"
)

type St struct {
	lg           logger.Full
	kaspiUrl     string
	certPath     string
	certPassword string
}

func New(lg logger.Full, kaspiUrl string, certPath string, certPassword string) *St {
	return &St{
		lg:           lg,
		kaspiUrl:     kaspiUrl,
		certPath:     certPath,
		certPassword: certPassword,
	}
}
