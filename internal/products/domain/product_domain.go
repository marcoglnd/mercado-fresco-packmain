package domain

import (
	"context"
)

type Product struct {
	Id                             int64   `json:"id"`
	Description                    string  `json:"description"`
	ExpirationRate                 int     `json:"expiration_rate"`
	FreezingRate                   int     `json:"freezing_rate"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netweight"`
	ProductCode                    string  `json:"product_code"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	Width                          float64 `json:"width"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type Repository interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int64) error
}

type Service interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int64) error
}

type RequestProducts struct {
	Description                    string  `json:"description" binding:"required"`
	ExpirationRate                 int     `json:"expiration_rate" binding:"required"`
	FreezingRate                   int     `json:"freezing_rate" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"netweight" binding:"required"`
	ProductCode                    string  `json:"product_code" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id" binding:"required"`
}

type RequestProductId struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type RequestProductsUpdated struct {
	Description                    string  `json:"description"`
	ExpirationRate                 int     `json:"expiration_rate"`
	FreezingRate                   int     `json:"freezing_rate"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netweight"`
	ProductCode                    string  `json:"product_code"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	Width                          float64 `json:"width"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}