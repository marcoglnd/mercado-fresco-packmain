package sections_test

import (
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/sections"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomSection() (section Section) {
	section = Section{
		ID:                 1,
		SectionNumber:      utils.RandomCode(),
		CurrentTemperature: utils.RandomCode(),
		MinimumTemperature: utils.RandomCode(),
		CurrentCapacity:    utils.RandomCode(),
		MinimumCapacity:    utils.RandomCode(),
		MaximumCapacity:    utils.RandomCode(),
		WarehouseId:        utils.RandomCode(),
		ProductTypeId:      utils.RandomCode(),
	}
	return
}

func createRandomListSection() (listOfSections []Section) {

	for i := 1; i <= 5; i++ {
		section := createRandomSection()
		section.ID = i
		listOfSections = append(listOfSections, section)
	}
	return
}

func TestGetAll(t *testing.T) {
	mock := new(mocks.Repository)

	sectionsArg := createRandomListSection()

	t.Run("GetAll in case of success", func(t *testing.T) {
		mock.On("GetAll").Return(sectionsArg, nil).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, list)

		for i := 0; i < len(sectionsArg); i++ {
			assert.Equal(t, sectionsArg[i].ID, list[i].ID)
			assert.Equal(t, sectionsArg[i].SectionNumber, list[i].SectionNumber)
			assert.Equal(t, sectionsArg[i].CurrentTemperature, list[i].CurrentTemperature)
			assert.Equal(t, sectionsArg[i].MinimumTemperature, list[i].MinimumTemperature)
			assert.Equal(t, sectionsArg[i].CurrentCapacity, list[i].CurrentCapacity)
			assert.Equal(t, sectionsArg[i].MinimumCapacity, list[i].MinimumCapacity)
			assert.Equal(t, sectionsArg[i].MaximumCapacity, list[i].MaximumCapacity)
			assert.Equal(t, sectionsArg[i].WarehouseId, list[i].WarehouseId)
			assert.Equal(t, sectionsArg[i].ProductTypeId, list[i].ProductTypeId)
		}

		mock.AssertExpectations(t)
	})

	t.Run("GetAll in case of error", func(t *testing.T) {
		mock.On("GetAll").Return(nil, errors.New("failed to retrieve Sections")).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.Error(t, err)
		assert.Empty(t, list)

		mock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mock := new(mocks.Repository)

	sectionArg := createRandomSection()

	t.Run("GetById in case of success", func(t *testing.T) {
		mock.On("GetById", sectionArg.ID).Return(sectionArg, nil).Once()

		service := NewService(mock)

		section, err := service.GetById(sectionArg.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, section)

		assert.Equal(t, sectionArg.ID, section.ID)
		assert.Equal(t, sectionArg.SectionNumber, section.SectionNumber)
		assert.Equal(t, sectionArg.CurrentTemperature, section.CurrentTemperature)
		assert.Equal(t, sectionArg.MinimumTemperature, section.MinimumTemperature)
		assert.Equal(t, sectionArg.CurrentCapacity, section.CurrentCapacity)
		assert.Equal(t, sectionArg.MinimumCapacity, section.MinimumCapacity)
		assert.Equal(t, sectionArg.MaximumCapacity, section.MaximumCapacity)
		assert.Equal(t, sectionArg.WarehouseId, section.WarehouseId)
		assert.Equal(t, sectionArg.ProductTypeId, section.ProductTypeId)

		mock.AssertExpectations(t)

	})

	t.Run("GetById in case of error", func(t *testing.T) {
		mock.On("GetById", 185).Return(Section{}, errors.New("failed to retrieve section")).Once()

		service := NewService(mock)

		section, err := service.GetById(185)

		assert.Error(t, err)
		assert.Empty(t, section)

		mock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Verify section's ID increases when a new section is created", func(t *testing.T) {

		sectionsArg := createRandomListSection()

		for _, section := range sectionsArg {
			mock.On("Create",
				section.SectionNumber,
				section.CurrentTemperature,
				section.MinimumTemperature,
				section.CurrentCapacity,
				section.MinimumCapacity,
				section.MaximumCapacity,
				section.WarehouseId,
				section.ProductTypeId).Return(section, nil).Once()
		}

		service := NewService(mock)

		var list []Section

		for _, sectionArg := range sectionsArg {
			newSection, err := service.Create(
				sectionArg.SectionNumber,
				sectionArg.CurrentTemperature,
				sectionArg.MinimumTemperature,
				sectionArg.CurrentCapacity,
				sectionArg.MinimumCapacity,
				sectionArg.MaximumCapacity,
				sectionArg.WarehouseId,
				sectionArg.ProductTypeId,
			)

			assert.NoError(t, err)
			assert.NotEmpty(t, newSection)

			assert.Equal(t, sectionArg.ID, newSection.ID)
			assert.Equal(t, sectionArg.SectionNumber, newSection.SectionNumber)
			assert.Equal(t, sectionArg.CurrentTemperature, newSection.CurrentTemperature)
			assert.Equal(t, sectionArg.MinimumTemperature, newSection.MinimumTemperature)
			assert.Equal(t, sectionArg.CurrentCapacity, newSection.CurrentCapacity)
			assert.Equal(t, sectionArg.MinimumCapacity, newSection.MinimumCapacity)
			assert.Equal(t, sectionArg.MaximumCapacity, newSection.MaximumCapacity)
			assert.Equal(t, sectionArg.WarehouseId, newSection.WarehouseId)
			assert.Equal(t, sectionArg.ProductTypeId, newSection.ProductTypeId)
			list = append(list, newSection)
		}
		assert.True(t, list[0].ID == list[1].ID-1)

		mock.AssertExpectations(t)
	})

	t.Run("Verify when a SectionNumber`s section already exists thrown an error", func(t *testing.T) {
		section1 := createRandomSection()
		section2 := createRandomSection()

		section2.SectionNumber = section1.SectionNumber

		expectedError := errors.New("SectionNumber already used")

		mock.On("Create",
			section1.SectionNumber,
			section1.CurrentTemperature,
			section1.MinimumTemperature,
			section1.CurrentCapacity,
			section1.MinimumCapacity,
			section1.MaximumCapacity,
			section1.WarehouseId,
			section1.ProductTypeId,
		).Return(section1, nil).Once()
		mock.On("Create",
			section2.SectionNumber,
			section2.CurrentTemperature,
			section2.MinimumTemperature,
			section2.CurrentCapacity,
			section2.MinimumCapacity,
			section2.MaximumCapacity,
			section2.WarehouseId,
			section2.ProductTypeId,
		).Return(Section{}, expectedError).Once()

		s := NewService(mock)
		newSection1, err := s.Create(
			section1.SectionNumber,
			section1.CurrentTemperature,
			section1.MinimumTemperature,
			section1.CurrentCapacity,
			section1.MinimumCapacity,
			section1.MaximumCapacity,
			section1.WarehouseId,
			section1.ProductTypeId,
		)

		assert.NoError(t, err)
		assert.NotEmpty(t, newSection1)

		assert.Equal(t, section1, newSection1)

		newSection2, err := s.Create(
			section2.SectionNumber,
			section2.CurrentTemperature,
			section2.MinimumTemperature,
			section2.CurrentCapacity,
			section2.MinimumCapacity,
			section2.MaximumCapacity,
			section2.WarehouseId,
			section2.ProductTypeId,
		)
		assert.Error(t, expectedError, err)
		assert.Empty(t, newSection2)

		assert.NotEqual(t, section2, newSection2)
		mock.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Update data in case of success", func(t *testing.T) {
		section1 := createRandomSection()
		section2 := createRandomSection()

		section2.ID = section1.ID

		mock.On("Create",
			section1.SectionNumber,
			section1.CurrentTemperature,
			section1.MinimumTemperature,
			section1.CurrentCapacity,
			section1.MinimumCapacity,
			section1.MaximumCapacity,
			section1.WarehouseId,
			section1.ProductTypeId,
		).Return(section1, nil).Once()
		mock.On("Update",
			section1.ID,
			section2.SectionNumber,
			section2.CurrentTemperature,
			section2.MinimumTemperature,
			section2.CurrentCapacity,
			section2.MinimumCapacity,
			section2.MaximumCapacity,
			section2.WarehouseId,
			section2.ProductTypeId,
		).Return(section2, nil).Once()

		s := NewService(mock)
		newSection1, err := s.Create(
			section1.SectionNumber,
			section1.CurrentTemperature,
			section1.MinimumTemperature,
			section1.CurrentCapacity,
			section1.MinimumCapacity,
			section1.MaximumCapacity,
			section1.WarehouseId,
			section1.ProductTypeId,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, newSection1)

		assert.Equal(t, section1, newSection1)

		newSection2, err := s.Update(
			section1.ID,
			section2.SectionNumber,
			section2.CurrentTemperature,
			section2.MinimumTemperature,
			section2.CurrentCapacity,
			section2.MinimumCapacity,
			section2.MaximumCapacity,
			section2.WarehouseId,
			section2.ProductTypeId,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, newSection2)

		assert.Equal(t, section1.ID, newSection2.ID)
		assert.NotEqual(t, section1.SectionNumber, newSection2.SectionNumber)
		assert.NotEqual(t, section1.CurrentTemperature, newSection2.CurrentTemperature)
		assert.NotEqual(t, section1.MinimumTemperature, newSection2.MinimumTemperature)
		assert.NotEqual(t, section1.CurrentCapacity, newSection2.CurrentCapacity)
		assert.NotEqual(t, section1.MinimumCapacity, newSection2.MinimumCapacity)
		assert.NotEqual(t, section1.MaximumCapacity, newSection2.MaximumCapacity)
		assert.NotEqual(t, section1.WarehouseId, newSection2.WarehouseId)
		assert.NotEqual(t, section1.ProductTypeId, newSection2.ProductTypeId)

		mock.AssertExpectations(t)
	})

	t.Run("Update throw an error in case of an nonexistent ID", func(t *testing.T) {
		section := createRandomSection()

		mock.On("Update",
			section.ID,
			section.SectionNumber,
			section.CurrentTemperature,
			section.MinimumTemperature,
			section.CurrentCapacity,
			section.MinimumCapacity,
			section.MaximumCapacity,
			section.WarehouseId,
			section.ProductTypeId,
		).Return(Section{}, errors.New("failed to retrieve section")).Once()

		service := NewService(mock)

		section, err := service.Update(
			section.ID,
			section.SectionNumber,
			section.CurrentTemperature,
			section.MinimumTemperature,
			section.CurrentCapacity,
			section.MinimumCapacity,
			section.MaximumCapacity,
			section.WarehouseId,
			section.ProductTypeId,
		)

		assert.Error(t, err)
		assert.Empty(t, section)

		mock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mock := new(mocks.Repository)

	sectionArg := createRandomSection()

	t.Run("Delete in case of success", func(t *testing.T) {
		mock.On("Create",
			sectionArg.SectionNumber,
			sectionArg.CurrentTemperature,
			sectionArg.MinimumTemperature,
			sectionArg.CurrentCapacity,
			sectionArg.MinimumCapacity,
			sectionArg.MaximumCapacity,
			sectionArg.WarehouseId,
			sectionArg.ProductTypeId,
		).Return(sectionArg, nil).Once()
		mock.On("GetAll").Return([]Section{sectionArg}, nil).Once()
		mock.On("Delete", sectionArg.ID).Return(nil).Once()
		mock.On("GetAll").Return([]Section{}, nil).Once()

		service := NewService(mock)

		newSection, err := service.Create(
			sectionArg.SectionNumber,
			sectionArg.CurrentTemperature,
			sectionArg.MinimumTemperature,
			sectionArg.CurrentCapacity,
			sectionArg.MinimumCapacity,
			sectionArg.MaximumCapacity,
			sectionArg.WarehouseId,
			sectionArg.ProductTypeId,
		)
		assert.NoError(t, err)
		list1, err := service.GetAll()
		assert.NoError(t, err)
		err = service.Delete(newSection.ID)
		assert.NoError(t, err)
		list2, err := service.GetAll()
		assert.NoError(t, err)

		assert.NotEmpty(t, list1)
		assert.NotEqual(t, list1, list2)
		assert.Empty(t, list2)

		mock.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mock.On("Delete", 185).Return(errors.New("section's ID not founded")).Once()

		service := NewService(mock)

		err := service.Delete(185)

		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}

