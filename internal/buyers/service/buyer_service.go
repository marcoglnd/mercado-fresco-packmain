package service

import (
	"context"
	"fmt"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
)

type buyerService struct {
	repository domain.BuyerRepository
}

func NewBuyerService(sr domain.BuyerRepository) domain.BuyerService {
	return &buyerService{repository: sr}
}

func (s buyerService) GetAll(ctx context.Context) (*[]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(ctx)
	if err != nil {
		return buyers, err
	}

	return buyers, nil
}

func (s buyerService) GetById(ctx context.Context, id int64) (*domain.Buyer, error) {
	buyer, err := s.repository.GetById(ctx, id)
	if err != nil {
		return buyer, err
	}

	return buyer, nil
}

func (s buyerService) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	buyersList, err := s.repository.GetAll(ctx)
	if err != nil {
		return &domain.Buyer{}, err
	}

	for i := range buyersList {
		if buyersList[i].CardNumberID == cardNumberId {
			return Buyer{}, fmt.Errorf("CardNumberID %s do Buyer já existe", cardNumberId)
		}
	}

	buyer, err := s.repository.Create(ctx, cardNumberId, firstName, lastName)
	if err != nil {
		return buyer, err
	}

	return buyer, nil
}

func (s buyerService) Update(ctx context.Context, id int64, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	buyer, err := s.repository.Update(ctx, id, cardNumberId, firstName, lastName)
	if err != nil {
		return buyer, err
	}

	return buyer, nil
}

func (s buyerService) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
