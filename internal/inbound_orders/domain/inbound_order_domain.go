package domain

import "context"

type InboundOrderRepository interface {
	Create(ctx context.Context, inboundOrder *InboundOrder) (*InboundOrder, error)
}

type InboundOrderService interface {
	Create(ctx context.Context, inboundOrder *InboundOrder) (*InboundOrder, error)
}
