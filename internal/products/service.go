package products

import "fmt"

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	CreateNewProduct(description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
	Update(
		id int, description string, expirationRate, freezingRate int,
		height, length, netWeight float64, productCode string,
		recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	listOfProducts, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listOfProducts, nil
}

func (s service) GetById(id int) (Product, error) {
	pr, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return pr, nil
}


func (s *service) VerifyProductCode(productCode string) (bool, error){
	list, err := s.repository.GetAll()
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


func (s *service) CreateNewProduct(
	description string, expirationRate, freezingRate int,
	height, length, netWeight float64, productCode string,
	recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	if verification, err := s.VerifyProductCode(productCode); !verification {
		return Product{}, err
	}
	newProd, err := s.repository.CreateNewProduct(
		id, description, expirationRate, freezingRate, height,
		length, netWeight, productCode, recommendedFreezingTemperature,
		width, productTypeId, sellerId)
	if err != nil {
		return Product{}, err
	}
	return newProd, nil
}

func (s service) Update(
	id int, description string, expirationRate, freezingRate int,
	height, length, netWeight float64, productCode string,
	recommendedFreezingTemperature, width float64, productTypeId, sellerId int) (Product, error) {
	product, err := s.repository.Update(
		id, description, expirationRate, freezingRate, height,
		length, netWeight, productCode, recommendedFreezingTemperature,
		width, productTypeId, sellerId)
	if err != nil {
		return Product{}, err
	}
	return product, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
