package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreatePayment(ctx context.Context, payment *entities.CreatePaymentDTO) error {
	q := `
		INSERT INTO payment (payment_method, payment_type, order_number, amount)
		VALUES ($1, $2)
		RETURNING id`

	var id string

	if err := r.client.QueryRow(ctx, q, payment.PaymentMethod, payment.PaymentType, payment.OrderNumber).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *St) FindAllPayments(ctx context.Context) (u []entities.Payment, err error) {
	q := `
		SELECT ID, ORDER_NUMBER, PAYMENT_METHOD, PAYMENT_TYPE, AMOUNT FROM public.payment`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments := make([]entities.Payment, 0)

	for rows.Next() {
		payment := entities.Payment{}

		err := rows.Scan(&payment.ID, &payment.OrderNumber, &payment.PaymentMethod, &payment.PaymentType, &payment.Amount)
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

func (r *St) FindOnePayment(ctx context.Context, orderNumber string) (entities.Payment, error) {
	q := `
		SELECT ID, ORDER_NUMBER, PAYMENT_METHOD, PAYMENT_TYPE, AMOUNT FROM public.payment WHERE order_id = &1`

	//Trace

	var payment entities.Payment
	err := r.client.QueryRow(ctx, q, orderNumber).Scan(&payment.ID, &payment.OrderNumber, &payment.PaymentMethod, &payment.PaymentType, &payment.Amount)
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}

//func (r *St) UpdatePayment(ctx context.Context, organization entities.Organization) error {
//	//TODO implement me
//	panic("implement me")
//}
//

func (r *St) DeletePayment(ctx context.Context, orderNumber string) error {
	q := `
		DELETE FROM payment
		WHERE order_number = $1;`

	var id string

	if err := r.client.QueryRow(ctx, q, orderNumber).Scan(&id); err != nil {
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
