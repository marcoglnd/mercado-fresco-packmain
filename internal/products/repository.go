package products

type Repository interface {
	GetAll() ([]Product, error)
}

type repository struct{
	db []Product
}

func (r *repository) GetAll() ([]Product, error) {
	return r.db, nil
}

func NewRepository(db []Product) Repository {
	return &repository{
		db: db,
	}
}
