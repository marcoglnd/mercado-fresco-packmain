package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockEmployee := utils.CreateRandomListEmployees()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("GetAll",
			mock.Anything,
		).Return(&mockEmployee, nil).Once()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.GET("/api/v1/employees", employeeController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of internal server error", func(t *testing.T) {
		mockEmployeeService := mocks.NewEmployeeService(t)
		mockEmployee := &[]domain.Employee{}

		mockEmployeeService.On("GetAll",
			mock.Anything,
		).Return(mockEmployee, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.GET("/api/v1/employees", employeeController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockEmployee, nil).Once()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/%v", mockEmployee.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.GET("/api/v1/employees/:id", employeeController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of invalid employee id", func(t *testing.T) {
		mockEmployeeService := mocks.NewEmployeeService(t)
		mockEmployee := &domain.Employee{}

		mockEmployeeService.On("GetById",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockEmployee, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/%v", "a")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.GET("/api/v1/employees/:id", employeeController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of nonexisting employee", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.GET("/api/v1/employees/:id", employeeController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockEmployee, nil).Once()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/%v", mockEmployee.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.PATCH("/api/v1/employees/:id", employeeController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of bad request", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockEmployee, errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/employees/%v", mockEmployee.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.PATCH("/api/v1/employees/:id", employeeController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of invalid employee id", func(t *testing.T) {
		mockEmployeeService := mocks.NewEmployeeService(t)
		mockEmployee := &domain.Employee{}

		mockEmployeeService.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(mockEmployee, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/%v", "a")
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.PATCH("/api/v1/employees/:id", employeeController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of nonexisting employee", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected not found error")).Maybe()

		payload, err := json.Marshal(mockEmployee)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.PATCH("/api/v1/employees/:id", employeeController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockEmployee := utils.CreateRandomEmployee()
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		PATH := fmt.Sprintf("/api/v1/employees/%v", mockEmployee.ID)
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.DELETE("/api/v1/employees/:id", employeeController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of invalid employee id", func(t *testing.T) {
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Delete",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/employees/%v", "a")
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.DELETE("/api/v1/employees/:id", employeeController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("In case of nonexisting employee", func(t *testing.T) {
		mockEmployeeService := mocks.NewEmployeeService(t)

		mockEmployeeService.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(errors.New("expected conflict error")).Maybe()

		PATH := fmt.Sprintf("/api/v1/employees/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: mockEmployeeService}

		engine.DELETE("/api/v1/employees/:id", employeeController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestReportInboundOrders(t *testing.T) {
	mockInboundOrder := utils.CreateRandomReportInboundOrder()
	employeeServiceMock := mocks.NewEmployeeService(t)

	t.Run("success", func(t *testing.T) {

		employeeServiceMock.On("ReportInboundOrders",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockInboundOrder, nil).Once()

		payload, err := json.Marshal(mockInboundOrder)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/reportInboundOrders?id=%v", mockInboundOrder.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: employeeServiceMock}

		engine.GET("/api/v1/employees/reportInboundOrders", employeeController.ReportInboundOrders())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		employeeServiceMock.AssertExpectations(t)
	})

	t.Run("In case of empty id", func(t *testing.T) {
		mockListReportInboundOrders := utils.CreateRamdomListReportInboundOrders()

		employeeServiceMock.On("ReportInboundOrders",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(&mockListReportInboundOrders, errors.New("bad request")).Maybe().On("ReportAllInboundOrders",
			mock.Anything,
		).Return(&mockListReportInboundOrders, nil).Once()

		payload, err := json.Marshal(mockListReportInboundOrders)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/employees/reportInboundOrders?id=%v", "")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: employeeServiceMock}

		engine.GET("/api/v1/employees/reportInboundOrders", employeeController.ReportInboundOrders())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		employeeServiceMock.AssertExpectations(t)
	})

	t.Run("In case of internal server error", func(t *testing.T) {
		mockListReportInboundOrders := utils.CreateRandomListInboundOrders()
		employeeServiceMock := mocks.NewEmployeeService(t)

		employeeServiceMock.On("ReportInboundOrders",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockListReportInboundOrders, nil).Maybe().On("ReportAllInboundOrders",
			mock.Anything,
		).Return(nil, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockListReportInboundOrders)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/reportInboundOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		employeeController := EmployeeController{service: employeeServiceMock}

		engine.GET("/api/v1/employees/reportInboundOrders", employeeController.ReportInboundOrders())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		employeeServiceMock.AssertExpectations(t)
	})
}
