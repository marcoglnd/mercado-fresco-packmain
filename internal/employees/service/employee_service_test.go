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
		).Return(&domain.Employee{}, errors.New("failed to create employee")).Once()

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

func TestUpdate(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On("GetById", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockEmployee, nil).On("Update", mock.Anything,
			mock.Anything).Return(&mockEmployee, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		employee, err := service.Update(context.Background(), &mockEmployee)

		assert.NoError(t, err)
		assert.NotEmpty(t, employee)
		assert.Equal(t, &mockEmployee, employee)

		mockEmployeeRepository.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On(
			"GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockEmployee, nil).On(
			"Update",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("failed to retrieve employee")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		employee, err := service.Update(context.Background(), &mockEmployee)

		assert.Error(t, err)
		assert.Empty(t, employee)

		mockEmployeeRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On("Delete", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		err := service.Delete(context.Background(), mockEmployee.ID)

		assert.NoError(t, err)

		mockEmployeeRepository.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockEmployee := utils.CreateRandomEmployee()

		mockEmployeeRepository.On("Delete", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(errors.New("employee's ID not founded")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		err := service.Delete(context.Background(), mockEmployee.ID)

		assert.Error(t, err)

		mockEmployeeRepository.AssertExpectations(t)
	})
}

func TestReportAllInboundOrders(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockInboundOrders := utils.CreateRamdomListReportInboundOrders()

		mockEmployeeRepository.On("ReportAllInboundOrders", mock.Anything).Return(&mockInboundOrders, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		newInboundOrders, err := service.ReportAllInboundOrders(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, &mockInboundOrders, newInboundOrders)

		mockEmployeeRepository.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)

		mockEmployeeRepository.On("ReportAllInboundOrders", mock.Anything).Return(nil, errors.New("failed to retrieve inbound orders")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		_, err := service.ReportAllInboundOrders(context.Background())

		assert.NotNil(t, err)

		mockEmployeeRepository.AssertExpectations(t)

	})
}

func TestReportInboundOrders(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockInboundOrder := utils.CreateRandomReportInboundOrder()

		mockEmployeeRepository.On("ReportInboundOrders", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockInboundOrder, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		inboundOrder, err := service.ReportInboundOrders(context.Background(), mockInboundOrder.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, inboundOrder)
		assert.Equal(t, &mockInboundOrder, inboundOrder)

		mockEmployeeRepository.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockEmployeeRepository := mocks.NewEmployeeRepository(t)
		mockInboundOrder := utils.CreateRandomReportInboundOrder()

		mockEmployeeRepository.On("ReportInboundOrders", mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("failed to retrieve inbound order")).Once()

		service := NewEmployeeService(mockEmployeeRepository)
		inboundOrder, err := service.ReportInboundOrders(context.Background(), mockInboundOrder.ID)

		assert.Error(t, err)
		assert.Empty(t, inboundOrder)

		mockEmployeeRepository.AssertExpectations(t)
	})
}
