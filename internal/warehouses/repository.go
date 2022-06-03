package warehouses

var warehouses []Warehouse = []Warehouse{}

type Repository interface {
	Create(
		id int,
		warehouseCode string,
		address string,
		telephone string,
		minimumCapacity int,
		minimumTemperature int,
	) (*Warehouse, error)
	Update(data interface{}) (*Warehouse, error)
	FindById(id int) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	Delete(id int) error
}

type repository struct{}

func (r *repository) Create(
	id int,
	warehouseCode string,
	address string,
	telephone string,
	minimumCapacity int,
	minimumTemperature int,
) (*Warehouse, error)
func (r *repository) Update(data interface{}) (*Warehouse, error)
func (r *repository) FindById(id int) (*Warehouse, error)
func (r *repository) GetAll() ([]Warehouse, error)
func (r *repository) Delete(id int) error
