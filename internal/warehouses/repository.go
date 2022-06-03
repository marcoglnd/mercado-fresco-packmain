package warehouses

var warehouses []Warehouse = []Warehouse{}

type Repository interface {
	Create(
		warehouseCode string,
		address string,
		telephone string,
		minimumCapacity int,
		minimumTemperature int,
	) (*Warehouse, error)
	Update(data interface{}) (*Warehouse, error)
	FindById(id int) (*Warehouse, error)
	FindByWarehouseCode(warehouseCode string) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	Delete(id int) error
}

type repository struct{}

func (r *repository) Create(
	warehouseCode string,
	address string,
	telephone string,
	minimumCapacity int,
	minimumTemperature int,
) (*Warehouse, error) {
	newWarehouse := &Warehouse{
		ID:                 len(warehouses) + 1,
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}
	warehouses = append(warehouses, *newWarehouse)
	return newWarehouse, nil
}
func (r *repository) Update(data interface{}) (*Warehouse, error) {
	return &Warehouse{}, nil
}
func (r *repository) FindById(id int) (*Warehouse, error) {
	return &Warehouse{}, nil
}
func (r *repository) FindByWarehouseCode(warehouseCode string) (*Warehouse, error) {
	var foundWarehouse *Warehouse
	for _, w := range warehouses {
		if w.WarehouseCode == warehouseCode {
			foundWarehouse = &w
			break
		}
	}
	return foundWarehouse, nil
}
func (r *repository) GetAll() ([]Warehouse, error) {
	return []Warehouse{}, nil
}
func (r *repository) Delete(id int) error {
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
