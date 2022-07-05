package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"

func CreateRandomProduct() domain.Product {
	product := domain.Product{
		Id:                             1,
		Description:                    RandomCategory(),
		ExpirationRate:                 RandomInt64(),
		FreezingRate:                   RandomInt64(),
		Height:                         RandomFloat64(),
		Length:                         RandomFloat64(),
		NetWeight:                      RandomFloat64(),
		ProductCode:                    RandomCategory(),
		RecommendedFreezingTemperature: RandomFloat64(),
		Width:                          RandomFloat64(),
		ProductTypeId:                  RandomInt64(),
		SellerId:                       RandomInt64(),
	}
	return product
}

func CreateRandomListProduct() []domain.Product {
	var listOfProducts []domain.Product
	for i := 1; i <= 5; i++ {
		product := CreateRandomProduct()
		product.Id = int64(i)
		listOfProducts = append(listOfProducts, product)
	}
	return listOfProducts
}

func CreateRandomProductRecords() domain.ProductRecords {
	records := domain.ProductRecords{
		LastUpdateDate: RandomString(10),
		PurchasePrice:  RandomFloat64(),
		SalePrice:      RandomFloat64(),
		ProductId:      RandomInt64(),
	}
	return records
}

func CreateRandomQtyOfRecords() domain.QtyOfRecords {
	qtyOfRecords := domain.QtyOfRecords{
		ProductId:    RandomInt64(),
		Description:  RandomString(10),
		RecordsCount: RandomInt64(),
	}
	return qtyOfRecords
}
