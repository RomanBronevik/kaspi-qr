package core

import (
	"context"
	"kaspi-qr/internal/adapters/notifier"
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"strconv"
	"strings"
	"time"

	qrcode "github.com/skip2/go-qrcode"
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

func (c *Payment) GetLink(ctx context.Context, id int64) (string, error) {
	return c.r.repo.PaymentGetLink(ctx, id)
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
	newId, err := c.create(ctx, &entities.PaymentCUSt{
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

	newPayment, err := c.Get(ctx, newId, true)
	if err != nil {
		return nil, err
	}

	return newPayment, nil
}

func (c *Payment) GetQrPicture(ctx context.Context, id int64) ([]byte, error) {
	const size = 350

	// get payment
	link, err := c.GetLink(ctx, id)
	if err != nil {
		return nil, err
	}
	if link == "" {
		return nil, errs.ObjectNotFound
	}

	//link := "https://google.kz?asd=asd&zxc=asdaasd"

	png, err := qrcode.Encode(link, qrcode.Medium, size)
	if err != nil {
		c.r.lg.Errorw("qrcode.Encode", err)
		return nil, err
	}

	return png, nil
}

func (c *Payment) GetStatus(ctx context.Context, id int64) (string, error) {
	var err error

	//c.r.lg.Infow("GetStatus", "id", id)

	// get payment
	payment, err := c.Get(ctx, id, true)
	if err != nil {
		return "", err
	}

	now := time.Now()
	allowanceDuration := 10 * time.Second
	newStatus := ""

	if payment.ExpireDt.Add(allowanceDuration).Before(now) {
		newStatus = cns.PaymentStatusExpired
	} else if payment.Status == cns.PaymentStatusCreated &&
		payment.StatusChangedAt.Add((time.Duration(payment.Pbo.LinkActivationWaitTimeout)*time.Second)+allowanceDuration).Before(now) {
		newStatus = cns.PaymentStatusExpired
	} else if payment.Status == cns.PaymentStatusLinkActivated &&
		payment.StatusChangedAt.Add((time.Duration(payment.Pbo.PaymentConfirmationTimeout)*time.Second)+allowanceDuration).Before(now) {
		newStatus = cns.PaymentStatusExpired
	} else {
		//get status from provider
		newStatus, err = c.r.prv.PaymentGetStatus(payment.Id)
		if err != nil {
			return "", err
		}
	}

	if newStatus == "" {
		newStatus = payment.Status
	}

	if newStatus != payment.Status {
		// update payment
		err = c.Update(ctx, id, &entities.PaymentCUSt{
			Status: &newStatus,
		})
		if err != nil {
			return "", err
		}
	}

	return newStatus, nil
}

func (c *Payment) statusToOrdStatus(v string) string {
	switch v {
	case cns.PaymentStatusCreated:
		return cns.OrdStatusCreated
	case cns.PaymentStatusLinkActivated:
		return cns.OrdStatusCreated
	case cns.PaymentStatusPaid:
		return cns.OrdStatusPaid
	case cns.PaymentStatusError:
		return cns.OrdStatusError
	case cns.PaymentStatusExpired:
		return cns.OrdStatusExpired
	case cns.PaymentStatusRefunded:
		return cns.OrdStatusRefunded
	default:
		c.r.lg.Errorw("unknown payment status", nil, "status", v)
		return cns.OrdStatusError
	}
}

func (c *Payment) createQrUrl(id int64) string {
	return strings.ReplaceAll(c.r.qrUrlTemplate, "{id}", strconv.FormatInt(id, 10))
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

	if obj.Status != nil {
		payment, err := c.Get(ctx, id, true)
		if err != nil {
			return err
		}

		ordStatus := c.statusToOrdStatus(payment.Status)

		// update ord
		err = c.r.Ord.Update(ctx, payment.OrdId, &entities.OrdCUSt{
			Status: &ordStatus,
		})
		if err != nil {
			return err
		}

		ord, err := c.r.Ord.Get(ctx, payment.OrdId, true)
		if err != nil {
			return err
		}

		if ord.SrcId != "" {
			src, err := c.r.Src.Get(ctx, ord.SrcId, false)
			if err != nil {
				return err
			}
			if src != nil {
				// notify
				_ = c.r.notifier.NotifyOrderStatusChange(src.NotifyUrl, &notifier.OrderStatusChangeReqSt{
					OrdId:     ord.Id,
					PaymentId: payment.Id,
					Status:    ord.Status,
				})
			}
		}

		//c.r.lg.Infow("payment status changed", "id", id, "status", payment.Status)
	}

	return nil
}

func (c *Payment) Delete(ctx context.Context, id int64) error {
	return c.r.repo.PaymentDelete(ctx, id)
}

// for testing

func (c *Payment) EmuPaymentScan(ctx context.Context, id int64) error {
	return c.r.prv.EmuPaymentScan(id)
}

func (c *Payment) EmuPaymentScanError(ctx context.Context, id int64) error {
	return c.r.prv.EmuPaymentScanError(id)
}

func (c *Payment) EmuPaymentConfirm(ctx context.Context, id int64) error {
	return c.r.prv.EmuPaymentConfirm(id)
}

func (c *Payment) EmuPaymentConfirmError(ctx context.Context, id int64) error {
	return c.r.prv.EmuPaymentConfirmError(id)
}
