package pg

import (
	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/adapters/logger"

	dopDb "github.com/rendau/dop/adapters/db"
)

type St struct {
	dopDb.RDBConnectionWithHelpers

	lg logger.WarnAndError
	db db.DB
}

func New(lg logger.WarnAndError, dopDb dopDb.RDBConnectionWithHelpers, db db.DB) *St {
	return &St{
		RDBConnectionWithHelpers: dopDb,
		lg:                       lg,
		db:                       db,
	}
}
