package service

import (
	"context"
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewBuyer(t *testing.T) {
	mockBuyerRepo := mocks.NewBuyerRepository(t)
	mockBuyer := utils.CreateRandomBuyer()

	t.Run("In case of success", func(t *testing.T) {
		mockBuyerRepo.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockBuyer, nil).Once()
		mockBuyerRepo.On("GetByCardNumberId", mock.Anything, mock.Anything).Return(nil, nil)

		s := NewBuyerService(mockBuyerRepo)

		newProduct, err := s.Create(context.Background(), mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)

		assert.NoError(t, err)
		assert.Equal(t, &mockBuyer, newProduct)

		mockBuyerRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockBuyerRepo.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&Buyer{}, errors.New("failed to create buyer")).Once()

		s := NewBuyerService(mockBuyerRepo)

		_, err := s.Create(context.Background(), mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)

		assert.Error(t, err)

		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockBuyerRepo := mocks.NewBuyerRepository(t)

	mockBuyers := utils.CreateRandomListBuyers()

	t.Run("In case of success", func(t *testing.T) {
		mockBuyerRepo.On("GetAll", mock.Anything).
			Return(&mockBuyers, nil).Once()

		s := NewBuyerService(mockBuyerRepo)
		list, err := s.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockBuyers, list)

		mockBuyerRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockBuyerRepo.On("GetAll", mock.Anything).
			Return(nil, errors.New("failed to retrieve buyers")).
			Once()

		s := NewBuyerService(mockBuyerRepo)
		_, err := s.GetAll(context.Background())

		assert.NotNil(t, err)

		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockBuyerRepo := mocks.NewBuyerRepository(t)

	mockBuyer := utils.CreateRandomBuyer()

	t.Run("In case of success", func(t *testing.T) {
		mockBuyerRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).Return(&mockBuyer, nil).Once()

		service := NewBuyerService(mockBuyerRepo)

		buyer, err := service.GetById(context.Background(), mockBuyer.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, buyer)

		assert.Equal(t, &mockBuyer, buyer)

		mockBuyerRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockBuyerRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve product")).Once()

		service := NewBuyerService(mockBuyerRepo)

		buyer, err := service.GetById(context.Background(), mockBuyer.ID)

		assert.Error(t, err)
		assert.Empty(t, buyer)

		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockBuyerRepo := mocks.NewBuyerRepository(t)

	mockBuyer := utils.CreateRandomBuyer()

	t.Run("In case of success", func(t *testing.T) {
		mockBuyerRepo.On(
			"Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockBuyer, nil).Once()

		service := NewBuyerService(mockBuyerRepo)
		buyer, err := service.Update(
			context.Background(), mockBuyer.ID, mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, buyer)

		assert.Equal(t, &mockBuyer, buyer)

		mockBuyerRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockBuyerRepo.On(
			"Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("failed to update buyer")).Once()

		service := NewBuyerService(mockBuyerRepo)
		product, err := service.Update(
			context.Background(), mockBuyer.ID, mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName,
		)
		assert.Error(t, err)
		assert.Empty(t, product)

		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockBuyerRepo := mocks.NewBuyerRepository(t)

	mockBuyer := utils.CreateRandomBuyer()

	t.Run("Delete in case of success", func(t *testing.T) {
		mockBuyerRepo.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		service := NewBuyerService(mockBuyerRepo)

		err := service.Delete(
			context.Background(), mockBuyer.ID,
		)
		assert.NoError(t, err)
		mockBuyerRepo.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mockBuyerRepo.On("Delete",
			mock.Anything, mock.AnythingOfType("int64"),
		).Return(errors.New("buyer's ID not founded")).Once()

		service := NewBuyerService(mockBuyerRepo)

		err := service.Delete(context.Background(), mockBuyer.ID)

		assert.Error(t, err)

		mockBuyerRepo.AssertExpectations(t)
	})
}
