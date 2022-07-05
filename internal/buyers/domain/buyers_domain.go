package domain

import (
	"context"
)

type Buyer struct {
	ID           int64  `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerRepository interface {
	GetAll(ctx context.Context) (*[]Buyer, error)
	GetById(ctx context.Context, id int64) (*Buyer, error)
	GetByCardNumberId(ctx context.Context, cardNumberId string) (*Buyer, error)
	Create(ctx context.Context, cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(ctx context.Context, id int64, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int64) error
}

type BuyerService interface {
	GetAll(ctx context.Context) (*[]Buyer, error)
	GetById(ctx context.Context, id int64) (*Buyer, error)
	Create(ctx context.Context, cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(ctx context.Context, id int64, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int64) error
}
