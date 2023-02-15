package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateTradePoint(ctx context.Context, tradePoint *entities.CreateTradePointDTO) error {
	q := `
		INSERT INTO trade_point (name, trade_point_id, organization_bin)
		VALUES ($1, $2, &3)
		RETURNING id`

	var id string

	if err := r.client.QueryRow(ctx, q, tradePoint.Name, tradePoint.TradePointId, tradePoint.OrganizationBin).Scan(&id); err != nil {
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

func (r *St) FindAllTradePoints(ctx context.Context) (u []entities.TradePoint, err error) {
	q := `
		SELECT ID, NAME, TRADE_POINT_ID, ORGANIZATION_BIN FROM public.trade_point`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	tradePoints := make([]entities.TradePoint, 0)

	for rows.Next() {
		var tradePoint entities.TradePoint

		err := rows.Scan(&tradePoint.ID, &tradePoint.Name, &tradePoint.TradePointId, &tradePoint.OrganizationBin)
		if err != nil {
			return nil, err
		}

		tradePoints = append(tradePoints, tradePoint)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tradePoints, nil
}

func (r *St) FindOneTradePoint(ctx context.Context, id string) (entities.TradePoint, error) {
	q := `
		SELECT ID, NAME, TRADE_POINT_ID, ORGANIZATION_BIN FROM public.trade_point WHERE id = &1`

	//Trace

	var tradePoint entities.TradePoint
	err := r.client.QueryRow(ctx, q, id).Scan(&tradePoint.ID, &tradePoint.Name, &tradePoint.TradePointId, &tradePoint.OrganizationBin)
	if err != nil {
		return entities.TradePoint{}, err
	}

	return tradePoint, nil

}

//	func (r *St) UpdateTradePoint(ctx context.Context, organization entities.Organization) error {
//		//TODO implement me
//		panic("implement me")
//	}
func (r *St) DeleteTradePoint(ctx context.Context, bin string, tradePointId string) error {
	q := `
		DELETE FROM trade_point
		WHERE organization_bin = $1 AND trade_point_id = $2;`

	var id string

	if err := r.client.QueryRow(ctx, q, bin, tradePointId).Scan(&id); err != nil {
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
