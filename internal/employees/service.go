package employees

type Service interface {
	GetAll() ([]Employee, error)
	GetEmployee(id int) (Employee, error)
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

func (s service) GetAll() ([]Employee, error) {
	es, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return es, nil
}

func (s service) GetEmployee(id int) (Employee, error) {
	es, err := s.repository.GetEmployee(id)
	if err != nil {
		return Employee{}, err
	}
	return es, nil
}

func (s service) Create(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Employee{}, err
	}

	lastID++

	employee, err := s.repository.Create(lastID, cardNymberId, firstName, lastName, warehouseId)

	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s service) Update(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	employee, err := s.repository.Update(id, cardNymberId, firstName, lastName, warehouseId)
	if err != nil {
		return Employee{}, err
	}
	return employee, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
