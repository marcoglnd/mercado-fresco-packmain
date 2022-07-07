package repository

import (
	"context"
	"database/sql"
	"errors"

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
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.WarehouseCode,
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

	warehouse.ID = lastID

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
	id int64,
) (*domain.Warehouse, error) {
	row := r.db.QueryRowContext(
		ctx, sqlGetById, id,
	)

	foundWarehouse := &domain.Warehouse{}
	err := row.Scan(
		&foundWarehouse.ID,
		&foundWarehouse.Address,
		&foundWarehouse.Telephone,
		&foundWarehouse.WarehouseCode,
		&foundWarehouse.MinimumCapacity,
		&foundWarehouse.MinimumTemperature,
		&foundWarehouse.LocalityId,
	)

	if err != nil {
		return nil, err
	}

	return foundWarehouse, nil
}

func (r *warehouseRepository) FindByWarehouseCode(
	ctx context.Context,
	warehouseCode string,
) (*domain.Warehouse, error) {
	row := r.db.QueryRowContext(
		ctx, sqlGetByWarehouseCode, warehouseCode,
	)

	foundWarehouse := &domain.Warehouse{}
	err := row.Scan(
		&foundWarehouse.ID,
		&foundWarehouse.Address,
		&foundWarehouse.Telephone,
		&foundWarehouse.WarehouseCode,
		&foundWarehouse.MinimumCapacity,
		&foundWarehouse.MinimumTemperature,
		&foundWarehouse.LocalityId,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return foundWarehouse, nil
}
func (r *warehouseRepository) GetAll(
	ctx context.Context,
) (*[]domain.Warehouse, error) {
	warehouses := []domain.Warehouse{}

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &warehouses, err
	}

	defer rows.Close()

	for rows.Next() {
		var warehouse domain.Warehouse

		if err := rows.Scan(
			&warehouse.ID,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.WarehouseCode,
			&warehouse.MinimumCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityId,
		); err != nil {
			return &warehouses, err
		}

		warehouses = append(warehouses, warehouse)
	}

	return &warehouses, nil
}
func (r *warehouseRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	result, err := r.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
