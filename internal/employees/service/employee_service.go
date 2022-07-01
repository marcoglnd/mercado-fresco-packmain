package service

type Service interface {
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	Create(cardNumberId, firstName, lastName string, warehouseId int) (Employee, error)
	Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error)
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
	listEmployees, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return listEmployees, nil
}

func (s service) GetById(id int) (Employee, error) {
	listEmployees, err := s.repository.GetById(id)
	if err != nil {
		return Employee{}, err
	}
	return listEmployees, nil
}

func (s service) Create(cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {

	employee, err := s.repository.Create(cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s service) Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	employee, err := s.repository.Update(id, cardNumberId, firstName, lastName, warehouseId)
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
