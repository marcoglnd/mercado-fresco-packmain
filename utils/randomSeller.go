package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"

func CreateRandomSeller() domain.Seller {
	seller := domain.Seller{
		ID:           1,
		Cid:          RandomCode(),
		Company_name: RandomCategory(),
		Address:      RandomCategory(),
		Telephone:    RandomCategory(),
	}
	return seller
}

func CreateRandomListSeller() []domain.Seller {
	var listOfProducts []domain.Seller
	for i := 1; i <= 5; i++ {
		seller := CreateRandomSeller()
		seller.ID = int64(i)
		listOfProducts = append(listOfProducts, seller)
	}
	return listOfProducts
}
