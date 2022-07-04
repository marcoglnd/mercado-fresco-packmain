package domain

import "context"

//go:generate mockgen -source=./domain.go -destination=../mocks/domain.go
type WarehouseRepository interface {
	Create(ctx context.Context, warehouse *Warehouse) (*Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	FindById(ctx context.Context, id int64) (*Warehouse, error)
	FindByWarehouseCode(ctx context.Context, warehouseCode string) (*Warehouse, error)
	GetAll(ctx context.Context) (*[]Warehouse, error)
	Delete(ctx context.Context, id int64) error
}

type WarehouseService interface {
	Create(ctx context.Context, warehouse *Warehouse) (*Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) (*Warehouse, error)
	FindById(ctx context.Context, id int64) (*Warehouse, error)
	FindByWarehouseCode(ctx context.Context, warehouseCode string) (*Warehouse, error)
	IsWarehouseCodeAvailable(ctx context.Context, warehouseCode string) error
	GetAll(ctx context.Context) (*[]Warehouse, error)
	Delete(ctx context.Context, id int64) error
}
