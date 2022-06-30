package products

import (
	"context"
	"errors"
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
	GetById(ctx context.Context, id int) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int) error
}

type Service interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int) (*Product, error)
	CreateNewProduct(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int) error
}

var (
	ErrIDNotFound = errors.New("section id not found")
)
