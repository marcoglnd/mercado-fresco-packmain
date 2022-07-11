package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain"

func CreateRandomInboundOrder() domain.InboundOrder {
	InboundOrder := domain.InboundOrder{
		ID:             1,
		OrderDate:      RandomString(10),
		OrderNumber:    RandomString(3),
		EmployeeId:     RandomInt64(),
		ProductBatchId: RandomInt64(),
		WarehouseId:    RandomInt64(),
	}
	return InboundOrder
}

func CreateRandomListInboundOrders() []domain.InboundOrder {
	var listOfInboundOrders []domain.InboundOrder
	for i := 1; i <= 5; i++ {
		InboundOrder := CreateRandomInboundOrder()
		InboundOrder.ID = int64(i)
		listOfInboundOrders = append(listOfInboundOrders, InboundOrder)
	}
	return listOfInboundOrders
}
