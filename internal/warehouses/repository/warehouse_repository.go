package repository

import (
	"context"
	"database/sql"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/domain"
)

type warehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) domain.WarehouseRepository {
	return &warehouseRepository{db: db}
}

func (r *warehouseRepository) Create(
	ctx context.Context,
	warehouse *domain.Warehouse,
) (*domain.Warehouse, error) {
	result, err := r.db.ExecContext(
		ctx,
		sqlStore,
		&warehouse.WarehouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimumCapacity,
		&warehouse.MinimumTemperature,
	)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	warehouse.ID = int(lastID)

	return warehouse, nil
}

func (r *warehouseRepository) Update(
	ctx context.Context,
	warehouse *domain.Warehouse,
) error {
	_, err := r.db.ExecContext(
		ctx,
		sqlUpdate,
		&warehouse.WarehouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimumCapacity,
		&warehouse.MinimumTemperature,
		&warehouse.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *warehouseRepository) FindById(
	ctx context.Context,
	id int,
) (*domain.Warehouse, error) {
	var foundWarehouse *domain.Warehouse
	return foundWarehouse, nil
}

func (r *warehouseRepository) FindByWarehouseCode(
	ctx context.Context,
	warehouseCode string,
) (*domain.Warehouse, error) {
	var foundWarehouse *domain.Warehouse
	return foundWarehouse, nil
}
func (r *warehouseRepository) GetAll(
	ctx context.Context,
) (*[]domain.Warehouse, error) {
	return nil, nil
}
func (r *warehouseRepository) Delete(
	ctx context.Context,
	id int,
) error {
	return nil
}
