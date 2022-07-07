package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewEmployee(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&mockEmployee, nil).Once()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.POST("/api/v1/employees", employeeController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		mockEmployeeService.AssertExpectations(t)

	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		mockEmployeeService := mocks.NewEmployeeService(t)
		mockEmployee := &domain.Employee{}

		mockEmployeeService.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(mockEmployee, errors.New("unprocessable entity")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.POST("/api/v1/employees", employeeController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.POST("/api/v1/employees", employeeController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}
