package kaspi

import (
	"kaspi-qr/internal/adapters/logger"
)

type St struct {
	lg         logger.Full
	httpClient *httpClientSt
}

func New(lg logger.Full, kaspiUrl, certPath, certPassword string) (*St, error) {
	httpClient, err := newHttpClient(lg, kaspiUrl, certPath, certPassword)
	if err != nil {
		lg.Errorw("generateCert", "err", err)
		return nil, err
	}

	return &St{
		lg:         lg,
		httpClient: httpClient,
	}, nil
}
