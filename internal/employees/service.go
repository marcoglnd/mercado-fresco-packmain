package employees

type Service interface {
	GetAll() ([]Employee, error)
	Create(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error)
	Update(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error)
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

func (s *service) GetAll() ([]Employee, error) {
	es, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return es, nil
}

func (s *service) Create(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Employee{}, err
	}

	lastID++

	employee, err := s.repository.Create(id, cardNymberId, firstName, lastName, warehouseId)

	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s *service) Update(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	return s.repository.Update(id, cardNymberId, firstName, lastName, warehouseId)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
