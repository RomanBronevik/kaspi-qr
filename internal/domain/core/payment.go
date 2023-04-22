package core

import (
	"context"
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"time"
)

type Payment struct {
	r *St
}

func NewPayment(r *St) *Payment {
	return &Payment{r: r}
}

func (c *Payment) ValidateCU(ctx context.Context, obj *entities.PaymentCUSt, id int64) error {
	//forCreate := id == 0

	return nil
}

func (c *Payment) List(ctx context.Context, pars *entities.PaymentListParsSt) ([]*entities.PaymentSt, error) {
	items, err := c.r.repo.PaymentList(ctx, pars)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Payment) Get(ctx context.Context, id int64, errNE bool) (*entities.PaymentSt, error) {
	result, err := c.r.repo.PaymentGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, errs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *Payment) IdExists(ctx context.Context, id int64) (bool, error) {
	return c.r.repo.PaymentIdExists(ctx, id)
}

func (c *Payment) CreateForOrd(ctx context.Context, ordId string) (*entities.PaymentSt, error) {
	var err error

	paymentStatus := cns.PaymentStatusCreated

	// get ord
	ord, err := c.r.Ord.Get(ctx, ordId, false)
	if err != nil {
		return nil, err
	}
	if ord == nil {
		return nil, errs.OrderNotFound
	}

	// get device
	device, err := c.r.Device.Get(ctx, ord.DeviceId, false)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, errs.DeviceNotFound
	}

	// create prvPayment in provider
	prvPayment, err := c.r.prv.PaymentLinkCreate(provider.PaymentCreateReqSt{
		ExternalId:      ord.Id,
		Amount:          ord.Amount,
		OrganizationBin: device.OrgBin,
		DeviceToken:     device.Token,
	})
	if err != nil {
		return nil, err
	}

	// create payment
	_, err = c.create(ctx, &entities.PaymentCUSt{
		Id:       &prvPayment.PaymentId,
		OrdId:    &ordId,
		Link:     &prvPayment.PaymentLink,
		Status:   &paymentStatus,
		Amount:   &ord.Amount,
		ExpireDt: &prvPayment.ExpireDate,
		Pbo: &entities.PaymentPboSt{
			StatusPollingInterval:      prvPayment.PaymentBehaviorOptions.StatusPollingInterval,
			LinkActivationWaitTimeout:  prvPayment.PaymentBehaviorOptions.LinkActivationWaitTimeout,
			PaymentConfirmationTimeout: prvPayment.PaymentBehaviorOptions.PaymentConfirmationTimeout,
		},
	})
	if err != nil {
		return nil, err
	}

	newPayment, err := c.Get(ctx, prvPayment.PaymentId, true)
	if err != nil {
		return nil, err
	}

	return newPayment, nil
}

func (c *Payment) create(ctx context.Context, obj *entities.PaymentCUSt) (int64, error) {
	return c.r.repo.PaymentCreate(ctx, obj)
}

func (c *Payment) Update(ctx context.Context, id int64, obj *entities.PaymentCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	now := time.Now()

	obj.Modified = &now

	if obj.Status != nil {
		obj.StatusChangedAt = &now
	}

	err = c.r.repo.PaymentUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Payment) Delete(ctx context.Context, id int64) error {
	return c.r.repo.PaymentDelete(ctx, id)
}
