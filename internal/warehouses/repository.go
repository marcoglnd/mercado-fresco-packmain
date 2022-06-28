package warehouses

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var warehouses []Warehouse = []Warehouse{}

//go:generate mockgen -source=./repository.go -destination=./mocks/repository_mock.go
type Repository interface {
	Create(
		warehouseCode string,
		address string,
		telephone string,
		minimumCapacity int,
		minimumTemperature float32,
	) (*Warehouse, error)
	Update(warehouse *Warehouse) error
	FindById(id int) (*Warehouse, error)
	FindByWarehouseCode(warehouseCode string) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(
	warehouseCode string,
	address string,
	telephone string,
	minimumCapacity int,
	minimumTemperature float32,
) (*Warehouse, error) {
	warehouse := Warehouse{
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
		LocalityId:         1,
	}
	result := r.db.Create(&warehouse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &warehouse, nil
}
func (r *repository) Update(warehouse *Warehouse) error {
	result := r.db.Save(&warehouse)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) FindById(id int) (*Warehouse, error) {
	foundWarehouse := &Warehouse{}
	result := r.db.First(foundWarehouse, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return foundWarehouse, nil
}
func (r *repository) FindByWarehouseCode(warehouseCode string) (*Warehouse, error) {
	foundWarehouse := &Warehouse{}
	result := r.db.First(foundWarehouse, "warehouse_code = ?", warehouseCode)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return foundWarehouse, nil
}
func (r *repository) GetAll() ([]Warehouse, error) {
	warehouses := []Warehouse{}
	result := r.db.Find(&warehouses)
	if result.Error != nil {
		return nil, result.Error
	}
	return warehouses, nil
}
func (r *repository) Delete(id int) error {
	result := r.db.Model(&Warehouse{}).Delete("id", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewRepository() Repository {
	dsn := "root:pass@tcp(127.0.0.1:3306)/mercado_fresco?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &repository{
		db: db,
	}
}
