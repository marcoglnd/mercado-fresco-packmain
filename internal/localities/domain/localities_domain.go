package domain

import "context"

// Modelo de sellers

type LocalityRepository interface {
	CreateLocality(ctx context.Context, local *Locality) (int64, error)
	GetLocalityByID(ctx context.Context, id int64) (*GetLocality, error)
	GetAllQtyOfSellers(ctx context.Context) (*[]QtyOfSellers, error)
	GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*QtyOfSellers, error)
}

type LocalityService interface {
	CreateLocality(ctx context.Context, local *Locality) (int64, error)
	GetLocalityByID(ctx context.Context, id int64) (*GetLocality, error)
	GetAllQtyOfSellers(ctx context.Context) (*[]QtyOfSellers, error)
	GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*QtyOfSellers, error)
}

type Locality struct {
	LocalityName string `json:"locality_name"`
	ProvinceID   int64  `json:"province_id"`
}

type GetLocality struct {
	ID           int64  `json:"ID"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type QtyOfSellers struct {
	LocalityID   int64  `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int64  `json:"sellers_count"`
}
