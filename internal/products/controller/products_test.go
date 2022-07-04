package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewProduct(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {

		productsServiceMock.On("CreateNewProduct",
			mock.Anything,
			mock.Anything,
		).Return(&mockProduct, nil).Once()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/products", productController.CreateNewProduct())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		mockProductBad := &domain.Product{}

		productsServiceMock.On("CreateNewProduct",
			mock.Anything,
			mock.Anything,
		).Return(mockProductBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/products", productController.CreateNewProduct())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		productsServiceMock.On("CreateNewProduct",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/products", productController.CreateNewProduct())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockProduct := utils.CreateRandomListProduct()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {

		productsServiceMock.On("GetAll",
			mock.Anything,
		).Return(&mockProduct, nil).Once()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products", productController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		mockProductBad := &[]domain.Product{}

		productsServiceMock.On("GetAll",
			mock.Anything,
		).Return(mockProductBad, errors.New("not found")).Maybe()

		payload, err := json.Marshal(mockProductBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products", productController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}