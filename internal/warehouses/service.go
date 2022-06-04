package warehouses

import (
	"fmt"
)

type Service interface {
	Create(
		warehouseCode string,
		address string,
		telephone string,
		minimumCapacity int,
		minimumTemperature int,
	) (*Warehouse, error)
	Update(
		id int,
		warehouseCode string,
		address string,
		telephone string,
		minimumCapacity int,
		minimumTemperature int,
	) (*Warehouse, error)
	FindById(id int) (*Warehouse, error)
	FindByWarehouseCode(warehouseCode string) (*Warehouse, error)
	IsWarehouseCodeAvailable(warehouseCode string) error
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

	if err := s.IsWarehouseCodeAvailable(warehouseCode); err != nil {
		return &Warehouse{}, err
	}

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

func (s *service) IsWarehouseCodeAvailable(warehouseCode string) error {
	warehouseDuplicated, err := s.FindByWarehouseCode(warehouseCode)
	if err != nil {
		return err
	}
	if warehouseDuplicated != nil {
		return fmt.Errorf("warehouseCode already exists")
	}
	return nil
}

func (s *service) Update(
	id int,
	warehouseCode string,
	address string,
	telephone string,
	minimumCapacity int,
	minimumTemperature int,
) (*Warehouse, error) {

	currentW, err := s.FindById(id)
	if err != nil {
		return &Warehouse{}, err
	}

	if warehouseCode != currentW.WarehouseCode &&
		warehouseCode != "" {
		if err := s.IsWarehouseCodeAvailable(warehouseCode); err != nil {
			return &Warehouse{}, err
		} else {
			currentW.WarehouseCode = warehouseCode
		}
	}

	updatedWarehouse := currentW

	if address != currentW.Address &&
		address != "" {
		updatedWarehouse.Address = address
	}

	if telephone != currentW.Telephone &&
		telephone != "" {
		updatedWarehouse.Telephone = telephone
	}

	if minimumCapacity != currentW.MinimumCapacity &&
		minimumCapacity != 0 {
		updatedWarehouse.MinimumCapacity = minimumCapacity
	}

	if minimumTemperature != currentW.MinimumTemperature &&
		minimumTemperature != 0 {
		updatedWarehouse.MinimumTemperature = minimumTemperature
	}

	if err := s.repository.Update(updatedWarehouse); err != nil {
		return &Warehouse{}, err
	}

	return updatedWarehouse, nil
}

func (s *service) FindById(id int) (*Warehouse, error) {
	foundWarehouse, err := s.repository.FindById(id)

	if err != nil {
		return &Warehouse{}, err
	}

	if foundWarehouse == nil {
		return &Warehouse{}, fmt.Errorf("could not find warehouse by id")
	}

	return foundWarehouse, nil
}

func (s *service) FindByWarehouseCode(warehouseCode string) (*Warehouse, error) {
	foundWarehouse, err := s.repository.FindByWarehouseCode(warehouseCode)

	if err != nil {
		return &Warehouse{}, err
	}

	return foundWarehouse, nil
}

func (s *service) GetAll() ([]Warehouse, error) {

	warehouses, err := s.repository.GetAll()

	if err != nil {
		return []Warehouse{}, err
	}

	return warehouses, nil
}

func (s *service) Delete(id int) error {
	return nil
}
