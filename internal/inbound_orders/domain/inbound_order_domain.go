package domain

import "context"

type InboundOrderRepository interface {
	GetAll(ctx context.Context) (*[]InboundOrder, error)
	Create(ctx context.Context, inboundOrder *InboundOrder) (*InboundOrder, error)
}

type InboundOrderService interface {
	GetAll(ctx context.Context) (*[]InboundOrder, error)
	Create(ctx context.Context, inboundOrder *InboundOrder) (*InboundOrder, error)
}
