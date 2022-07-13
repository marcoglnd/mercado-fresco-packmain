package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewSection(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, nil).Once()

		s := NewService(mockSectionRepo)

		newSection, err := s.Create(context.Background(), &mockSection)

		assert.NoError(t, err)
		assert.Equal(t, &mockSection, newSection)

		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&domain.Section{}, errors.New("failed to create buyer")).Once()

		s := NewService(mockSectionRepo)

		_, err := s.Create(context.Background(), &mockSection)

		assert.Error(t, err)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSections := utils.CreateRandomListSection()

		mockSectionRepo.On("GetAll", mock.Anything).
			Return(&mockSections, nil).Once()

		s := NewService(mockSectionRepo)
		list, err := s.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, &mockSections, list)

		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)

		mockSectionRepo.On("GetAll", mock.Anything).
			Return(nil, errors.New("failed to retrieve sections")).
			Once()

		s := NewService(mockSectionRepo)
		_, err := s.GetAll(context.Background())

		assert.NotNil(t, err)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).Return(&mockSection, nil).Once()

		service := NewService(mockSectionRepo)

		section, err := service.GetById(context.Background(), mockSection.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, section)

		assert.Equal(t, &mockSection, section)

		mockSectionRepo.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil, errors.New("failed to retrieve product")).Once()

		service := NewService(mockSectionRepo)

		section, err := service.GetById(context.Background(), mockSection.ID)

		assert.Error(t, err)
		assert.Empty(t, section)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On(
			"GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).
			Return(&mockSection, nil).On(
			"Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, nil).Once()

		service := NewService(mockSectionRepo)
		section, err := service.Update(
			context.Background(), &mockSection,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, section)

		assert.Equal(t, &mockSection, section)

		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On(
			"GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).
			Return(&mockSection, nil).On(
			"Update",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("failed to update section")).Once()

		service := NewService(mockSectionRepo)
		product, err := service.Update(
			context.Background(), &mockSection,
		)
		assert.Error(t, err)
		assert.Empty(t, product)

		mockSectionRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete in case of success", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		service := NewService(mockSectionRepo)

		err := service.Delete(
			context.Background(), mockSection.ID,
		)
		assert.NoError(t, err)
		mockSectionRepo.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mockSectionRepo := mocks.NewRepository(t)
		mockSection := utils.CreateRandomSection()

		mockSectionRepo.On("Delete",
			mock.Anything, mock.AnythingOfType("int64"),
		).Return(errors.New("buyer's ID not founded")).Once()

		service := NewService(mockSectionRepo)

		err := service.Delete(context.Background(), mockSection.ID)

		assert.Error(t, err)

		mockSectionRepo.AssertExpectations(t)
	})
}
