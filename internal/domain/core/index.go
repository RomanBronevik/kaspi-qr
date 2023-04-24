package core

import (
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/adapters/notifier"
	"kaspi-qr/internal/adapters/provider"
	"sync"

	"kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/adapters/repo/pg"
)

type St struct {
	lg            logger.Lite
	repo          repo.Repo
	prv           provider.Provider
	notifier      notifier.Notifier
	qrUrlTemplate string

	wg         sync.WaitGroup
	City       *City
	Device     *Device
	Ord        *Ord
	Payment    *Payment
	TradePoint *TradePoint
	Src        *Src
}

func New(lg logger.Lite, repo *pg.St, prv provider.Provider, notifier notifier.Notifier, qrUrlTemplate string) *St {
	c := &St{
		lg:            lg,
		repo:          repo,
		prv:           prv,
		notifier:      notifier,
		qrUrlTemplate: qrUrlTemplate,
	}

	c.City = NewCity(c)
	c.Device = NewDevice(c)
	c.Ord = NewOrd(c)
	c.Payment = NewPayment(c)
	c.TradePoint = NewTradePoint(c)
	c.Src = NewSrc(c)

	return c
}

func (s *St) WaitJobs() {
	s.wg.Wait()
}
