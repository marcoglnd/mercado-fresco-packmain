package domain

import (
	"context"
)

type PurchaseOrder struct {
	ID            int64  `json:"id" binding:"required"`
	OrderNumber   string `json:"order_number" binding:"required"`
	OrderDate     string `json:"order_date" binding:"required"`
	TrackingCode  string `json:"tracking_code" binding:"required"`
	BuyerId       int64  `json:"buyer_id" binding:"required"`
	CarrierId     int64  `json:"carrier_id" binding:"required"`
	OrderStatusId int64  `json:"order_status_id" binding:"required"`
	WarehouseId   int64  `json:"warehouse_id" binding:"required"`
}

type PurchaseOrderRepository interface {
	Create(
		ctx context.Context, orderNumber, orderDate, trackingCode string, buyerId, carrierId, orderStatusId, warehouseId int64) (*PurchaseOrder, error)
	GetByOrderNumber(ctx context.Context, orderNumber string) (*PurchaseOrder, error)
}

type PurchaseOrderService interface {
	Create(
		ctx context.Context, orderNumber, orderDate, trackingCode string, buyerId, carrierId, orderStatusId, warehouseId int64) (*PurchaseOrder, error)
}
