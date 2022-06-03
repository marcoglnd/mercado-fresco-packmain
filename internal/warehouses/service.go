package warehouses

type Service interface {
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

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}
func (s *service) Create(
	id int,
	warehouseCode string,
	address string,
	telephone string,
	minimumCapacity int,
	minimumTemperature int,
) (*Warehouse, error)
func (s *service) Update(data interface{}) (*Warehouse, error)
func (s *service) FindById(id int) (*Warehouse, error)
func (s *service) GetAll() ([]Warehouse, error)
func (s *service) Delete(id int) error
