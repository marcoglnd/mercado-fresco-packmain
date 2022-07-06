package service_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewBuyer(t *testing.T) {
	mockSectionRepo := mocks.NewRepository(t)
	mockSection := utils.CreateRandomSection()

	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, nil).Once()
		mockSectionRepo.On("GetBySectionNumber", mock.Anything, mock.Anything).Return(nil, nil)

		s := NewService(mockSectionRepo)

		newProduct, err := s.Create(context.Background(), mockSection., mockSection.FirstName, mockSection.LastName)

		assert.NoError(t, err)
		assert.Equal(t, &mockSection, newProduct)

		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&Buyer{}, errors.New("failed to create buyer")).Once()

		s := NewBuyerService(mockSectionRepo)

		_, err := s.Create(context.Background(), mockSection.CardNumberID, mockSection.FirstName, mockSection.LastName)

		assert.Error(t, err)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockSectionRepo := mocks.NewRepository(t)

	mockSections := utils.CreateRandomListBuyers()

	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo.On("GetAll", mock.Anything).
			Return(&mockSections, nil).Once()

		s := NewBuyerService(mockSectionRepo)
		list, err := s.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockSections, list)

		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo.On("GetAll", mock.Anything).
			Return(nil, errors.New("failed to retrieve buyers")).
			Once()

		s := NewBuyerService(mockSectionRepo)
		_, err := s.GetAll(context.Background())

		assert.NotNil(t, err)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockSectionRepo := mocks.NewRepository(t)

	mockSection := utils.CreateRandomSection()

	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).Return(&mockSection, nil).Once()

		service := NewBuyerService(mockSectionRepo)

		buyer, err := service.GetById(context.Background(), mockSection.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, buyer)

		assert.Equal(t, &mockSection, buyer)

		mockSectionRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve product")).Once()

		service := NewBuyerService(mockSectionRepo)

		buyer, err := service.GetById(context.Background(), mockSection.ID)

		assert.Error(t, err)
		assert.Empty(t, buyer)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockSectionRepo := mocks.NewRepository(t)

	mockSection := utils.CreateRandomSection()

	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo.On(
			"Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, nil).Once()

		service := NewBuyerService(mockSectionRepo)
		buyer, err := service.Update(
			context.Background(), mockSection.ID, mockSection.CardNumberID, mockSection.FirstName, mockSection.LastName,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, buyer)

		assert.Equal(t, &mockSection, buyer)

		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo.On(
			"Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("failed to update buyer")).Once()

		service := NewBuyerService(mockSectionRepo)
		product, err := service.Update(
			context.Background(), mockSection.ID, mockSection.CardNumberID, mockSection.FirstName, mockSection.LastName,
		)
		assert.Error(t, err)
		assert.Empty(t, product)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockSectionRepo := mocks.NewRepository(t)

	mockSection := utils.CreateRandomSection()

	t.Run("Delete in case of success", func(t *testing.T) {
		mockSectionRepo.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		service := NewBuyerService(mockSectionRepo)

		err := service.Delete(
			context.Background(), mockSection.ID,
		)
		assert.NoError(t, err)
		mockSectionRepo.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mockSectionRepo.On("Delete",
			mock.Anything, mock.AnythingOfType("int64"),
		).Return(errors.New("buyer's ID not founded")).Once()

		service := NewBuyerService(mockSectionRepo)

		err := service.Delete(context.Background(), mockSection.ID)

		assert.Error(t, err)

		mockSectionRepo.AssertExpectations(t)
	})
}
