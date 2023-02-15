package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateOrder(ctx context.Context, order *entities.CreateOrderDTO) error {
	q := `
		INSERT INTO order (order_number, organization_id, paid, sent_to_1c)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	var id string

	if err := r.client.QueryRow(ctx, q, order.OrderNumber, order.OrganizationId, order.Paid, order.SentTo1C).Scan(&id); err != nil {
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

func (r *St) FindAllOrders(ctx context.Context) (u []entities.Order, err error) {
	q := `
		SELECT ID, ORDER_NUMBER, ORGANIZATION_BIN, PAID, SENT_TO_1C FROM public.order`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	orders := make([]entities.Order, 0)

	for rows.Next() {
		var order entities.Order

		err := rows.Scan(&order.ID, &order.OrderNumber, &order.OrganizationBin, &order.Paid, &order.SentTo1C)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *St) FindOneOrder(ctx context.Context, orderNumber string) (entities.Order, error) {
	q := `
		SELECT ID, ORDER_NUMBER, ORGANIZATION_BIN, PAID, SENT_TO_1C FROM public.order WHERE ORDER_NUMBER = &1`

	//Trace

	var order entities.Order
	err := r.client.QueryRow(ctx, q, orderNumber).Scan(&order.ID, &order.OrderNumber, &order.OrganizationBin, &order.Paid, &order.SentTo1C)
	if err != nil {
		return entities.Order{}, err
	}

	return order, nil

}

//	func (r *St) UpdateOrder(ctx context.Context, organization entities.Organization) error {
//		//TODO implement me
//		panic("implement me")
//	}
func (r *St) DeleteOrder(ctx context.Context, orderNumber string) error {
	q := `
		DELETE FROM orders
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
