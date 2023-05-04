package core

import (
	"context"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
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

		c.StatusCheck()

		time.Sleep(7 * time.Second)
	}
}

func (c *PaymentCheck) StatusCheck() {
	defer dopTools.PanicRecover(c.r.lg, "statusCheck")

	ctx := context.Background()

	// get payments
	payments, err := c.r.Payment.List(ctx, &entities.PaymentListParsSt{
		Statuses: dopTools.NewSlicePtr(
			cns.PaymentStatusCreated,
			cns.PaymentStatusLinkActivated,
		),
	})
	if err != nil {
		return
	}

	for _, payment := range payments {
		if c.r.IsStopped() {
			return
		}

		c.r.wg.Add(1)
		c.StatusCheckForPayment(ctx, payment)
	}
}

func (c *PaymentCheck) StatusCheckForPayment(ctx context.Context, payment *entities.PaymentSt) {
	defer c.r.wg.Done()
	defer dopTools.PanicRecover(c.r.lg, "statusCheck")

	_, _ = c.r.Payment.GetStatus(ctx, payment.Id)
}
