package tests

import (
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"
)

var (
	app = struct {
		repo  *pg.St
		cr    *core.St
		usc   *usecases.St
		kaspi *kaspi.St
	}{}
)
