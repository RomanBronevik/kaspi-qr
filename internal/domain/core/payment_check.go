package core

import (
	"context"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"sync"
	"time"

	"github.com/rendau/dop/dopTools"
)

type PaymentCheck struct {
	r *St
}

func NewPaymentCheck(r *St) *PaymentCheck {
	return &PaymentCheck{r: r}
}

// jobs

func (c *PaymentCheck) StatusChecker() {
	// first time sleep
	time.Sleep(10 * time.Second)

	for {
		if c.r.IsStopped() {
			return
		}

		c.r.wg.Add(1)
		c.StatusCheck()

		time.Sleep(7 * time.Second)
	}
}

func (c *PaymentCheck) StatusCheck() {
	defer c.r.wg.Done()
	defer dopTools.PanicRecover(c.r.lg, "StatusCheck")

	const workerCount = 10

	// get payments
	payments, _, err := c.r.Payment.List(context.Background(), &entities.PaymentListParsSt{
		Statuses: dopTools.NewSlicePtr(
			cns.PaymentStatusCreated,
			cns.PaymentStatusLinkActivated,
		),
	})
	if err != nil {
		return
	}

	jobCh := make(chan *entities.PaymentSt, len(payments))
	wg := &sync.WaitGroup{}

	// start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go c.statusCheckRoutine(wg, jobCh)
	}

	for _, payment := range payments {
		jobCh <- payment
	}

	close(jobCh)

	wg.Wait()
}

func (c *PaymentCheck) statusCheckRoutine(wg *sync.WaitGroup, jobCh <-chan *entities.PaymentSt) {
	defer wg.Done()
	defer dopTools.PanicRecover(c.r.lg, "statusCheckRoutine")

	for payment := range jobCh {
		if c.r.IsStopped() {
			return
		}

		_, _ = c.r.Payment.GetStatus(context.Background(), payment.Id)
	}
}
