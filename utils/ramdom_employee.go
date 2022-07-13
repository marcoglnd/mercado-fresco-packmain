package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"

func CreateRandomEmployee() domain.Employee {
	employee := domain.Employee{
		ID:           1,
		CardNumberId: RandomString(3),
		FirstName:    RandomString(6),
		LastName:     RandomString(6),
		WarehouseId:  RandomInt64(),
	}
	return employee
}

func CreateRandomListEmployees() []domain.Employee {
	var listOfEmployees []domain.Employee
	for i := 1; i <= 5; i++ {
		employee := CreateRandomEmployee()
		employee.ID = int64(i)
		listOfEmployees = append(listOfEmployees, employee)
	}
	return listOfEmployees
}

func CreateRandomReportInboundOrder() domain.InboundOrderResponse {
	inboundOrder := domain.InboundOrderResponse{
		ID:                 1,
		CardNumberId:       RandomString(3),
		FirstName:          RandomString(6),
		LastName:           RandomString(6),
		WarehouseId:        RandomInt64(),
		InboundOrdersCount: RandomInt64(),
	}
	return inboundOrder
}

func CreateRamdomListReportInboundOrders() []domain.InboundOrderResponse {
	var listOfInboundOrders []domain.InboundOrderResponse
	for i := 1; i <= 5; i++ {
		inboundOrder := CreateRandomReportInboundOrder()
		inboundOrder.ID = int64(i)
		listOfInboundOrders = append(listOfInboundOrders, inboundOrder)
	}
	return listOfInboundOrders
}
