package core

import (
	"kaspi-qr/internal/adapters/provider"
	"sync"

	"kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/adapters/repo/pg"
)

type St struct {
	repo repo.Repo
	prv  provider.Provider

	wg         sync.WaitGroup
	City       *City
	Device     *Device
	Ord        *Ord
	Payment    *Payment
	TradePoint *TradePoint
}

func New(repo *pg.St, prv provider.Provider) *St {
	c := &St{
		repo: repo,
		prv:  prv,
	}

	c.City = NewCity(c)
	c.Device = NewDevice(c)
	c.Ord = NewOrd(c)
	c.Payment = NewPayment(c)
	c.TradePoint = NewTradePoint(c)

	return c
}

func (s *St) WaitJobs() {
	s.wg.Wait()
}
