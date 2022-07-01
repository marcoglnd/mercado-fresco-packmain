package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewProduct(t *testing.T) {
	mockProductsRepo := mocks.NewService(t)
	mockProduct := utils.CreateRandomProduct()

	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo.On("CreateNewProduct",
			mock.Anything,
			mock.Anything,
		).Return(&mockProduct, nil).Once()

		s := NewService(mockProductsRepo)

		newProduct, err := s.CreateNewProduct(context.Background(), &mockProduct)

		assert.NoError(t, err)
		assert.Equal(t, &mockProduct, newProduct)

		mockProductsRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo.On("CreateNewProduct",
			mock.Anything,
			mock.Anything,
		).Return(&domain.Product{}, errors.New("failed to create product")).Once()

		s := NewService(mockProductsRepo)

		_, err := s.CreateNewProduct(context.Background(), &mockProduct)

		assert.Error(t, err)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockProductsRepo := mocks.NewService(t)

	mockProducts := utils.CreateRandomListProduct()

	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo.On("GetAll", mock.Anything).
			Return(&mockProducts, nil).Once()

		s := NewService(mockProductsRepo)
		list, err := s.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockProducts, list)

		mockProductsRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo.On("GetAll", mock.Anything).
			Return(nil, errors.New("failed to retrieve products")).
			Once()

		s := NewService(mockProductsRepo)
		_, err := s.GetAll(context.Background())

		assert.NotNil(t, err)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockProductsRepo := mocks.NewService(t)

	mockProduct := utils.CreateRandomProduct()

	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct, nil).Once()

		service := NewService(mockProductsRepo)

		product, err := service.GetById(context.Background(), mockProduct.Id)

		assert.NoError(t, err)
		assert.NotEmpty(t, product)

		assert.Equal(t, &mockProduct, product)

		mockProductsRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve product")).Once()

		service := NewService(mockProductsRepo)

		product, err := service.GetById(context.Background(), mockProduct.Id)

		assert.Error(t, err)
		assert.Empty(t, product)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockProductsRepo := mocks.NewService(t)

	mockProduct := utils.CreateRandomProduct()

	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo.On(
			"GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).
			Return(&mockProduct, nil).On(
			"Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockProduct, nil).Once()

		service := NewService(mockProductsRepo)
		product, err := service.Update(
			context.Background(), &mockProduct,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, product)

		assert.Equal(t, &mockProduct, product)

		mockProductsRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo.On(
			"GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).
			Return(&mockProduct, nil).On(
			"Update",
			mock.Anything,
			mock.Anything,
		).Return(nil,  errors.New("failed to retrieve product")).Once()

		service := NewService(mockProductsRepo)
		product, err := service.Update(
			context.Background(), &mockProduct,
		)
		assert.Error(t, err)
		assert.Empty(t, product)

		mockProductsRepo.AssertExpectations(t)
	})
}
