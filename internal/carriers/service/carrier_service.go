package service

import (
	"context"
	"fmt"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"
)

type carrierService struct {
	repository domain.CarrierRepository
}

func NewCarrierService(cr domain.CarrierRepository) domain.CarrierService {
	return &carrierService{repository: cr}
}

func (s *carrierService) Create(ctx context.Context, carrier *domain.Carrier) (*domain.Carrier, error) {

	if err := s.IsCidAvailable(ctx, carrier.Cid); err != nil {
		return nil, err
	}

	carrier, err := s.repository.Create(ctx, carrier)

	if err != nil {
		return nil, err
	}

	return carrier, nil
}

func (s *carrierService) IsCidAvailable(ctx context.Context, cid string) error {
	carrierDuplicated, err := s.repository.FindByCid(ctx, cid)
	if err != nil {
		return err
	}
	if carrierDuplicated != nil {
		return fmt.Errorf("cid already exists")
	}
	return nil
}

func (s *carrierService) FindById(ctx context.Context, id int64) (*domain.Carrier, error) {
	foundCarrier, err := s.repository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	if foundCarrier == nil {
		return nil, fmt.Errorf("could not find carrier by id")
	}

	return foundCarrier, nil
}

func (s *carrierService) FindByCid(ctx context.Context, cid string) (*domain.Carrier, error) {
	foundCarrier, err := s.repository.FindByCid(ctx, cid)

	if err != nil {
		return nil, err
	}

	return foundCarrier, nil
}

func (s *carrierService) GetAllCarriersReport(
	ctx context.Context,
) (*[]domain.CarrierReport, error) {
	carriersReport, err := s.repository.GetAllCarriersReport(ctx)

	if err != nil {
		return nil, err
	}

	return carriersReport, nil
}

func (s *carrierService) GetCarriersReportById(
	ctx context.Context,
	id int64,
) (*domain.CarrierReport, error) {
	carrierReport, err := s.repository.GetCarriersReportById(ctx, id)

	if err != nil {
		return nil, err
	}

	return carrierReport, nil
}
