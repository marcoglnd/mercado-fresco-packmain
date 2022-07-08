package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/inboundOrders/domain"
)

type inboundOrderService struct {
	repository domain.InboundOrdersRepository
}

func NewInboundOrderService(ir domain.InboundOrdersRepository) domain.InboundOrdersService {
	return &inboundOrderService{repository: ir}
}

func (i inboundOrderService) Create(ctx context.Context, inboundOrder *domain.InboundOrder) (*domain.InboundOrder, error) {
	inboundOrder, err := i.repository.Create(ctx, inboundOrder)

	if err != nil {
		return inboundOrder, err
	}

	return inboundOrder, nil
}
