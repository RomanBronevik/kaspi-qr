package core

import (
	"context"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
)

type TradePoint struct {
	r *St
}

func NewTradePoint(r *St) *TradePoint {
	return &TradePoint{r: r}
}

func (c *TradePoint) List(ctx context.Context, pars *entities.TradePointListParsSt) ([]*entities.TradePointSt, error) {
	if pars.OrgBin == nil {
		return nil, errs.OrgBinRequired
	}

	items, err := c.r.prv.TradePointList(*pars.OrgBin)
	if err != nil {
		return nil, err
	}

	result := make([]*entities.TradePointSt, len(items))
	for i, item := range items {
		result[i] = &entities.TradePointSt{
			Id:   item.TradePointId,
			Name: item.TradePointName,
		}
	}

	return result, nil
}
