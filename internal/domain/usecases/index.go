package usecases

import (
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/domain/core"
)

type St struct {
	lg logger.Lite
	db db.Transaction
	cr *core.St
}

func New(lg logger.Lite, db db.Transaction, cr *core.St) *St {
	return &St{
		lg: lg,
		db: db,
		cr: cr,
	}
}
