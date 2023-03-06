package usecases

import (
	"kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/domain/core"
)

type St struct {
	db *pg.St
	cr *core.St
}

func New(db *pg.St) *St {
	return &St{
		db: db,
	}
}

func (u *St) SetCore(core *core.St) {
	u.cr = core
}
