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

var es []Employee = []Employee{}
var lastID int

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) GetAll() ([]Employee, error) {
	return es, nil
}

func (repository) GetById(id int) (Employee, error) {
	var e Employee
	foundEmployee := false
	for i := range es {
		if es[i].ID == id {
			e = es[i]
			foundEmployee = true
		}
	}
	if !foundEmployee {
		return Employee{}, fmt.Errorf("Employee %d not found", id)
	}
	return e, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Create(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	for i := range es {
		if es[i].CardNumberId == cardNumberId {
			return Employee{}, fmt.Errorf("CardNumberId %s already exist", cardNumberId)
		}
	}
	e := Employee{id, cardNumberId, firstName, lastName, warehouseId}
	es = append(es, e)
	lastID = e.ID
	return e, nil
}

func (repository) Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (Employee, error) {
	e := Employee{id, cardNumberId, firstName, lastName, warehouseId}
	updated := false
	for i := range es {
		if es[i].ID == id {
			e.ID = id
			es[i] = e
			updated = true
		}
	}
	if !updated {
		return Employee{}, fmt.Errorf("Employee %d not found", id)
	}
	return e, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range es {
		if es[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Employee %d not found", id)
	}
	es = append(es[:index], es[index+1:]...)
	return nil
}
