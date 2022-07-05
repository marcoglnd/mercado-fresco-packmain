package domain

import (
	"context"
)

type Product struct {
	Id                             int64   `json:"id"`
	Description                    string  `json:"description"`
	ExpirationRate                 int64   `json:"expiration_rate"`
	FreezingRate                   int64   `json:"freezing_rate"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netweight"`
	ProductCode                    string  `json:"product_code"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	Width                          float64 `json:"width"`
	ProductTypeId                  int64   `json:"product_type_id"`
	SellerId                       int64   `json:"seller_id"`
}

type Repository interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int64) error

	CreateProductRecords(ctx context.Context, record *ProductRecords) (int64, error)
	GetProductRecordsById(ctx context.Context, id int64) (*ProductRecords, error)

	GetQtyOfRecords(ctx context.Context, id int64) (*QtyOfRecords, error)
}

type Service interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int64) error

	CreateProductRecords(ctx context.Context, record *ProductRecords) (int64, error)
	GetProductRecordsById(ctx context.Context, id int64) (*ProductRecords, error)

	GetQtyOfRecords(ctx context.Context, id int64) (*QtyOfRecords, error)
}

type RequestProducts struct {
	Description                    string  `json:"description" binding:"required"`
	ExpirationRate                 int64   `json:"expiration_rate" binding:"required"`
	FreezingRate                   int64   `json:"freezing_rate" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"netweight" binding:"required"`
	ProductCode                    string  `json:"product_code" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	ProductTypeId                  int64   `json:"product_type_id" binding:"required"`
	SellerId                       int64   `json:"seller_id" binding:"required"`
}

type RequestProductId struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type RequestProductsUpdated struct {
	Description                    string  `json:"description"`
	ExpirationRate                 int64   `json:"expiration_rate"`
	FreezingRate                   int64   `json:"freezing_rate"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netweight"`
	ProductCode                    string  `json:"product_code"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	Width                          float64 `json:"width"`
	ProductTypeId                  int64   `json:"product_type_id"`
	SellerId                       int64   `json:"seller_id"`
}

type ProductRecords struct {
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int64   `json:"product_id"`
}

type RequestProductRecords struct {
	PurchasePrice float64 `json:"purchase_price"`
	SalePrice     float64 `json:"sale_price"`
	ProductId     int64   `json:"product_id"`
}

type RequestProductRecordId struct {
	Id int64 `form:"id" binding:"required,min=1"`
}

type QtyOfRecords struct {
	ProductId    int64  `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int64  `json:"records_count"`
}
