package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/domain"

func CreateRandomPurchaseOrder() domain.PurchaseOrder {
	purchaseOrder := domain.PurchaseOrder{
		ID:            1,
		OrderNumber:   RandomString(6),
		OrderDate:     RandomString(6),
		TrackingCode:  RandomString(6),
		BuyerId:       RandomInt(0, 10),
		CarrierId:     RandomInt(0, 10),
		OrderStatusId: 1,
		WarehouseId:   RandomInt(0, 10),
	}
	return purchaseOrder
}
