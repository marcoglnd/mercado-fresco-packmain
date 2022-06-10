package employees

import "fmt"

type Repository interface {
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	Create(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error)
	LastID() (int, error)
	Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error)
	Delete(id int) error
}

var listEmployees []Employee = []Employee{}
var lastID int

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) GetAll() ([]Employee, error) {
	return listEmployees, nil
}

func (repository) GetById(id int) (Employee, error) {
	var employee Employee
	foundEmployee := false
	for i := range listEmployees {
		if listEmployees[i].ID == id {
			employee = listEmployees[i]
			foundEmployee = true
		}
	}
	if !foundEmployee {
		return Employee{}, fmt.Errorf("Employee %d not found", id)
	}
	return employee, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Create(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	employee := Employee{id, cardNumberId, firstName, lastName, warehouseId}
	listEmployees = append(listEmployees, employee)
	lastID = employee.ID
	return employee, nil
}

func (repository) Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	employee := Employee{id, cardNumberId, firstName, lastName, warehouseId}
	updated := false
	for i := range listEmployees {
		if listEmployees[i].ID == id {
			employee.ID = id
			listEmployees[i] = employee
			updated = true
		}
	}
	if !updated {
		return Employee{}, fmt.Errorf("Employee %d not found", id)
	}
	return employee, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range listEmployees {
		if listEmployees[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Employee %d not found", id)
	}
	listEmployees = append(listEmployees[:index], listEmployees[index+1:]...)
	return nil
}
