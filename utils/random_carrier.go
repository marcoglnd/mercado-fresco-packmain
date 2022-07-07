package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"

func CreateRandomCarrier() domain.Carrier {
	carrier := domain.Carrier{
		ID:          0,
		Cid:         RandomString(4),
		CompanyName: RandomString(10),
		Address:     RandomString(10),
		Telephone:   RandomString(6),
		LocalityId:  1,
	}
	return carrier
}

func CreateRandomListCarriers() []domain.Carrier {
	var listOfCarriers []domain.Carrier
	for i := 1; i <= 5; i++ {
		carrier := CreateRandomCarrier()
		carrier.ID = int64(i)
		listOfCarriers = append(listOfCarriers, carrier)
	}
	return listOfCarriers
}
