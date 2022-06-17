package employees

import "github.com/stretchr/testify/mock"

type EmployeeService struct {
	mock.Mock
}

func (e *EmployeeService) GetAll() ([]Employee, error) {
	return []Employee{}, nil
}

func (e *EmployeeService) GetById(id int) (Employee, error) {
	return Employee{}, nil
}

func (e *EmployeeService) Create(cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	return Employee{}, nil
}

func (e *EmployeeService) Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	return Employee{}, nil
}

func (e *EmployeeService) Delete(id int) error {
	return nil
}
