package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewInboundOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockInboundOrder := utils.CreateRandomInboundOrder()
		mockInboundOrderService := mocks.NewInboundOrderService(t)

		mockInboundOrderService.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&mockInboundOrder, nil).Once()

		payload, err := json.Marshal(mockInboundOrder)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		inboundOrderController := InboundOrderController{service: mockInboundOrderService}

		engine.POST("/api/v1/inboundOrders", inboundOrderController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		mockInboundOrderService.AssertExpectations(t)
	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		mockInboundOrder := &domain.InboundOrder{}
		mockInboundOrderService := mocks.NewInboundOrderService(t)

		mockInboundOrderService.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(mockInboundOrder, errors.New("unprocessable entity")).Maybe()

		payload, err := json.Marshal(mockInboundOrder)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		inboundOrderController := InboundOrderController{service: mockInboundOrderService}

		engine.POST("/api/v1/inboundOrders", inboundOrderController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		mockInboundOrderService.AssertExpectations(t)

	})

	t.Run("fail with status conflict", func(t *testing.T) {
		mockInboundOrder := utils.CreateRandomInboundOrder()
		mockInboundOrderService := mocks.NewInboundOrderService(t)

		mockInboundOrderService.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockInboundOrder)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		inboundOrderController := InboundOrderController{service: mockInboundOrderService}

		engine.POST("/api/v1/inboundOrders", inboundOrderController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		mockInboundOrderService.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockInboundOrders := utils.CreateRandomListInboundOrders()
		mockInboundOrderService := mocks.NewInboundOrderService(t)

		mockInboundOrderService.On("GetAll",
			mock.Anything,
		).Return(&mockInboundOrders, nil).Once()

		payload, err := json.Marshal(mockInboundOrders)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		inboundOrderController := InboundOrderController{service: mockInboundOrderService}

		engine.GET("/api/v1/inboundOrders", inboundOrderController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockInboundOrderService.AssertExpectations(t)
	})

	t.Run("In case of internal server error", func(t *testing.T) {
		mockInboundOrderService := mocks.NewInboundOrderService(t)
		mockInboundOrders := &[]domain.InboundOrder{}

		mockInboundOrderService.On("GetAll",
			mock.Anything,
		).Return(mockInboundOrders, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockInboundOrders)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		inboundOrderController := InboundOrderController{service: mockInboundOrderService}

		engine.GET("/api/v1/inboundOrders", inboundOrderController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		mockInboundOrderService.AssertExpectations(t)
	})
}
