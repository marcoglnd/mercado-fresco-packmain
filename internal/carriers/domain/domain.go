package domain

import "context"

//go:generate mockgen -source=./domain.go -destination=../mocks/domain.go

type CarrierRepository interface {
	Create(ctx context.Context, carrier *Carrier) (*Carrier, error)
	FindById(ctx context.Context, id int64) (*Carrier, error)
	FindByCid(ctx context.Context, cid string) (*Carrier, error)
	GetAll(ctx context.Context) (*[]Carrier, error)
	GetAllCarriersReport(ctx context.Context) (*[]CarrierReport, error)
	GetCarriersReportById(ctx context.Context, id int64) (*CarrierReport, error)
}

type CarrierService interface {
	Create(ctx context.Context, carrier *Carrier) (*Carrier, error)
	FindById(ctx context.Context, id int64) (*Carrier, error)
	FindByCid(ctx context.Context, cid string) (*Carrier, error)
	IsCidAvailable(ctx context.Context, cid string) error
	GetAllCarriersReport(ctx context.Context) (*[]CarrierReport, error)
	GetCarriersReportById(ctx context.Context, id int64) (*CarrierReport, error)
}
