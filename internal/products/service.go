package products

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) ([]Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	return s.repository.GetAll()
}

func (s service) GetById(id int) ([]Product, error) {
	pr, err := s.repository.GetById(id)
	if err != nil {
		return []Product{}, err
	}
	return []Product{pr}, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}