package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewEmployee(t *testing.T) {
	mockEmployeeRepository := mocks.NewEmployeeRepository(t)
	mockEmployee := utils.CreateRandomEmployee()

	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&mockEmployee, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		newEmployee, err := service.Create(context.Background(), &mockEmployee)

		assert.NoError(t, err)
		assert.Equal(t, &mockEmployee, newEmployee)

		mockEmployeeRepository.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&domain.Employee{}, errors.New("failes to create employee")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		_, err := service.Create(context.Background(), &mockEmployee)

		assert.Error(t, err)

		mockEmployeeRepository.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployees := utils.CreateRandomListEmployees()

		mockEmployeeRepository.On("GetAll", mock.Anything).Return(&mockEmployees, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		newEmployees, err := service.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, &mockEmployees, newEmployees)

		mockEmployeeRepository.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)

		mockEmployeeRepository.On("GetAll", mock.Anything).Return(nil, errors.New("failed to retrieve employees")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		_, err := service.GetAll(context.Background())

		assert.NotNil(t, err)

		mockEmployeeRepository.AssertExpectations(t)

	})
}

func TestGetBiId(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On("GetById", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockEmployee, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		employee, err := service.GetById(context.Background(), mockEmployee.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, employee)
		assert.Equal(t, &mockEmployee, employee)

		mockEmployeeRepository.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On("GetById", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("failed to retrieve employee")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		employee, err := service.GetById(context.Background(), mockEmployee.ID)

		assert.Error(t, err)
		assert.Empty(t, employee)

		mockEmployeeRepository.AssertExpectations(t)
	})
}
