package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
)

type sellerService struct {
	repository domain.SellerRepository
}

func NewService(r domain.SellerRepository) domain.SellerService {
	return &sellerService{
		repository: r,
	}
}

func (s sellerService) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	sellers, err := s.repository.GetAll(ctx)
	if err != nil {
		return sellers, err
	}
	return sellers, nil
}

func (s sellerService) GetByID(ctx context.Context, id int64) (*domain.Seller, error) {
	seller, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func (s sellerService) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	seller, err := s.repository.Create(ctx, seller)
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func (s sellerService) Update(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	seller, err := s.repository.Update(ctx, seller)
	if err != nil {
		return seller, err
	}
	return seller, err
}

func (s sellerService) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
