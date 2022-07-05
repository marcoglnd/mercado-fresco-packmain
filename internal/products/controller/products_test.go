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

	t.Run("In case of internal server error", func(t *testing.T) {
		mockProductBad := &[]domain.Product{}

		productsServiceMock.On("GetAll",
			mock.Anything,
		).Return(mockProductBad, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockProductBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products", productController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {

		productsServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockProduct, nil).Once()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/%v", mockProduct.Id)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/:id", productController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid product id", func(t *testing.T) {
		mockProductBad := &domain.Product{}

		productsServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockProductBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/%v", "a")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/:id", productController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting product", func(t *testing.T) {
		productsServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/:id", productController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {

		productsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockProduct, nil).Once()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/%v", mockProduct.Id)
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.PATCH("/api/v1/products/:id", productController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of unprocessable entity", func(t *testing.T) {
		productsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockProduct, errors.New("unprocessable entity")).Maybe()

		PATH := fmt.Sprintf("/api/v1/products/%v", mockProduct.Id)
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.PATCH("/api/v1/products/:id", productController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid product id", func(t *testing.T) {
		mockProductBad := &domain.Product{}

		productsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(mockProductBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/%v", "a")
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.PATCH("/api/v1/products/:id", productController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting product", func(t *testing.T) {
		productsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected not found error")).Maybe()

		payload, err := json.Marshal(mockProduct)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.PATCH("/api/v1/products/:id", productController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {
		productsServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		PATH := fmt.Sprintf("/api/v1/products/%v", mockProduct.Id)
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.DELETE("/api/v1/products/:id", productController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid product id", func(t *testing.T) {
		productsServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/products/%v", "a")
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.DELETE("/api/v1/products/:id", productController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting product", func(t *testing.T) {
		productsServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(errors.New("expected conflict error")).Maybe()

		PATH := fmt.Sprintf("/api/v1/products/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.DELETE("/api/v1/products/:id", productController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestCreateProductRecords(t *testing.T) {
	mockProductRecords := utils.CreateRandomProductRecords()
	mockProductRecordsId := utils.RandomInt64()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {

		productsServiceMock.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(mockProductRecordsId, nil).Once().
			On("GetProductRecordsById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(&mockProductRecords, nil).Once()

		payload, err := json.Marshal(mockProductRecords)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productRecords", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productRecords", productController.CreateProductRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		mockProductRecordsBad := &domain.ProductRecords{}

		productsServiceMock.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(int64(0), errors.New("bad request")).Maybe().
			On("GetProductRecordsById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductRecordsBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductRecordsBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productRecords", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productRecords", productController.CreateProductRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status unprocessable entity", func(t *testing.T) {
		mockProductRecordsBad := &domain.ProductRecords{}
		mockProductRecordsId := utils.RandomInt64()

		productsServiceMock.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(mockProductRecordsId, errors.New("bad request")).Maybe().
			On("GetProductRecordsById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductRecordsBad, nil).Maybe()

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productRecords", nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productRecords", productController.CreateProductRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestGetQtyOfRecordsById(t *testing.T) {
	mockQtyOfRecords := utils.CreateRandomQtyOfRecords()
	mockQtyOfRecordsId := utils.RandomInt64()

	productsServiceMock := mocks.NewService(t)

	t.Run("success", func(t *testing.T) {

		productsServiceMock.On("GetQtyOfRecordsById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockQtyOfRecords, nil).Once()

		payload, err := json.Marshal(mockQtyOfRecords)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/reportRecords?id=%v", mockQtyOfRecordsId)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecordsById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid qty of records id", func(t *testing.T) {
		mockQtyOfRecordsBad := &domain.QtyOfRecords{}

		productsServiceMock.On("GetQtyOfRecordsById",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockQtyOfRecordsBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockQtyOfRecordsBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/reportRecords?id=%v", "a")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecordsById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting qty of records", func(t *testing.T) {
		productsServiceMock.On("GetQtyOfRecordsById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockQtyOfRecords)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/products/reportRecords?id=%v", mockQtyOfRecordsId)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecordsById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}