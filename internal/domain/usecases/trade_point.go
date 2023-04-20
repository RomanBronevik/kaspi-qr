package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

func (u *St) TradePointList(ctx context.Context,
	pars *entities.TradePointListParsSt) ([]*entities.TradePointSt, error) {
	return u.cr.TradePoint.List(ctx, pars)
}
