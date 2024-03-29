package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Product, error) {
	listOfProducts, err := s.repository.GetAll(ctx)
	if err != nil {
		return listOfProducts, err
	}

	return listOfProducts, nil
}

func (s service) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	product, err := s.repository.GetById(ctx, id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	newProd, err := s.repository.CreateNewProduct(ctx, product)
	if err != nil {
		return newProd, err
	}
	return newProd, nil
}

func (s *service) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	current, err := s.GetById(ctx, product.Id)
	if err != nil {
		return product, err
	}

	if len(product.Description) > 0 {
		current.Description = product.Description
	}

	if product.ExpirationRate > 0 || product.ExpirationRate < 0 {
		current.ExpirationRate = product.ExpirationRate
	}

	if product.FreezingRate > 0 || product.FreezingRate < 0 {
		current.FreezingRate = product.FreezingRate
	}

	if product.Height > 0 {
		current.Height = product.Height
	}

	if product.Length > 0 {
		current.Length = product.Length
	}

	if product.Width > 0 {
		current.Width = product.Width
	}

	if product.NetWeight > 0 {
		current.NetWeight = product.NetWeight
	}

	if product.ProductTypeId > 0 {
		current.ProductTypeId = product.ProductTypeId
	}

	if product.RecommendedFreezingTemperature > 0 || product.RecommendedFreezingTemperature < 0 {
		current.RecommendedFreezingTemperature = product.RecommendedFreezingTemperature
	}

	if product.SellerId > 0 {
		current.SellerId = product.SellerId
	}

	if len(product.ProductCode) > 0 {
		current.ProductCode = product.ProductCode
	}

	product, err = s.repository.Update(ctx, current)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateProductRecords(ctx context.Context, record *domain.ProductRecords) (int64, error) {
	newRecordId, err := s.repository.CreateProductRecords(ctx, record)
	if err != nil {
		return newRecordId, err
	}
	return newRecordId, nil
}

func (s *service) GetProductRecordsById(ctx context.Context, id int64) (*domain.ProductRecords, error) {
	newRecord, err := s.repository.GetProductRecordsById(ctx, id)
	if err != nil {
		return newRecord, err
	}
	return newRecord, nil
}

func (s service) GetQtyOfRecordsById(ctx context.Context, id int64) (*domain.QtyOfRecords, error) {
	report, err := s.repository.GetQtyOfRecordsById(ctx, id)
	if err != nil {
		return report, err
	}

	return report, nil
}

func (s service) GetQtyOfAllRecords(ctx context.Context) (*[]domain.QtyOfRecords, error) {
	report, err := s.repository.GetQtyOfAllRecords(ctx)
	if err != nil {
		return report, err
	}

	return report, nil
}

func (s *service) CreateProductBatches(ctx context.Context, batche *domain.ProductBatches) (int64, error) {
	newBatchId, err := s.repository.CreateProductBatches(ctx, batche)
	if err != nil {
		return newBatchId, err
	}
	return newBatchId, nil
}

func (s *service) GetProductBatchesById(ctx context.Context, id int64) (*domain.ProductBatches, error) {
	newBatch, err := s.repository.GetProductBatchesById(ctx, id)
	if err != nil {
		return newBatch, err
	}
	return newBatch, nil
}

func (s service) GetQtdProductsBySectionId(ctx context.Context, id int64) (*domain.QtdOfProducts, error) {
	report, err := s.repository.GetQtdProductsBySectionId(ctx, id)
	if err != nil {
		return report, err
	}

	return report, nil
}

func (s service) GetQtdOfAllProducts(ctx context.Context) (*[]domain.QtdOfProducts, error) {
	report, err := s.repository.GetQtdOfAllProducts(ctx)
	if err != nil {
		return report, err
	}

	return report, nil
}
