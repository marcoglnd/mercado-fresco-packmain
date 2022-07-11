package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll(ctx context.Context) (*[]domain.Section, error) {
	sectionsList, err := s.repository.GetAll(ctx)
	if err != nil {
		return sectionsList, err
	}
	return sectionsList, nil
}

func (s service) GetById(ctx context.Context, id int64) (*domain.Section, error) {
	section, err := s.repository.GetById(ctx, id)
	if err != nil {
		return section, err
	}
	return section, err
}

func (s *service) Create(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	section, err := s.repository.Create(ctx, section)

	if err != nil {
		return section, err
	}

	return section, nil
}

func (s *service) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	section, err := s.repository.Update(ctx, section)
	if err != nil {
		return section, err
	}
	return section, err
}

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return err
}
