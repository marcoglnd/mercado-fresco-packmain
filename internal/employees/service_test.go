package employees_test

import (
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomEmployee() (employee Employee) {
	employee = Employee{
		ID:           1,
		CardNumberId: utils.RandomCategory(),
		FirstName:    utils.RandomCategory(),
		LastName:     utils.RandomCategory(),
		WarehouseId:  utils.RandomCode(),
	}
	return
}

func createRandomListEmployee() (listOfEmployee []Employee) {

	for i := 1; i <= 5; i++ {
		employee := createRandomEmployee()
		employee.ID = i
		listOfEmployee = append(listOfEmployee, employee)
	}
	return
}

func TestGetAll(t *testing.T) {
	mock := new(mocks.Repository)

	employeesArg := createRandomListEmployee()

	t.Run("GetAll in case of success", func(t *testing.T) {
		mock.On("GetAll").Return(employeesArg, nil)

		service := NewService(mock)

		list, err := service.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, list)

		for i := 0; i < len(employeesArg); i++ {
			assert.Equal(t, employeesArg[i].ID, list[i].ID)
			assert.Equal(t, employeesArg[i].CardNumberId, list[i].CardNumberId)
			assert.Equal(t, employeesArg[i].FirstName, list[i].FirstName)
			assert.Equal(t, employeesArg[i].LastName, list[i].LastName)
			assert.Equal(t, employeesArg[i].WarehouseId, list[i].WarehouseId)
		}
		mock.AssertExpectations(t)
	})

	t.Run("GetAll in case of error", func(t *testing.T) {
		mock.On("GetAll").Return(nil, errors.New("failed to retrieve employees"))

		service := NewService(mock)

		list, err := service.GetAll()

		assert.Error(t, err)
		assert.Empty(t, list)

		mock.AssertExpectations(t)
	})

}

func TestGetById(t *testing.T) {
	mock := new(mocks.Repository)

	employeeArg := createRandomEmployee()

	t.Run("GetById in case of success", func(t *testing.T) {
		mock.On("GetById", employeeArg.ID).Return(employeeArg, nil)

		service := NewService(mock)

		employee, err := service.GetById(employeeArg.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, employee)

		assert.Equal(t, employeeArg.ID, employee.ID)
		assert.Equal(t, employeeArg.CardNumberId, employee.CardNumberId)
		assert.Equal(t, employeeArg.FirstName, employee.FirstName)
		assert.Equal(t, employeeArg.LastName, employee.LastName)
		assert.Equal(t, employeeArg.WarehouseId, employee.WarehouseId)

		mock.AssertExpectations(t)

	})

	t.Run("GetById in case of error", func(t *testing.T) {
		mock.On("GetById", 42).Return(Employee{}, errors.New("failed to retrieve employee"))

		service := NewService(mock)

		employee, err := service.GetById(42)

		assert.Error(t, err)
		assert.Empty(t, employee)

		mock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Verify employee`s ID increases when a new employee is created", func(t *testing.T) {

		employeesArg := createRandomListEmployee()

		for _, employee := range employeesArg {
			mock.On("Create", employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId).Return(employee, nil)
		}

		service := NewService(mock)

		var list []Employee

		for _, employee := range employeesArg {
			newEmployee, err := service.Create(employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)

			assert.NoError(t, err)
			assert.NotEmpty(t, newEmployee)

			assert.Equal(t, employee.ID, newEmployee.ID)
			assert.Equal(t, employee.CardNumberId, newEmployee.CardNumberId)
			assert.Equal(t, employee.FirstName, newEmployee.FirstName)
			assert.Equal(t, employee.LastName, newEmployee.LastName)
			assert.Equal(t, employee.WarehouseId, newEmployee.WarehouseId)

			list = append(list, newEmployee)

		}
		assert.True(t, list[0].ID == list[1].ID-1)

		mock.AssertExpectations(t)
	})

	t.Run("Verify when CardNumberId`s employee already exists thrown an error", func(t *testing.T) {
		employee1 := createRandomEmployee()
		employee2 := createRandomEmployee()

		employee2.CardNumberId = employee1.CardNumberId

		expectedError := errors.New("CardNumberId already used")

		mock.On("Create", employee1.CardNumberId, employee1.FirstName, employee1.LastName, employee1.WarehouseId).Return(employee1, nil)
		mock.On("Create", employee2.CardNumberId, employee2.FirstName, employee2.LastName, employee2.WarehouseId).Return(Employee{}, expectedError)

		service := NewService(mock)

		newEmployee1, err := service.Create(employee1.CardNumberId, employee1.FirstName, employee1.LastName, employee1.WarehouseId)
		assert.NoError(t, err)
		assert.NotEmpty(t, newEmployee1)

		assert.Equal(t, employee1, newEmployee1)

		newEmployee2, err := service.Create(employee2.CardNumberId, employee2.FirstName, employee2.LastName, employee2.WarehouseId)
		assert.Error(t, err)
		assert.Empty(t, newEmployee2)

		assert.NotEqual(t, employee2, newEmployee2)

		mock.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Update data in case of success", func(t *testing.T) {
		employee1 := createRandomEmployee()
		employee2 := createRandomEmployee()

		employee2.ID = employee1.ID

		mock.On("Create", employee1.CardNumberId, employee1.FirstName, employee1.LastName,
			employee1.WarehouseId).Return(employee1, nil)

		mock.On("Update", employee1.ID, employee2.CardNumberId, employee2.FirstName, employee2.LastName,
			employee2.WarehouseId).Return(employee2, nil)

		service := NewService(mock)

		newEmployee1, err := service.Create(employee1.CardNumberId, employee1.FirstName, employee1.LastName, employee1.WarehouseId)
		assert.NoError(t, err)
		assert.NotEmpty(t, newEmployee1)

		assert.Equal(t, employee1, newEmployee1)

		newEmployee2, err := service.Update(employee1.ID, employee2.CardNumberId, employee2.FirstName, employee2.LastName, employee2.WarehouseId)
		assert.NoError(t, err)
		assert.NotEmpty(t, newEmployee2)

		assert.Equal(t, employee1.ID, newEmployee2.ID)
		assert.NotEqual(t, employee1.CardNumberId, newEmployee2.CardNumberId)
		assert.NotEqual(t, employee1.FirstName, newEmployee2.FirstName)
		assert.NotEqual(t, employee1.LastName, newEmployee2.LastName)
		assert.NotEqual(t, employee1.WarehouseId, newEmployee2.WarehouseId)

		mock.AssertExpectations(t)

	})

	t.Run("Update throw an error in case of an nanoexistent ID", func(t *testing.T) {
		employee := createRandomEmployee()

		mock.On("Update", employee.ID, employee.CardNumberId, employee.FirstName, employee.LastName,
			employee.WarehouseId).Return(Employee{}, errors.New("failed to retrieve employee"))

		service := NewService(mock)

		employee, err := service.Update(employee.ID, employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)

		assert.Error(t, err)
		assert.Empty(t, employee)

		mock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mock := new(mocks.Repository)

	employeeArg := createRandomEmployee()

	t.Run("Delete in case of success", func(t *testing.T) {
		mock.On("Create", employeeArg.CardNumberId, employeeArg.FirstName, employeeArg.LastName,
			employeeArg.WarehouseId).Return(employeeArg, nil)

		mock.On("GetAll").Return([]Employee{employeeArg}, nil)

		mock.On("Delete", employeeArg.ID).Return(nil)

		mock.On("GetAll").Return([]Employee{}, nil)

		service := NewService(mock)

		newEmployee, err := service.Create(employeeArg.CardNumberId, employeeArg.FirstName, employeeArg.LastName, employeeArg.WarehouseId)
		assert.NoError(t, err)

		list1, err := service.GetAll()
		assert.NoError(t, err)

		err = service.Delete(newEmployee.ID)
		assert.NoError(t, err)

		list2, err := service.GetAll()
		assert.NoError(t, err)

		assert.NotEmpty(t, list1)
		assert.NotEqual(t, list1, list2)
		assert.Empty(t, list2)

		mock.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mock.On("Delete", 42).Return(errors.New("employee's ID not founded"))

		service := NewService(mock)

		err := service.Delete(42)

		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}
