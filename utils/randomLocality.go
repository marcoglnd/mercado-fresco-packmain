package utils

import (
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
)

func CreateRandomLocality() domain.Locality {
	records := domain.Locality{
		LocalityName: RandomString(10),
		ProvinceID:   RandomInt64(),
	}
	return records
}

func CreateRandomGetLocality() domain.GetLocality {
	records := domain.GetLocality{
		ID:           RandomInt64(),
		LocalityName: RandomString(10),
		ProvinceName: RandomString(10),
		CountryName:  RandomString(10),
	}
	return records
}

func CreateRandomQtyOfSellers() domain.QtyOfSellers {
	qtyOfSellers := domain.QtyOfSellers{
		LocalityID:   RandomInt64(),
		LocalityName: RandomString(10),
		SellersCount: RandomInt64(),
	}
	return qtyOfSellers
}

func CreateRandomListQtyOfSellers() []domain.QtyOfSellers {
	var listOfSellers []domain.QtyOfSellers
	for i := 1; i <= 5; i++ {
		qtySellers := CreateRandomQtyOfSellers()
		listOfSellers = append(listOfSellers, qtySellers)
	}
	return listOfSellers
}
