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

func (r *St) CreateOrder(ctx context.Context, order *entities.CreateOrderDTO) error {
	q := `
		INSERT INTO orders (created, modified, order_number, organization_bin, status)
		VALUES ($1, $2, $3, $4, $5)`

	if err := r.db.Exec(ctx, q, order.Created, order.Modified, order.OrderNumber, order.OrganizationBin, order.Status); err != nil {
		return r.ErorrHandler(err)
	}
	return nil
}

func (r *St) FindAllOrders(ctx context.Context) (u []entities.Order, err error) {
	q := `
		SELECT CREATED, MODIFIED, ORDER_NUMBER, ORGANIZATION_BIN, STATUS FROM orders`
	u, err = r.ListOrders(ctx, q)
	return u, err
}

func (r *St) FindAllUnpaidOrders(ctx context.Context) (u []entities.UnPaidOrder, err error) {
	q := `
			SELECT o.CREATED, o.ORDER_NUMBER, o.ORGANIZATION_BIN , p.PAYMENT_ID
			FROM orders o
         		LEFT JOIN payment p
                	   ON o.ORDER_NUMBER = p.ORDER_NUMBER
			WHERE p.status = 'Created' OR p.status = 'Wait'
			`
	rows, err := r.db.Query(ctx, q)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	orders := make([]entities.UnPaidOrder, 0)

	for rows.Next() {
		var order entities.UnPaidOrder

		err := rows.Scan(&order.Created, &order.OrderNumber, &order.OrganizationBin, &order.PaymentId)
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

func (r *St) FindAllPaidOrders(ctx context.Context) (u []entities.PaidOrder, err error) {
	q := `
			SELECT o.ORDER_NUMBER, o.ORGANIZATION_BIN , p.PAYMENT_ID
			FROM orders o
         		LEFT JOIN payment p
                	   ON o.ORDER_NUMBER = p.ORDER_NUMBER
			WHERE p.status = 'Success'
			`
	rows, err := r.db.Query(ctx, q)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	orders := make([]entities.PaidOrder, 0)

	for rows.Next() {
		var order entities.PaidOrder

		err := rows.Scan(&order.OrderNumber, &order.OrganizationBin, &order.PaymentId)
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
		SELECT CREATED, MODIFIED, ORDER_NUMBER, ORGANIZATION_BIN, STATUS FROM orders WHERE ORDER_NUMBER = $1`

	//Trace

	var order entities.Order
	err := r.db.QueryRow(ctx, q, orderNumber).Scan(&order.Created, &order.Modified, &order.OrderNumber, &order.OrganizationBin, &order.Status)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Order{}, err
	}

	return order, nil

}

func (r *St) UpdateOrderStatus(ctx context.Context, orderNumber string, status string) error {
	q := `
		UPDATE orders SET status = $1, modified = $2 
		               WHERE order_number = $3;`

	if err := r.db.Exec(ctx, q, status, time.Now().Local(), orderNumber); err != nil {
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

func (r *St) DeleteOrder(ctx context.Context, orderNumber string) error {
	q := `
		DELETE FROM orders
		WHERE order_number = $1;`

	if err := r.db.Exec(ctx, q, orderNumber); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return r.ErorrHandler(err)
	}

	return nil
}

func (r *St) ListOrders(ctx context.Context, query string) (u []entities.Order, err error) {
	rows, err := r.db.Query(ctx, query)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	orders := make([]entities.Order, 0)

	for rows.Next() {
		var order entities.Order

		err := rows.Scan(&order.Created, &order.Modified, &order.OrderNumber, &order.OrganizationBin, &order.Status)
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
