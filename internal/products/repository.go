package products

import "fmt"

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
}

type repository struct {
	db []Product
}

func (r *repository) GetAll() ([]Product, error) {
	return r.db, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var product Product
	foundProduct := false
	for i := range r.db {
		if r.db[i].Id == id {
			product = r.db[i]
			foundProduct = true
		}
	}
	if !foundProduct {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return product, nil
}

func NewRepository(db []Product) Repository {
	return &repository{
		db: db,
	}
}
