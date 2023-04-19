package old

import (
	"context"
	"errors"
	"fmt"
	"kaspi-qr/internal/cns"
	"kaspi-qr/internal/domain/entities"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (d *St) CreateOrder(ctx context.Context, order *entities.CreateOrderDTO) error {
	q := `
		INSERT INTO orders (created, modified, order_number, organization_bin, status)
		VALUES ($1, $2, $3, $4, $5)`

	return d.db.Exec(ctx, q, order.Created, order.Modified, order.OrderNumber, order.OrganizationBin, order.Status)
}

func (d *St) FindAllOrders(ctx context.Context) (u []*entities.Order, err error) {
	q := `
		SELECT CREATED, MODIFIED, ORDER_NUMBER, ORGANIZATION_BIN, STATUS
		FROM orders`

	rows, err := d.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*entities.Order, 0)

	for rows.Next() {
		order := &entities.Order{}

		err = rows.Scan(&order.Created, &order.Modified, &order.OrderNumber, &order.OrganizationBin, &order.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (d *St) FindAllUnpaidOrders(ctx context.Context) ([]*entities.UnPaidOrder, error) {
	q := `
		SELECT o.CREATED, o.ORDER_NUMBER, o.ORGANIZATION_BIN , p.PAYMENT_ID
		FROM orders o
			LEFT JOIN payment p
				   ON o.ORDER_NUMBER = p.ORDER_NUMBER
		WHERE p.status = $1 OR p.status = $2`

	rows, err := d.db.Query(ctx, q, cns.StatusCreated, cns.StatusWait)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*entities.UnPaidOrder, 0)

	for rows.Next() {
		order := &entities.UnPaidOrder{}

		err = rows.Scan(&order.Created, &order.OrderNumber, &order.OrganizationBin, &order.PaymentId)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (d *St) FindAllPaidOrders(ctx context.Context) ([]*entities.PaidOrder, error) {
	q := `
		SELECT o.CREATED, o.ORDER_NUMBER, o.ORGANIZATION_BIN , p.PAYMENT_ID
		FROM orders o
			LEFT JOIN payment p
				   ON o.ORDER_NUMBER = p.ORDER_NUMBER
		WHERE p.status = $1`

	rows, err := d.db.Query(ctx, q, cns.StatusSuccess)
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

func (d *St) FindOneOrder(ctx context.Context, orderNumber string) (entities.Order, error) {
	q := `
		SELECT CREATED, MODIFIED, ORDER_NUMBER, ORGANIZATION_BIN, STATUS FROM orders WHERE ORDER_NUMBER = $1`

	//Trace

	var order entities.Order
	err := d.db.QueryRow(ctx, q, orderNumber).Scan(&order.Created, &order.Modified, &order.OrderNumber, &order.OrganizationBin, &order.Status)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Order{}, err
	}

	return order, nil

}

func (d *St) UpdateOrderStatus(ctx context.Context, orderNumber string, status string) error {
	q := `
		UPDATE orders SET status = $1, modified = $2 
		               WHERE order_number = $3;`

	if err := d.db.Exec(ctx, q, status, time.Now().Local(), orderNumber); err != nil {
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

func (d *St) DeleteOrder(ctx context.Context, orderNumber string) error {
	q := `
		DELETE FROM orders
		WHERE order_number = $1;`

	if err := d.db.Exec(ctx, q, orderNumber); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return d.ErorrHandler(err)
	}

	return nil
}
