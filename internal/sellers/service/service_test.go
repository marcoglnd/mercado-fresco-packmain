package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {

	mockSeller := utils.CreateRandomListSeller()
	sellerRepositoryMock := mocks.NewSellerRepository(t)

	t.Run("ok", func(t *testing.T) {
		sellerRepositoryMock.On("GetAll", mock.Anything).
			Return(&mockSeller, nil).Once()

		service := NewService(sellerRepositoryMock)
		list, err := service.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockSeller, list)

		sellerRepositoryMock.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		sellerRepositoryMock.On("GetAll", mock.Anything).
			Return(nil, errors.New("failed to retrieve sellers")).
			Once()

		service := NewService(sellerRepositoryMock)
		_, err := service.GetAll(context.Background())

		assert.NotNil(t, err)

		sellerRepositoryMock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {

	mockSeller := utils.CreateRandomSeller()
	sellerRepositoryMock := mocks.NewSellerRepository(t)

	t.Run("existent", func(t *testing.T) {
		sellerRepositoryMock.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockSeller, nil).Once()

		service := NewService(sellerRepositoryMock)

		seller, err := service.GetByID(context.Background(), mockSeller.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, seller)

		assert.Equal(t, &mockSeller, seller)

		sellerRepositoryMock.AssertExpectations(t)
	})

	t.Run("non existent", func(t *testing.T) {
		sellerRepositoryMock.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve seller")).Once()

		service := NewService(sellerRepositoryMock)

		seller, err := service.GetByID(context.Background(), mockSeller.ID)

		assert.Error(t, err)
		assert.Empty(t, seller)

		sellerRepositoryMock.AssertExpectations(t)
	})
}

// func TestCreate(t *testing.T) {

// 	sellerRepositoryMock := mocks.NewSellerRepository(t)
// 	mockSeller := utils.CreateRandomSeller()

// 	t.Run("ok", func(t *testing.T) {
// 		sellerRepositoryMock.On("Create",
// 			mock.Anything,
// 			mock.Anything,
// 			mock.Anything,
// 			mock.Anything,
// 		).Return(&mockSeller, nil).Once()
// 		sellerRepositoryMock.On("GetByID", mock.Anything, mock.Anything).Return(nil, nil)

// 		service := NewService(sellerRepositoryMock)
// 		seller, err := service.Create(context.Background(), &mockSeller)

// 		assert.NoError(t, err)
// 		assert.Equal(t, &mockSeller, seller)

// 		sellerRepositoryMock.AssertExpectations(t)
// 	})

// 	t.Run("fail", func(t *testing.T) {
// 		sellerRepositoryMock.On("Create",
// 			mock.Anything,
// 			mock.Anything,
// 			mock.Anything,
// 			mock.Anything,
// 		).Return(&domain.Seller{}, errors.New("failed to create seller")).Once()

// 		service := NewService(sellerRepositoryMock)
// 		_, err := service.Create(context.Background(), &mockSeller)

// 		assert.Error(t, err)

// 		sellerRepositoryMock.AssertExpectations(t)
// 	})
// }
func TestUpdate(t *testing.T) {

	mockSeller := utils.CreateRandomSeller()
	sellerRepositoryMock := mocks.NewSellerRepository(t)

	t.Run("ok", func(t *testing.T) {
		sellerRepositoryMock.On(
			"Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSeller, nil).Once()

		service := NewService(sellerRepositoryMock)
		seller, err := service.Update(context.Background(), &mockSeller)
		assert.NoError(t, err)
		assert.NotEmpty(t, seller)

		assert.Equal(t, &mockSeller, seller)

		sellerRepositoryMock.AssertExpectations(t)
	})

	t.Run("non existent", func(t *testing.T) {
		sellerRepositoryMock.On(
			"Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("failed to update seller")).Once()

		service := NewService(sellerRepositoryMock)
		seller, err := service.Update(context.Background(), &mockSeller)
		assert.Error(t, err)
		assert.Empty(t, seller)

		sellerRepositoryMock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {

	mockSeller := utils.CreateRandomSeller()
	sellerRepositoryMock := mocks.NewSellerRepository(t)

	t.Run("ok", func(t *testing.T) {
		sellerRepositoryMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		service := NewService(sellerRepositoryMock)
		err := service.Delete(
			context.Background(), mockSeller.ID,
		)
		assert.NoError(t, err)
		sellerRepositoryMock.AssertExpectations(t)
	})

	t.Run("non existent", func(t *testing.T) {
		sellerRepositoryMock.On("Delete",
			mock.Anything, mock.AnythingOfType("int64"),
		).Return(errors.New("seller's ID not founded")).Once()

		service := NewService(sellerRepositoryMock)
		err := service.Delete(context.Background(), mockSeller.ID)

		assert.Error(t, err)

		sellerRepositoryMock.AssertExpectations(t)
	})
}
