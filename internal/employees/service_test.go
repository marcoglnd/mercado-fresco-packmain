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
