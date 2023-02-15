package core

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateTradePoint(ctx *gin.Context, obj *entities.CreateTradePointDTO) error {
	err := s.repo.CreateTradePoint(ctx, obj)

	return err
}

func (s *St) DeleteTradePoint(ctx *gin.Context, bin string, tradePointId string) error {
	err := s.repo.DeleteTradePoint(ctx, bin, tradePointId)

	return err
}

//
//func (s *St) UpdateTradePoint(ctx *gin.Context, obj *entities.TradePoint) error {
//	err := s.repo.UpdateTradePoint(ctx, obj)
//
//	return err
//}

func (s *St) FindAllTradePoints(ctx *gin.Context) ([]entities.TradePoint, error) {
	tradePoints, err := s.repo.FindAllTradePoints(ctx)

	return tradePoints, err
}

func (s *St) FindOneTradePoint(ctx *gin.Context, obj *entities.TradePoint) (entities.TradePoint, error) {
	tradePoint, err := s.repo.FindOneTradePoint(ctx, obj.ID)

	return tradePoint, err
}
