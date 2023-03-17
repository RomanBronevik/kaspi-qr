package pg

import (
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/adapters/logger"
)

type St struct {
	lg logger.WarnAndError
	db db.DB
}

func New(lg logger.WarnAndError, db db.DB) *St {
	return &St{
		lg: lg,
		db: db,
	}
}
