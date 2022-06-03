package products

import "fmt"

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	LastId() (int, error)
	CreateNewProduct(
		id int, description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
	Update(
		id int, description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
}

var listOfProducts []Product = []Product{}

type repository struct{}

func (repository) GetAll() ([]Product, error) {
	return listOfProducts, nil
}

func (repository) GetById(id int) (Product, error) {
	var product Product
	foundProduct := false
	for i := range listOfProducts {
		if listOfProducts[i].Id == id {
			product = listOfProducts[i]
			foundProduct = true
		}
	}
	if !foundProduct {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return product, nil
}

func (repository) LastId() (int, error) {
	if len(listOfProducts) == 0 {
		return 0, nil
	}
	lastId := listOfProducts[len(listOfProducts)-1].Id + 1
	return lastId, nil
}

func (repository) CreateNewProduct(
	id int, description string, expirationRate, freezingRate int,
	height, length, netWeight float64, productCode string,
	recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error) {
	prod := Product{
		Id:                             id,
		Description:                    description,
		ExpirationRate:                 expirationRate,
		FreezingRate:                   freezingRate,
		Height:                         height,
		Length:                         length,
		NetWeight:                      netWeight,
		ProductCode:                    productCode,
		RecommendedFreezingTemperature: recommendedFreezingTemperature,
		Width:                          width,
		ProductTypeId:                  productTypeId,
		SellerId:                       sellerId,
	}
	listOfProducts = append(listOfProducts, prod)
	return prod, nil
}

func (repository) Update(
	id int, description string, expirationRate, freezingRate int,
	height, length, netWeight float64, productCode string,
	recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error) {
	prod := Product{
		Id:                             id,
		Description:                    description,
		ExpirationRate:                 expirationRate,
		FreezingRate:                   freezingRate,
		Height:                         height,
		Length:                         length,
		NetWeight:                      netWeight,
		ProductCode:                    productCode,
		RecommendedFreezingTemperature: recommendedFreezingTemperature,
		Width:                          width,
		ProductTypeId:                  productTypeId,
		SellerId:                       sellerId,
	}
	updated := false
	for i := range listOfProducts {
		if listOfProducts[i].Id == id {
			prod.Id = id
			listOfProducts[i] = prod
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return prod, nil
}

func NewRepository() Repository {
	return &repository{}
}
