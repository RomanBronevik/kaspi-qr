package usecases

import (
	"kaspi-qr/internal/domain/core"

	"github.com/rendau/dop/adapters/db"

	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg logger.Lite
	db db.RDBContextTransaction
	cr *core.St
}

func New(lg logger.Lite, db db.RDBContextTransaction, cr *core.St) *St {
	return &St{
		lg: lg,
		db: db,
		cr: cr,
	}
}
