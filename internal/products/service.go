package products

type Service interface {
	GetAll() ([]Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	return s.repository.GetAll()
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}