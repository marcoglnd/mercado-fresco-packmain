package domain

import "context"

type InboundOrdersRepository interface {
	Create(ctx context.Context, inboundOrder *InboundOrder) (*InboundOrder, error)
}

type InboundOrdersService interface {
	Create(ctx context.Context, inboundOrder *InboundOrder) (*InboundOrder, error)
}
