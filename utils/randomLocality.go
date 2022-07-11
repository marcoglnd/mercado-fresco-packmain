package utils

import (
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
)

func CreateRandomLocality() domain.GetLocality {
	records := domain.GetLocality{
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
