package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"kaspi-qr/internal/domain/entities"
	"time"
)

func (r *St) CreatePayment(ctx context.Context, payment *entities.CreatePaymentDTO) error {
	q := `
		INSERT INTO payment (created, modified, status, order_number, payment_id, payment_method, wait_timeout, polling_interval, payment_confirmation_timeout, amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	if err := r.db.Exec(ctx, q, payment.Created, payment.Modified, payment.Status, payment.OrderNumber, payment.PaymentId, payment.PaymentMethod, payment.WaitTimeout, payment.PollingInterval, payment.PaymentConfirmationTimeout, payment.Amount); err != nil {
		return r.ErorrHandler(err)
	}

	return nil
}

func (r *St) FindAllPayments(ctx context.Context) (u []entities.Payment, err error) {
	q := `
		SELECT created, modified, status, order_number, payment_id, payment_method, wait_timeout, polling_interval, payment_confirmation_timeout, amount FROM public.payment`
	rows, err := r.db.Query(ctx, q)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	payments := make([]entities.Payment, 0)

	for rows.Next() {
		payment := entities.Payment{}

		err := rows.Scan(&payment.Created, &payment.Modified, &payment.Status, &payment.OrderNumber, &payment.PaymentId, &payment.PaymentMethod, &payment.WaitTimeout, &payment.PollingInterval, &payment.PaymentConfirmationTimeout, &payment.Amount)
		if err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *St) FindOnePaymentByPaymentId(ctx context.Context, paymentId string) (entities.Payment, error) {
	q := `
		SELECT created, modified, status, order_number, payment_id, payment_method, wait_timeout, polling_interval, payment_confirmation_timeout, amount FROM payment WHERE payment_id = $1`

	//Trace

	var payment entities.Payment
	err := r.db.QueryRow(ctx, q, paymentId).Scan(&payment.Created, &payment.Modified, &payment.Status, &payment.OrderNumber, &payment.PaymentId, &payment.PaymentMethod, &payment.WaitTimeout, &payment.PollingInterval, &payment.PaymentConfirmationTimeout, &payment.Amount)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Payment{}, err
	}

	return payment, nil
}

func (r *St) FindLastPaymentByDesc(ctx context.Context, orderNumber string) (entities.Payment, error) {
	q := `
		SELECT created, modified, Status, order_number, payment_id, payment_method, wait_timeout, polling_interval, payment_confirmation_timeout, amount 
		FROM payment 
		WHERE order_number = $1
		ORDER BY wait_timeout DESC
		LIMIT 1`

	//Trace

	var payment entities.Payment
	err := r.db.QueryRow(ctx, q, orderNumber).Scan(&payment.Created, &payment.Modified, &payment.Status, &payment.OrderNumber, &payment.PaymentId, &payment.PaymentMethod, &payment.WaitTimeout, &payment.PollingInterval, &payment.PaymentConfirmationTimeout, &payment.Amount)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Payment{}, err
	}

	return payment, nil
}

func (r *St) UpdatePaymentRecordsToFail(ctx context.Context, orderNumber string) error {
	q := `
		UPDATE payment SET status = 'Error', modified = $1 
		               WHERE order_number = $2 and (status = 'Created' OR status = 'Wait');`

	if err := r.db.Exec(ctx, q, time.Now().Local(), orderNumber); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r *St) UpdatePaymentStatus(ctx context.Context, paymentId string, status string) error {
	q := `
		UPDATE payment SET status = $1, modified = $2 
		               WHERE payment_id = $3;`

	if err := r.db.Exec(ctx, q, status, time.Now().Local(), paymentId); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r *St) DeletePayment(ctx context.Context, orderNumber string) error {
	q := `
		DELETE FROM payment
		WHERE order_number = $1;`

	if err := r.db.Exec(ctx, q, orderNumber); err != nil {
		return r.ErorrHandler(err)
	}

	return nil
}

func (r *St) FindOnePaymentByOrderNumber(ctx context.Context, orderNumber string) (entities.Payment, error) {
	q := `
		SELECT created, modified, status, order_number, payment_id, payment_method, wait_timeout, polling_interval, payment_confirmation_timeout, amount FROM payment WHERE order_number = $1`
	var payment entities.Payment
	err := r.db.QueryRow(ctx, q, orderNumber).Scan(&payment.Created, &payment.Modified, &payment.Status, &payment.OrderNumber, &payment.PaymentId, &payment.PaymentMethod, &payment.WaitTimeout, &payment.PollingInterval, &payment.PaymentConfirmationTimeout, &payment.Amount)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Payment{}, err
	}

	return payment, nil
}
