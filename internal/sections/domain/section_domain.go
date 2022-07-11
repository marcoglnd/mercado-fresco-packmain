package domain

import (
	"context"
)

type Section struct {
	ID                 int64   `json:"id"`
	SectionNumber      int64   `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int64   `json:"current_capacity"`
	MinimumCapacity    int64   `json:"minimum_capacity"`
	MaximumCapacity    int64 `json:"maximum_capacity"`
	WarehouseId        int64 `json:"warehouse_id"`
	ProductTypeId      int64 `json:"product_type_id"`
}

type Service interface {
	GetAll(ctx context.Context) (*[]Section, error)
	GetById(ctx context.Context, id int64) (*Section, error)
	Create(ctx context.Context, section *Section) (*Section, error)
	Update(ctx context.Context, section *Section) (*Section, error)
	Delete(ctx context.Context, id int64) error
}

type Repository interface {
	GetAll(ctx context.Context) (*[]Section, error)
	GetById(ctx context.Context, id int64) (*Section, error)
	Create(ctx context.Context, section *Section) (*Section, error)
	Update(ctx context.Context, section *Section) (*Section, error)
	Delete(ctx context.Context, id int64) error
}

type RequestSectionId struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type RequestSections struct {
	SectionNumber      int64   `json:"section_number" binding:"required"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int64   `json:"current_capacity" binding:"required"`
	MinimumCapacity    int64   `json:"minimum_capacity" binding:"required"`
	MaximumCapacity int64 `json:"maximum_capacity" binding:"required"`
	WarehouseId     int64 `json:"warehouse_id" binding:"required"`
	ProductTypeId   int64 `json:"product_type_id" binding:"required"`
}

type RequestSectionsUpdated struct {
	SectionNumber      int64   `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int64   `json:"current_capacity"`
	MinimumCapacity    int64   `json:"minimum_capacity"`
	MaximumCapacity int64 `json:"maximum_capacity"`
	WarehouseId     int64 `json:"warehouse_id"`
	ProductTypeId   int64 `json:"product_type_id"`
}
