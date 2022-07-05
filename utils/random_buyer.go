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
