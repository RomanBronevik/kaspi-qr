package pg

import (
	dopDb "github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	dopDb.RDBConnectionWithHelpers

	lg logger.Lite
}

func New(lg logger.Lite, dopDb dopDb.RDBConnectionWithHelpers) *St {
	return &St{
		RDBConnectionWithHelpers: dopDb,
		lg:                       lg,
	}
}
