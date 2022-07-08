package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
)

type localityService struct {
	repository domain.LocalityRepository
}

func NewService(r domain.LocalityRepository) domain.LocalityService {
	return &localityService{
		repository: r,
	}
}

func (s localityService) CreateLocality(ctx context.Context, local *domain.Locality) (int64, error) {
	newLocality, err := s.repository.CreateLocality(ctx, local)
	if err != nil {
		return newLocality, err
	}
	return newLocality, nil
}

func (s localityService) GetLocalityByID(ctx context.Context, id int64) (*domain.GetLocality, error) {
	getLocality, err := s.repository.GetLocalityByID(ctx, id)
	if err != nil {
		return getLocality, err
	}
	return getLocality, nil
}

func (s localityService) GetQtyOfSellers(ctx context.Context) (*[]domain.QtyOfSellers, error) {
	listOfSellers, err := s.repository.GetQtyOfSellers(ctx)
	if err != nil {
		return listOfSellers, err
	}
	return listOfSellers, nil
}

func (s localityService) GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*domain.QtyOfSellers, error) {
	getSellersByLocalityID, err := s.repository.GetQtyOfSellersByLocalityId(ctx, id)
	if err != nil {
		return getSellersByLocalityID, err
	}
	return getSellersByLocalityID, nil
}
