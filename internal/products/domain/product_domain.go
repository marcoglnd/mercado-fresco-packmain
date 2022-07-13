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
	NetWeight                      float64 `json:"net_weight"`
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

	GetQtyOfRecordsById(ctx context.Context, id int64) (*QtyOfRecords, error)
	GetQtyOfAllRecords(ctx context.Context) (*[]QtyOfRecords, error)

	CreateProductBatches(ctx context.Context, batch *ProductBatches) (int64, error)
	GetProductBatchesById(ctx context.Context, id int64) (*ProductBatches, error)

	GetQtdProductsBySectionId(ctx context.Context, id int64) (*QtdOfProducts, error)
	GetQtdOfAllProducts(ctx context.Context) (*[]QtdOfProducts, error)
}

type Service interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int64) error

	CreateProductRecords(ctx context.Context, record *ProductRecords) (int64, error)
	GetProductRecordsById(ctx context.Context, id int64) (*ProductRecords, error)

	GetQtyOfRecordsById(ctx context.Context, id int64) (*QtyOfRecords, error)
	GetQtyOfAllRecords(ctx context.Context) (*[]QtyOfRecords, error)

	CreateProductBatches(ctx context.Context, batch *ProductBatches) (int64, error)
	GetProductBatchesById(ctx context.Context, id int64) (*ProductBatches, error)

	GetQtdProductsBySectionId(ctx context.Context, id int64) (*QtdOfProducts, error)
	GetQtdOfAllProducts(ctx context.Context) (*[]QtdOfProducts, error)
}

type RequestProducts struct {
	Description                    string  `json:"description" binding:"required"`
	ExpirationRate                 int64   `json:"expiration_rate" binding:"required"`
	FreezingRate                   int64   `json:"freezing_rate" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"net_weight" binding:"required"`
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
	NetWeight                      float64 `json:"net_weight"`
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
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	SalePrice     float64 `json:"sale_price" binding:"required"`
	ProductId     int64   `json:"product_id" binding:"required"`
}

type RequestProductRecordId struct {
	Id int64 `form:"id" binding:"required,min=1"`
}
type QtyOfRecords struct {
	ProductId    int64  `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int64  `json:"records_count"`
}

type RequestProductBatches struct {
	BatchNumber        int64   `json:"batch_number" binding:"required"`
	CurrentQuantity    int64   `json:"current_quantity" binding:"required"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required"`
	DueDate            string  `json:"due_date" binding:"required"`
	InitialQuantity    int64   `json:"initial_quantity" binding:"required"`
	ManufacturingDate  string  `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  int64   `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required"`
	ProductId          int64   `json:"product_id" binding:"required"`
	SectionId          int64   `json:"section_id" binding:"required"`
}

type ProductBatches struct {
	BatchNumber        int64   `json:"batch_number"`
	CurrentQuantity    int64   `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int64   `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  int64   `json:"manufacturing_hour"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	ProductId          int64   `json:"product_id"`
	SectionId          int64   `json:"section_id"`
}

type RequestQtdProductsBySectionId struct {
	Id int64 `form:"id" binding:"required,min=1"`
}
type QtdOfProducts struct {
	SectionId     int64 `json:"section_id"`
	SectionNumber int64 `json:"section_number"`
	ProductsCount int64 `json:"products_count"`
}
