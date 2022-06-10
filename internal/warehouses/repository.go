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
	Update(warehouse *Warehouse) error
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
func (r *repository) Update(warehouse *Warehouse) error {
	for i := range warehouses {
		if warehouses[i].ID == warehouse.ID {
			warehouses[i].WarehouseCode = warehouse.WarehouseCode
			warehouses[i].Address = warehouse.Address
			warehouses[i].Telephone = warehouse.Telephone
			warehouses[i].MinimumCapacity = warehouse.MinimumCapacity
			warehouses[i].MinimumTemperature = warehouse.MinimumTemperature
			break
		}
	}
	return nil
}

func (r *repository) FindById(id int) (*Warehouse, error) {
	var foundWarehouse *Warehouse
	for _, w := range warehouses {
		if w.ID == id {
			foundWarehouse = &w
			break
		}
	}
	return foundWarehouse, nil
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
	return warehouses, nil
}
func (r *repository) Delete(id int) error {
	for i, w := range warehouses {
		if w.ID == id {
			warehouses = append(warehouses[:i], warehouses[i+1:]...)
			break
		}
	}
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
