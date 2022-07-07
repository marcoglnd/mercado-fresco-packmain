package service

import (
	"context"
	"fmt"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/domain"
)

type warehouseService struct {
	repository domain.WarehouseRepository
}

func NewWarehouseService(wr domain.WarehouseRepository) domain.WarehouseService {
	return &warehouseService{repository: wr}
}

func (s *warehouseService) Create(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {

	if err := s.IsWarehouseCodeAvailable(ctx, warehouse.WarehouseCode); err != nil {
		return nil, err
	}

	warehouse, err := s.repository.Create(ctx, warehouse)

	if err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (s *warehouseService) IsWarehouseCodeAvailable(ctx context.Context, warehouseCode string) error {
	warehouseDuplicated, err := s.repository.FindByWarehouseCode(ctx, warehouseCode)
	if err != nil {
		return err
	}
	if warehouseDuplicated != nil {
		return fmt.Errorf("warehouseCode already exists")
	}
	return nil
}

func (s *warehouseService) Update(ctx context.Context, updatedWarehouse *domain.Warehouse) (*domain.Warehouse, error) {
	currentWarehouse, err := s.repository.FindById(ctx, updatedWarehouse.ID)
	if err != nil {
		return nil, err
	}

	if updatedWarehouse.WarehouseCode != currentWarehouse.WarehouseCode &&
		updatedWarehouse.WarehouseCode != "" {
		if err := s.IsWarehouseCodeAvailable(ctx, updatedWarehouse.WarehouseCode); err != nil {
			return nil, err
		} else {
			currentWarehouse.WarehouseCode = updatedWarehouse.WarehouseCode
		}
	}

	if updatedWarehouse.Address != currentWarehouse.Address &&
		updatedWarehouse.Address != "" {
		currentWarehouse.Address = updatedWarehouse.Address
	}

	if updatedWarehouse.Telephone != currentWarehouse.Telephone &&
		updatedWarehouse.Telephone != "" {
		currentWarehouse.Telephone = updatedWarehouse.Telephone
	}

	if updatedWarehouse.MinimumCapacity != currentWarehouse.MinimumCapacity &&
		updatedWarehouse.MinimumCapacity != 0 {
		currentWarehouse.MinimumCapacity = updatedWarehouse.MinimumCapacity
	}

	if updatedWarehouse.MinimumTemperature != currentWarehouse.MinimumTemperature &&
		updatedWarehouse.MinimumTemperature != 0 {
		currentWarehouse.MinimumTemperature = updatedWarehouse.MinimumTemperature
	}

	if err := s.repository.Update(ctx, currentWarehouse); err != nil {
		return nil, err
	}

	return currentWarehouse, nil
}

func (s *warehouseService) FindById(ctx context.Context, id int64) (*domain.Warehouse, error) {
	foundWarehouse, err := s.repository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	if foundWarehouse == nil {
		return nil, fmt.Errorf("could not find warehouse by id")
	}

	return foundWarehouse, nil
}

func (s *warehouseService) FindByWarehouseCode(ctx context.Context, warehouseCode string) (*domain.Warehouse, error) {
	foundWarehouse, err := s.repository.FindByWarehouseCode(ctx, warehouseCode)

	if err != nil {
		return nil, err
	}

	return foundWarehouse, nil
}

func (s *warehouseService) GetAll(ctx context.Context) (*[]domain.Warehouse, error) {

	warehouses, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return warehouses, nil
}

func (s *warehouseService) Delete(ctx context.Context, id int64) error {

	if _, err := s.FindById(ctx, id); err != nil {
		return err
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
