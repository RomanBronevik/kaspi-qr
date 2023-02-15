package pg

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/jackc/pgconn"
//	"kaspi-qr/internal/domain/entities"
//)
//
//func (r *St) CreateRequestLog(ctx context.Context, requestLog *entities.CreateRequestLogDTO) error {
//	q := `
//		INSERT INTO trade_point (name, trade_point_id, organization_id)
//		VALUES ($1, $2, &3)
//		RETURNING id`
//
//	var id string
//
//	if err := r.client.QueryRow(ctx, q, tradePoint.Name, tradePoint.TradePointId, tradePoint.OrganizationId).Scan(&id); err != nil {
//		var pgErr *pgconn.PgError
//		if errors.Is(err, pgErr) {
//			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
//			fmt.Println(newErr)
//			return newErr
//		}
//		return err
//	}
//
//	return nil
//}
//
//func (r *St) FindAllRequestLogs(ctx context.Context) (u []entities.TradePoint, err error) {
//	q := `
//		SELECT ID, NAME, TRADE_POINT_ID, ORGANIZATION_ID FROM public.trade_point`
//	rows, err := r.client.Query(ctx, q)
//	if err != nil {
//		return nil, err
//	}
//
//	tradePoints := make([]entities.TradePoint, 0)
//
//	for rows.Next() {
//		var tradePoint entities.TradePoint
//
//		err := rows.Scan(&tradePoint.ID, &tradePoint.Name, &tradePoint.TradePointId, &tradePoint.OrganizationId)
//		if err != nil {
//			return nil, err
//		}
//
//		tradePoints = append(tradePoints, tradePoint)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return tradePoints, nil
//}
//
//func (r *St) FindOneRequestLog(ctx context.Context, id string) (entities.TradePoint, error) {
//	q := `
//		SELECT ID, NAME, TRADE_POINT_ID, ORGANIZATION_ID FROM public.trade_point WHERE id = &1`
//
//	//Trace
//
//	var tradePoint entities.TradePoint
//	err := r.client.QueryRow(ctx, q, id).Scan(&tradePoint.ID, &tradePoint.Name, &tradePoint.TradePointId, &tradePoint.OrganizationId)
//	if err != nil {
//		return entities.TradePoint{}, err
//	}
//
//	return tradePoint, nil
//
//}
//
//func (r *St) DeleteRequestLog(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
