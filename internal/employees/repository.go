package employees

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
	Store(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Employee, error) {
	return es, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id, cardNymberId int, firstName, lastName string, warehouseId int) (Employee, error) {
	e := Employee{id, cardNymberId, firstName, lastName, warehouseId}
	es = append(es, e)
	lastID = e.ID
	return e, nil
}
