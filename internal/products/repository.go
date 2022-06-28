package products

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/marcoglnd/mercado-fresco-packmain/db"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	LastId() (int, error)
	CreateNewProduct(
		description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
	Update(
		id int, description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
	Delete(id int) error
}

var listOfProducts []Product = []Product{}
var lastId int = 1
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
		return 1, nil
	}
	lastId := listOfProducts[len(listOfProducts)-1].Id + 1
	return lastId, nil
}

func (r *repository) VerifyProductCode(productCode string) (bool, error) {
	list, err := r.GetAll()
	if err != nil {
		return false, err
	}
	for _, prod := range list {
		if prod.ProductCode == productCode {
			return false, fmt.Errorf("product_code already used")
		}
	}
	return true, nil
}

func (r *repository) CreateNewProduct(
	description string,
	expirationRate,
	freezingRate int,
	height,
	length,
	netWeight float64,
	productCode string,
	recommendedFreezingTemperature,
	width float64,
	productTypeId,
	sellerId int,
) (Product, error) {
	// id, err := r.LastId()
	// if err != nil {
	// 	return Product{}, err
	// }
	// if verification, err := r.VerifyProductCode(productCode); !verification {
	// 	return Product{}, err
	// }
	prod := Product{
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
	// listOfProducts = append(listOfProducts, prod)
	// return prod, nil
	db := db.StorageDB
	stmt, err := db.Prepare(sqlStore)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result

	result, err = stmt.Exec(
		// lastId,
		description,
		expirationRate,
		freezingRate,
		height,
		length,
		netWeight,
		productCode,
		recommendedFreezingTemperature,
		width,
		productTypeId,
		sellerId,
	)
	if err != nil {
		return Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	prod.Id = int(insertedId)
	lastId = prod.Id
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
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return prod, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range listOfProducts {
		if listOfProducts[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product %d not found", id)
	}
	listOfProducts = append(listOfProducts[:index], listOfProducts[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
