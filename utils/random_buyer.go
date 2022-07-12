package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"

func CreateRandomBuyer() domain.Buyer {
	buyer := domain.Buyer{
		ID:           1,
		CardNumberID: RandomString(3),
		FirstName:    RandomString(6),
		LastName:     RandomString(6),
	}
	return buyer
}

func CreateRandomListBuyers() []domain.Buyer {
	var listOfBuyers []domain.Buyer
	for i := 1; i <= 5; i++ {
		buyer := CreateRandomBuyer()
		buyer.ID = int64(i)
		listOfBuyers = append(listOfBuyers, buyer)
	}
	return listOfBuyers
}

func CreateRandomReportPurchaseOrder() domain.PurchaseOrdersResponse {
	report := domain.PurchaseOrdersResponse{
		ID:                  1,
		CardNumberID:        RandomString(3),
		FirstName:           RandomString(6),
		LastName:            RandomString(6),
		PurchaseOrdersCount: RandomInt(0, 10),
	}
	return report
}

func CreateRandomListReportPurchaseOrder() []domain.PurchaseOrdersResponse {
	var listOfReports []domain.PurchaseOrdersResponse
	for i := 1; i <= 5; i++ {
		report := CreateRandomReportPurchaseOrder()
		report.ID = int64(i)
		listOfReports = append(listOfReports, report)
	}
	return listOfReports
}
