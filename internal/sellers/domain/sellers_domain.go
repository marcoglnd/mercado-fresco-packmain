package domain

import "context"

// Modelo de sellers
type Seller struct {
	ID           int64  `json:"id"`
	Cid          int64  `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
}

type SellerRepository interface {
	GetAll(ctx context.Context) (*[]Seller, error)
	GetByID(ctx context.Context, id int64) (*Seller, error)
	Create(ctx context.Context, seller *Seller) (*Seller, error)
	Update(ctx context.Context, seller *Seller) (*Seller, error)
	Delete(ctx context.Context, id int64) error
}

type SellerService interface {
	GetAll(ctx context.Context) (*[]Seller, error)
	GetByID(ctx context.Context, id int64) (*Seller, error)
	Create(ctx context.Context, seller *Seller) (*Seller, error)
	Update(ctx context.Context, seller *Seller) (*Seller, error)
	Delete(ctx context.Context, id int64) error
}
