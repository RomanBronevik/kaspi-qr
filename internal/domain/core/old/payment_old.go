package old

import (
	"context"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *St) CreatePayment(ctx *gin.Context, obj *entities.CreatePaymentDTO) error {
	err := s.repo.CreatePayment(ctx, obj)

	return err
}

func (s *St) DeletePayment(ctx *gin.Context, orderNumber string) error {
	err := s.repo.DeletePayment(ctx, orderNumber)

	return err
}

func (s *St) FindAllPayments(ctx *gin.Context) ([]entities.Payment, error) {
	payments, err := s.repo.FindAllPayments(ctx)

	return payments, err
}

func (s *St) FindOnePaymentByOrderNumber(ctx *gin.Context, orderNumber string) (entities.Payment, error) {

	payment, err := s.repo.FindOnePaymentByOrderNumber(ctx, orderNumber)

	return payment, err
}

func (s *St) FindOnePaymentByPaymentId(ctx *gin.Context, paymentId string) (entities.Payment, error) {

	payment, err := s.repo.FindOnePaymentByPaymentId(ctx, paymentId)

	return payment, err
}

func (s *St) StatusPaid(status string) bool {
	if status == cns.StatusProcessed {
		return true
	}
	return false
}

func (s *St) StatusWait(status string) bool {
	if status == cns.StatusWait {
		return true
	}
	return false
}

func (s *St) StatusCreated(status string) bool {
	if status == cns.StatusCreated {
		return true
	}
	return false
}

func (s *St) StatusCanceled(status string) bool {
	if status == cns.StatusError {
		return true
	}
	return false
}

func (s *St) UpdatePaymentRecordsToFail(ctx context.Context, orderNumber string) error {
	err := s.repo.UpdatePaymentRecordsToFail(ctx, orderNumber)

	return err
}

func (s *St) PaymentExist(c *gin.Context, orderNumber string) bool {
	payment, _ := s.FindOnePaymentByOrderNumber(c, orderNumber)

	empty := entities.Payment{}
	if payment == empty {
		return false
	}

	return true
}

func (s *St) FindLastPaymentByDesc(c *gin.Context, orderNumber string) (entities.Payment, error) {
	payment, err := s.repo.FindLastPaymentByDesc(c, orderNumber)
	if err != nil {
		return entities.Payment{}, err
	}
	return payment, nil
}

func (s *St) CancelPreviousPayment(c *gin.Context, orderNumber string) error {
	if s.PaymentExist(c, orderNumber) {
		err := s.UpdatePaymentRecordsToFail(c, orderNumber)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *St) QrCreatePaymentRecord(c *gin.Context, currTime time.Time, input entities.KaspiPaymentInput, output entities.QrTokenOutput) error {

	err := s.CancelPreviousPayment(c, input.ExternalId)
	if err != nil {
		return err
	}

	paymentId := strconv.Itoa(output.Data.QrPaymentId)
	QrPaymentBehaviorOptions := output.Data.QrPaymentBehaviorOptions

	paymentDtoSt := entities.CreatePaymentDTO{
		Created:                    currTime,
		Modified:                   currTime,
		Status:                     cns.StatusCreated,
		OrderNumber:                input.ExternalId,
		PaymentId:                  paymentId,
		PaymentMethod:              cns.PaymentMethodQr,
		WaitTimeout:                currTime.Add(time.Second * time.Duration(QrPaymentBehaviorOptions.QrCodeScanWaitTimeout)),
		PollingInterval:            QrPaymentBehaviorOptions.StatusPollingInterval,
		PaymentConfirmationTimeout: QrPaymentBehaviorOptions.PaymentConfirmationTimeout,
		Amount:                     input.Amount,
	}

	err = s.CreatePayment(c, &paymentDtoSt)
	if err != nil {
		return err
	}

	return nil
}

func (s *St) LinkCreatePaymentRecord(c *gin.Context, currTime time.Time, input entities.KaspiPaymentInput, output entities.PaymentLinkRequestOutput) error {
	err := s.CancelPreviousPayment(c, input.ExternalId)
	if err != nil {
		return err
	}

	paymentId := strconv.Itoa(output.Data.PaymentId)
	PaymentBehaviorOptions := output.Data.PaymentBehaviorOptions

	paymentDtoSt := entities.CreatePaymentDTO{
		Created:                    currTime,
		Modified:                   currTime,
		Status:                     cns.StatusCreated,
		OrderNumber:                input.ExternalId,
		PaymentId:                  paymentId,
		PaymentMethod:              cns.PaymentMethodLink,
		WaitTimeout:                currTime.Add(time.Second * time.Duration(PaymentBehaviorOptions.LinkActivationWaitTimeout)),
		PollingInterval:            PaymentBehaviorOptions.StatusPollingInterval,
		PaymentConfirmationTimeout: PaymentBehaviorOptions.PaymentConfirmationTimeout,
		Amount:                     input.Amount,
	}

	err = s.CreatePayment(c, &paymentDtoSt)
	if err != nil {
		return err
	}

	return nil
}

func (s *St) CheckPaymentStatus(ctx *gin.Context) error {
	var err error

	var mutex = &sync.Mutex{}
	var isLocked bool

	orders, err := s.FindAllUnpaidOrders(ctx)

	if err != nil {
		return err
	}

	orderQuantity := len(orders)

	if orderQuantity == 0 {
		return nil
	}

	workerCount := findCh(orderQuantity)

	ch := make(chan error, 10)

	for i := 0; i < workerCount; i++ {
		s.checkUnpaid(ctx, mutex, isLocked, orders, i*cns.UnInteger, orderQuantity, ch)
	}

	return <-ch
}

func (s *St) checkUnpaid(ctx *gin.Context, mutex *sync.Mutex, isLocked bool, orders []entities.UnPaidOrder, indexI int, quantity int, doneCh chan error) {
	if isLocked {
		doneCh <- nil
		return
	}

	mutex.Lock()
	setLock(isLocked, true)

	defer mutex.Unlock()
	defer setLock(isLocked, false)

	i := indexI
	var err error
	for ; i < indexI+cns.UnInteger; i++ {
		if i >= quantity {
			break
		}
		value := orders[i]

		payment, err := s.FindOnePaymentByPaymentId(ctx, value.PaymentId)
		if err != nil {
			continue
		}

		active, err := s.checkOrderActivity(ctx, value.OrderNumber, value.Created)
		if err != nil {
			continue
		}

		if !active {
			continue
		}

		if s.StatusCreated(payment.Status) && payment.WaitTimeout.Before(time.Now().Local()) {
			err = s.UpdatePaymentStatus(ctx, value.PaymentId, cns.StatusError)
		}

		if s.StatusWait(payment.Status) {
			ConfirmationTimeout := payment.Modified.Add(time.Duration(payment.PaymentConfirmationTimeout) * time.Second)
			if time.Now().Local().After(ConfirmationTimeout) {
				err = s.UpdatePaymentStatus(ctx, value.PaymentId, cns.StatusError)
			}
		}

		operationStatusSt, err := s.kaspi.OperationStatus(value.PaymentId)
		if err != nil {
			continue
		}

		if operationStatusSt.StatusCode != 0 {
			continue
		}

		operationStatus := operationStatusSt.Data.Status

		if payment.Status == operationStatus {
			continue
		}

		switch operationStatus {
		case kaspi.PaymentStatusWait:
			err = s.UpdatePaymentStatus(ctx, value.PaymentId, cns.StatusWait)
		case kaspi.PaymentStatusProcessed:
			err = s.UpdatePaymentStatus(ctx, value.PaymentId, cns.StatusProcessed)
		case kaspi.PaymentStatusError:
			err = s.UpdatePaymentStatus(ctx, value.PaymentId, cns.StatusError)
		default:
			continue
		}

		if err != nil {
			continue
		}
	}

	doneCh <- err
}

func (s *St) checkOrderActivity(c *gin.Context, orderNumber string, created time.Time) (bool, error) {
	if time.Now().Local().After(created.Add(time.Duration(cns.HoursQuantity) * time.Hour)) {
		err := s.UpdateOrderStatus(c, orderNumber, cns.StatusError)
		if err != nil {
			return false, err
		}
		err = s.CancelPreviousPayment(c, orderNumber)
		if err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func (s *St) UpdatePaymentStatus(ctx context.Context, paymentId string, status string) error {
	err := s.repo.UpdatePaymentStatus(ctx, paymentId, status)

	return err
}

func setLock(isLocked bool, lock bool) {
	isLocked = lock
}
