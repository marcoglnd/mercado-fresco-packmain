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