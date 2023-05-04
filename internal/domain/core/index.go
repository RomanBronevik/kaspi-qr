package core

import (
	"kaspi-qr/internal/adapters/notifier"
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/adapters/repo/pg"
	"sync"

	"github.com/rendau/dop/adapters/logger"
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

	City         *City
	Device       *Device
	Ord          *Ord
	Payment      *Payment
	PaymentCheck *PaymentCheck
	TradePoint   *TradePoint
	Src          *Src
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
	c.PaymentCheck = NewPaymentCheck(c)
	c.TradePoint = NewTradePoint(c)
	c.Src = NewSrc(c)

	return c
}

func (c *St) Start() {
	go c.PaymentCheck.StatusChecker()
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
