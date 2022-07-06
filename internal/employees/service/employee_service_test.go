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
