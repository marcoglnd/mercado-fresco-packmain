package products

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	LastId() (int, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(
		id int, description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
	Delete(id int) error
}

var listOfProducts []Product = []Product{}

type repository struct{ db *sql.DB }

func (repository) GetAll() ([]Product, error) {
	return listOfProducts, nil
}

func (repository) GetById(id int) (Product, error) {
	var product Product
	foundProduct := false
	for i := range listOfProducts {
		if listOfProducts[i].Id == int64(id) {
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
	return int(lastId), nil
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

func (r *repository) CreateNewProduct(ctx context.Context, product *Product) (*Product, error) {
	newProduct := Product{}
	result, err := r.db.ExecContext(
		ctx,
		sqlStore,
		&product.Description,
		&product.ExpirationRate,
		&product.FreezingRate,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.ProductCode,
		&product.RecommendedFreezingTemperature,
		&product.Width,
		&product.ProductTypeId,
		&product.SellerId,
	)
	if err != nil {
		return &newProduct, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return &newProduct, err
	}
	newProduct.Id = insertedId
	return &newProduct, nil
}

func (repository) Update(
	id int, description string, expirationRate, freezingRate int,
	height, length, netWeight float64, productCode string,
	recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error) {
	prod := Product{
		Id:                             int64(id),
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
		if listOfProducts[i].Id == int64(id) {
			prod.Id = int64(id)
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
		if listOfProducts[i].Id == int64(id) {
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

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
