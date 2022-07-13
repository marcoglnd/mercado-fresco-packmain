package utils

import (
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
)

func CreateRandomSeller() domain.Seller {
	seller := domain.Seller{
		ID:           1,
		Cid:          RandomInt64(),
		Company_name: RandomCategory(),
		Address:      RandomCategory(),
		Telephone:    RandomCategory(),
		LocalityID:   RandomInt64(),
	}
	return seller
}

func CreateRandomListSeller() []domain.Seller {
	var listOfSellers []domain.Seller
	for i := 1; i <= 5; i++ {
		seller := CreateRandomSeller()
		seller.ID = int64(i)
		listOfSellers = append(listOfSellers, seller)
	}
	return listOfSellers
}
