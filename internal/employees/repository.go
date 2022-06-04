package employees

import "fmt"

type Employee struct {
	ID           int    `json:"id"`
	CardNumberId int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

var es []Employee
var lastID int

type Repository interface {
	GetAll() ([]Employee, error)
	GetEmployee(id int) (Employee, error)
	Create(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error)
	LastID() (int, error)
	Update(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Employee, error) {
	return es, nil
}

func (r *repository) GetEmployee(id int) (Employee, error) {
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

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Create(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	e := Employee{id, cardNymberId, firstName, lastName, warehouseId}
	es = append(es, e)
	lastID = e.ID
	return e, nil
}

func (r *repository) Update(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	e := Employee{id, cardNymberId, firstName, lastName, warehouseId}
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

func (r *repository) Delete(id int) error {
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
