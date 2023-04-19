package core

import (
	"sync"

	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/adapters/repo/pg"
)

type St struct {
	repo  repo.Repo
	kaspi *kaspi.St

	wg      sync.WaitGroup
	City    *City
	Device  *Device
	Ord     *Ord
	Payment *Payment
}

func New(repo *pg.St, kaspi *kaspi.St) *St {
	c := &St{
		repo:  repo,
		kaspi: kaspi,
	}

	c.City = NewCity(c)
	c.Device = NewDevice(c)
	c.Ord = NewOrd(c)
	c.Payment = NewPayment(c)

	return c
}

func (s *St) WaitJobs() {
	s.wg.Wait()
}
