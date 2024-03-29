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
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewBuyer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockBuyer := utils.CreateRandomBuyer()
		buyerServiceMock := mocks.NewBuyerService(t)

		buyerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockBuyer, nil).Once()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.POST("/api/v1/buyers", buyerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		buyerServiceMock := mocks.NewBuyerService(t)
		mockBuyer := utils.CreateRandomBuyer()
		mockBuyerBad := &domain.Buyer{}

		buyerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(mockBuyerBad, errors.New("error: unprocessable entity")).Maybe()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.POST("/api/v1/buyers", buyerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("fail with empty values", func(t *testing.T) {
		buyerServiceMock := mocks.NewBuyerService(t)
		mockBuyerBad := &domain.Buyer{}

		buyerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(mockBuyerBad, errors.New("error: unprocessable entity")).Maybe()

		payload, err := json.Marshal(mockBuyerBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.POST("/api/v1/buyers", buyerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		mockBuyer := utils.CreateRandomBuyer()
		buyerServiceMock := mocks.NewBuyerService(t)

		buyerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, domain.ErrDuplicatedID).Maybe()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.POST("/api/v1/buyers", buyerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockBuyer := utils.CreateRandomListBuyers()

	buyerServiceMock := mocks.NewBuyerService(t)

	t.Run("success", func(t *testing.T) {

		buyerServiceMock.On("GetAll",
			mock.Anything,
		).Return(&mockBuyer, nil).Once()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers", buyerController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of internal server error", func(t *testing.T) {
		mockBuyerBad := &[]domain.Buyer{}

		buyerServiceMock.On("GetAll",
			mock.Anything,
		).Return(mockBuyerBad, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockBuyerBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers", buyerController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockBuyer := utils.CreateRandomBuyer()

	buyerServiceMock := mocks.NewBuyerService(t)

	t.Run("success", func(t *testing.T) {

		buyerServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockBuyer, nil).Once()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/%v", mockBuyer.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers/:id", buyerController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid buyer id", func(t *testing.T) {
		mockBuyerBad := &domain.Buyer{}

		buyerServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockBuyerBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockBuyerBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/%v", "a")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers/:id", buyerController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting buyer", func(t *testing.T) {
		buyerServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/%v", utils.RandomInt(0, 999))
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers/:id", buyerController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockBuyer := utils.CreateRandomBuyer()

	buyerServiceMock := mocks.NewBuyerService(t)

	t.Run("success", func(t *testing.T) {

		buyerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockBuyer, nil).Once()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/%v", mockBuyer.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.PATCH("/api/v1/buyers/:id", buyerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid id", func(t *testing.T) {
		buyerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockBuyer, errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/buyers/%v", "a")
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.PATCH("/api/v1/buyers/:id", buyerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of bad request", func(t *testing.T) {
		buyerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockBuyer, errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/buyers/%v", mockBuyer.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.PATCH("/api/v1/buyers/:id", buyerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting buyer", func(t *testing.T) {
		buyerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("id not found error")).Maybe()

		payload, err := json.Marshal(mockBuyer)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/%v", utils.RandomInt(0, 999))
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.PATCH("/api/v1/buyers/:id", buyerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockBuyer := utils.CreateRandomBuyer()

	buyerServiceMock := mocks.NewBuyerService(t)

	t.Run("success", func(t *testing.T) {
		buyerServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		PATH := fmt.Sprintf("/api/v1/buyers/%v", mockBuyer.ID)
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.DELETE("/api/v1/buyers/:id", buyerController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid buyer id", func(t *testing.T) {
		buyerServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/buyers/%v", "a")
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.DELETE("/api/v1/buyers/:id", buyerController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting buyer", func(t *testing.T) {
		buyerServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(errors.New("expected conflict error")).Maybe()

		PATH := fmt.Sprintf("/api/v1/buyers/%v", utils.RandomInt(0, 999))
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.DELETE("/api/v1/buyers/:id", buyerController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})
}

func TestReportPurchaseOrders(t *testing.T) {
	mockReport := utils.CreateRandomReportPurchaseOrder()

	buyerServiceMock := mocks.NewBuyerService(t)

	t.Run("success", func(t *testing.T) {

		buyerServiceMock.On("ReportPurchaseOrders",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockReport, nil).Once()

		payload, err := json.Marshal(mockReport)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/reportPurchaseOrders?id=%v", mockReport.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers/reportPurchaseOrders", buyerController.ReportPurchaseOrders())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of empty id", func(t *testing.T) {
		mockListOfReportPurchaseOrders := utils.CreateRandomListReportPurchaseOrder()

		buyerServiceMock.On("ReportPurchaseOrders",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(&mockListOfReportPurchaseOrders, errors.New("bad request")).Maybe().On("ReportAllPurchaseOrders",
			mock.Anything,
		).Return(&mockListOfReportPurchaseOrders, nil).Once()

		payload, err := json.Marshal(mockListOfReportPurchaseOrders)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/buyers/reportPurchaseOrders?id=%v", "")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers/reportPurchaseOrders", buyerController.ReportPurchaseOrders())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})

	t.Run("In case of internal server error", func(t *testing.T) {
		mockListOfReportPurchaseOrders := utils.CreateRandomListReportPurchaseOrder()
		buyerServiceMock := mocks.NewBuyerService(t)

		buyerServiceMock.On("ReportPurchaseOrders",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockListOfReportPurchaseOrders, nil).Maybe().On("ReportAllPurchaseOrders",
			mock.Anything,
		).Return(nil, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockListOfReportPurchaseOrders)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		buyerController := BuyerController{buyer: buyerServiceMock}

		engine.GET("/api/v1/buyers/reportPurchaseOrders", buyerController.ReportPurchaseOrders())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		buyerServiceMock.AssertExpectations(t)
	})
}
