package warehouses

type Service interface {
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

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Create(
	warehouseCode string,
	address string,
	telephone string,
	minimumCapacity int,
	minimumTemperature int,
) (*Warehouse, error) {

	warehouse, err := s.repository.Create(
		warehouseCode,
		address,
		telephone,
		minimumCapacity,
		minimumTemperature,
	)

	if err != nil {
		return &Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) Update(data interface{}) (*Warehouse, error) {
	return &Warehouse{}, nil
}
func (s *service) FindById(id int) (*Warehouse, error) {
	return &Warehouse{}, nil
}
func (s *service) FindByWarehouseCode(warehouseCode string) (*Warehouse, error) {
	foundWarehouse, err := s.repository.FindByWarehouseCode(warehouseCode)

	if err != nil {
		return &Warehouse{}, err
	}

	return foundWarehouse, nil
}
func (s *service) GetAll() ([]Warehouse, error) {
	return []Warehouse{}, nil
}
func (s *service) Delete(id int) error {
	return nil
}
