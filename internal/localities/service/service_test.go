package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLocality(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)
		mockLocality := utils.CreateRandomLocality()
		mockLocalityId := utils.RandomInt64()

		localitiesRepositoryMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(mockLocalityId, nil).Once()

		service := NewService(localitiesRepositoryMock)
		localityId, err := service.CreateLocality(context.Background(), &mockLocality)

		assert.NoError(t, err)
		assert.Equal(t, mockLocalityId, localityId)

		localitiesRepositoryMock.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)
		mockLocality := utils.CreateRandomLocality()

		localitiesRepositoryMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(int64(0), errors.New("failed to create locality")).Once()

		service := NewService(localitiesRepositoryMock)
		_, err := service.CreateLocality(context.Background(), &mockLocality)

		assert.Error(t, err)

		localitiesRepositoryMock.AssertExpectations(t)
	})
}

func TestGetLocalityByID(t *testing.T) {
	t.Run("In case of error", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)
		mockLocalityId := utils.RandomInt64()

		localitiesRepositoryMock.On("GetLocalityByID", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve locality")).Once()

		service := NewService(localitiesRepositoryMock)

		locality, err := service.GetLocalityByID(context.Background(), mockLocalityId)

		assert.Error(t, err)
		assert.Empty(t, locality)

		localitiesRepositoryMock.AssertExpectations(t)
	})
}

func TestGetAllQtyOfSellers(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)
		mockQtyOfSellers := utils.CreateRandomListQtyOfSellers()

		localitiesRepositoryMock.On("GetAllQtyOfSellers", mock.Anything).
			Return(&mockQtyOfSellers, nil).Once()

		service := NewService(localitiesRepositoryMock)
		list, err := service.GetAllQtyOfSellers(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockQtyOfSellers, list)

		localitiesRepositoryMock.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)

		localitiesRepositoryMock.On("GetAllQtyOfSellers", mock.Anything).
			Return(nil, errors.New("failed to retrieve sellers")).
			Once()

		service := NewService(localitiesRepositoryMock)
		_, err := service.GetAllQtyOfSellers(context.Background())

		assert.NotNil(t, err)

		localitiesRepositoryMock.AssertExpectations(t)
	})
}

func TestGetQtyOfSellersByLocalityId(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)
		mockQtyOfSellers := utils.CreateRandomQtyOfSellers()
		mockQtyOfSellersId := utils.RandomInt64()

		localitiesRepositoryMock.On("GetQtyOfSellersByLocalityId", mock.Anything, mock.AnythingOfType("int64")).
			Return(&mockQtyOfSellers, nil).Once()

		service := NewService(localitiesRepositoryMock)

		productRecords, err := service.GetQtyOfSellersByLocalityId(context.Background(), mockQtyOfSellersId)

		assert.NoError(t, err)
		assert.NotEmpty(t, productRecords)

		assert.Equal(t, &mockQtyOfSellers, productRecords)

		localitiesRepositoryMock.AssertExpectations(t)

	})

	t.Run("fail", func(t *testing.T) {
		localitiesRepositoryMock := mocks.NewLocalityRepository(t)
		mockQtyOfRecordsId := utils.RandomInt64()

		localitiesRepositoryMock.On("GetQtyOfSellersByLocalityId", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve qty of product records")).Once()

		service := NewService(localitiesRepositoryMock)

		productRecords, err := service.GetQtyOfSellersByLocalityId(context.Background(), mockQtyOfRecordsId)

		assert.Error(t, err)
		assert.Empty(t, productRecords)

		localitiesRepositoryMock.AssertExpectations(t)
	})
}
