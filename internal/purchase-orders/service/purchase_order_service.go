package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/domain"
)

type purchaseOrderService struct {
	repository domain.PurchaseOrderRepository
}

func NewPurchaseOrderService(sr domain.PurchaseOrderRepository) domain.PurchaseOrderService {
	return &purchaseOrderService{repository: sr}
}

func (s purchaseOrderService) Create(ctx context.Context,
	orderNumber,
	orderDate,
	trackingCode string,
	buyerId,
	carrierId,
	orderStatusId,
	warehouseId int64,
) (*domain.PurchaseOrder, error) {
	foundPurchaseOrder, err := s.repository.GetByOrderNumber(ctx, orderNumber)
	if err != nil {
		return nil, err
	}

	if foundPurchaseOrder != nil {
		return nil, domain.ErrDuplicatedID
	}

	purchaseOrder, err := s.repository.Create(
		ctx,
		orderNumber,
		orderDate,
		trackingCode,
		buyerId,
		carrierId,
		orderStatusId,
		warehouseId,
	)
	if err != nil {
		return purchaseOrder, err
	}

	return purchaseOrder, nil
}
