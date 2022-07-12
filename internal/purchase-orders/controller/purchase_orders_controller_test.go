package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewPurchaseOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockPurchaseOrder := utils.CreateRandomPurchaseOrder()
		purchaseOrderServiceMock := mocks.NewPurchaseOrderService(t)

		purchaseOrderServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockPurchaseOrder, nil).Once()

		payload, err := json.Marshal(mockPurchaseOrder)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/purchaseOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		purchaseOrderController := PurchaseOrderController{purchaseOrder: purchaseOrderServiceMock}

		engine.POST("/api/v1/purchaseOrders", purchaseOrderController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		purchaseOrderServiceMock.AssertExpectations(t)
	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		purchaseOrderServiceMock := mocks.NewPurchaseOrderService(t)
		mockBadPurchaseOrder := &domain.PurchaseOrder{}

		purchaseOrderServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(mockBadPurchaseOrder, errors.New("error: unprocessable entity")).Maybe()

		payload, err := json.Marshal(mockBadPurchaseOrder)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/purchaseOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		purchaseOrderController := PurchaseOrderController{purchaseOrder: purchaseOrderServiceMock}

		engine.POST("/api/v1/purchaseOrders", purchaseOrderController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		purchaseOrderServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		mockPurchaseOrder := utils.CreateRandomPurchaseOrder()
		purchaseOrderServiceMock := mocks.NewPurchaseOrderService(t)

		purchaseOrderServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, domain.ErrDuplicatedOrderNumber).Maybe()

		payload, err := json.Marshal(mockPurchaseOrder)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/purchaseOrders", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		purchaseOrderController := PurchaseOrderController{purchaseOrder: purchaseOrderServiceMock}

		engine.POST("/api/v1/purchaseOrders", purchaseOrderController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		purchaseOrderServiceMock.AssertExpectations(t)
	})
}
