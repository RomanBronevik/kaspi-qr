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

	wg        sync.WaitGroup
	stopped   bool
	stoppedMu sync.RWMutex

	City       *City
	Device     *Device
	Ord        *Ord
	Payment    *Payment
	TradePoint *TradePoint
	Src        *Src
}

func New(
	lg logger.Lite,
	repo *pg.St,
	prv provider.Provider,
	notifier notifier.Notifier,
	qrUrlTemplate string,
) *St {
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

func (c *St) Start() {
	go c.Payment.StatusChecker()
}

func (c *St) IsStopped() bool {
	c.stoppedMu.RLock()
	defer c.stoppedMu.RUnlock()
	return c.stopped
}

func (c *St) StopAndWaitJobs() {
	c.stoppedMu.Lock()
	c.stopped = true
	c.stoppedMu.Unlock()

	c.wg.Wait()
}
