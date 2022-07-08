package domain

import "context"

// Modelo de sellers
type Seller struct {
	ID           int64  `json:"id"`
	Cid          int64  `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	LocalityID   int64  `json:"locality_id"`
}

type SellerRepository interface {
	GetAll(ctx context.Context) (*[]Seller, error)
	GetByID(ctx context.Context, id int64) (*Seller, error)
	Create(ctx context.Context, seller *Seller) (*Seller, error)
	Update(ctx context.Context, seller *Seller) (*Seller, error)
	Delete(ctx context.Context, id int64) error
	// CreateLocality(ctx context.Context, local *Locality) (int64, error)
	// GetLocalityByID(ctx context.Context, id int64) (*GetLocality, error)
	// GetQtyOfSellers(ctx context.Context) (*[]QtyOfSellers, error)
	// GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*QtyOfSellers, error)
}

type SellerService interface {
	GetAll(ctx context.Context) (*[]Seller, error)
	GetByID(ctx context.Context, id int64) (*Seller, error)
	Create(ctx context.Context, seller *Seller) (*Seller, error)
	Update(ctx context.Context, seller *Seller) (*Seller, error)
	Delete(ctx context.Context, id int64) error
	// CreateLocality(ctx context.Context, local *Locality) (int64, error)
	// GetLocalityByID(ctx context.Context, id int64) (*GetLocality, error)
	// GetQtyOfSellers(ctx context.Context) (*[]QtyOfSellers, error)
	// GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*QtyOfSellers, error)
}

// type Locality struct {
// 	LocalityName string `json:"locality_name"`
// 	ProvinceID   int64  `json:"province_id"`
// }

// type GetLocality struct {
// 	ID           int64  `json:"ID"`
// 	LocalityName string `json:"locality_name"`
// 	ProvinceName string `json:"province_name"`
// 	CountryName  string `json:"country_name"`
// }

// type QtyOfSellers struct {
// 	LocalityID   int64  `json:"locality_id"`
// 	LocalityName string `json:"locality_name"`
// 	SellersCount int64  `json:"sellers_count"`
// }
