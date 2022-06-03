package employees

type Service interface {
	GetAll() ([]Employee, error)
	Store(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Employee, error) {
	es, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return es, nil
}

func (s *service) Store(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Employee{}, err
	}

	lastID++

	employee, err := s.repository.Store(id, cardNymberId, firstName, lastName, warehouseId)

	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}
