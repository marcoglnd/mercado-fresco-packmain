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
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

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
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

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
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProducts := utils.CreateRandomListProduct()

		mockProductsRepo.On("GetAll", mock.Anything).
			Return(&mockProducts, nil).Once()

		s := NewService(mockProductsRepo)
		list, err := s.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockProducts, list)

		mockProductsRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)

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
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

		mockProductsRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct, nil).Once()

		service := NewService(mockProductsRepo)

		product, err := service.GetById(context.Background(), mockProduct.Id)

		assert.NoError(t, err)
		assert.NotEmpty(t, product)

		assert.Equal(t, &mockProduct, product)

		mockProductsRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

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
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

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
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

		mockProductsRepo.On(
			"GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).
			Return(&mockProduct, nil).On(
			"Update",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("failed to retrieve product")).Once()

		service := NewService(mockProductsRepo)
		product, err := service.Update(
			context.Background(), &mockProduct,
		)
		assert.Error(t, err)
		assert.Empty(t, product)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete in case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

		mockProductsRepo.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		service := NewService(mockProductsRepo)

		err := service.Delete(
			context.Background(), mockProduct.Id,
		)
		assert.NoError(t, err)
		mockProductsRepo.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProduct := utils.CreateRandomProduct()

		mockProductsRepo.On("Delete",
			mock.Anything, mock.AnythingOfType("int64"),
		).Return(errors.New("product's ID not founded")).Once()

		service := NewService(mockProductsRepo)

		err := service.Delete(context.Background(), mockProduct.Id)

		assert.Error(t, err)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestCreateProductRecords(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProductRecords := utils.CreateRandomProductRecords()
		mockProductRecordsId := utils.RandomInt64()

		mockProductsRepo.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(mockProductRecordsId, nil).Once()

		s := NewService(mockProductsRepo)

		newRecordId, err := s.CreateProductRecords(context.Background(), &mockProductRecords)

		assert.NoError(t, err)
		assert.Equal(t, mockProductRecordsId, newRecordId)

		mockProductsRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProductRecords := utils.CreateRandomProductRecords()

		mockProductsRepo.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(int64(0), errors.New("failed to create product records")).Once()

		s := NewService(mockProductsRepo)

		_, err := s.CreateProductRecords(context.Background(), &mockProductRecords)

		assert.Error(t, err)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestGetProductRecordsById(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProductRecords := utils.CreateRandomProductRecords()
		mockProductRecordsId := utils.RandomInt64()

		mockProductsRepo.On("GetProductRecordsById", mock.Anything, mock.AnythingOfType("int64")).
			Return(&mockProductRecords, nil).Once()

		service := NewService(mockProductsRepo)

		productRecords, err := service.GetProductRecordsById(context.Background(), mockProductRecordsId)

		assert.NoError(t, err)
		assert.NotEmpty(t, productRecords)

		assert.Equal(t, &mockProductRecords, productRecords)

		mockProductsRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockProductRecordsId := utils.RandomInt64()

		mockProductsRepo.On("GetProductRecordsById", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve product records")).Once()

		service := NewService(mockProductsRepo)

		productRecords, err := service.GetProductRecordsById(context.Background(), mockProductRecordsId)

		assert.Error(t, err)
		assert.Empty(t, productRecords)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestGetQtyOfRecordsById(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockQtyOfRecords := utils.CreateRandomQtyOfRecords()
		mockQtyOfRecordsId := utils.RandomInt64()

		mockProductsRepo.On("GetQtyOfRecordsById", mock.Anything, mock.AnythingOfType("int64")).
			Return(&mockQtyOfRecords, nil).Once()

		service := NewService(mockProductsRepo)

		productRecords, err := service.GetQtyOfRecordsById(context.Background(), mockQtyOfRecordsId)

		assert.NoError(t, err)
		assert.NotEmpty(t, productRecords)

		assert.Equal(t, &mockQtyOfRecords, productRecords)

		mockProductsRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockProductsRepo := mocks.NewRepository(t)
		mockQtyOfRecordsId := utils.RandomInt64()

		mockProductsRepo.On("GetQtyOfRecordsById", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve qty of product records")).Once()

		service := NewService(mockProductsRepo)

		productRecords, err := service.GetQtyOfRecordsById(context.Background(), mockQtyOfRecordsId)

		assert.Error(t, err)
		assert.Empty(t, productRecords)

		mockProductsRepo.AssertExpectations(t)
	})
}

func TestGetQtyOfAllRecords(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockReportsRepo := mocks.NewRepository(t)
		mockReports := utils.CreateRandomListQtyOfRecords()

		mockReportsRepo.On("GetQtyOfAllRecords", mock.Anything).
			Return(&mockReports, nil).Once()

		s := NewService(mockReportsRepo)
		list, err := s.GetQtyOfAllRecords(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockReports, list)

		mockReportsRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockReportsRepo := mocks.NewRepository(t)

		mockReportsRepo.On("GetQtyOfAllRecords", mock.Anything).
			Return(nil, errors.New("failed to retrieve reports")).
			Once()

		s := NewService(mockReportsRepo)
		_, err := s.GetQtyOfAllRecords(context.Background())

		assert.NotNil(t, err)

		mockReportsRepo.AssertExpectations(t)
	})
}
