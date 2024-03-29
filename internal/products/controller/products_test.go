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

	t.Run("success", func(t *testing.T) {
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
		productsServiceMock := mocks.NewService(t)
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
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
	t.Run("success", func(t *testing.T) {
		mockProduct := utils.CreateRandomListProduct()
		productsServiceMock := mocks.NewService(t)

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
		productsServiceMock := mocks.NewService(t)
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
	t.Run("success", func(t *testing.T) {
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
		productsServiceMock := mocks.NewService(t)
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
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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

	t.Run("success", func(t *testing.T) {
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
		productsServiceMock := mocks.NewService(t)
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
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
	t.Run("success", func(t *testing.T) {
		mockProduct := utils.CreateRandomProduct()
		productsServiceMock := mocks.NewService(t)

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
		productsServiceMock := mocks.NewService(t)

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
		productsServiceMock := mocks.NewService(t)

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
	t.Run("success", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductRecords := utils.CreateRandomProductRecords()
		mockProductRecordsId := utils.RandomInt64()

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

	t.Run("fail with status internal server error", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductRecords := utils.CreateRandomProductRecords()
		mockProductRecordsId := utils.RandomInt64()
		mockProductRecordsBad := &domain.ProductRecords{}

		productsServiceMock.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(mockProductRecordsId, nil).
			On("GetProductRecordsById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductRecordsBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductRecords)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productRecords", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productRecords", productController.CreateProductRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status unprocessable entity", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
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

	t.Run("fail with status conflict", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductRecordsBad := &domain.ProductRecords{}
		mockProductRecords := utils.CreateRandomProductRecords()

		productsServiceMock.On("CreateProductRecords",
			mock.Anything,
			mock.Anything,
		).Return(int64(0), errors.New("bad request")).
			On("GetProductRecordsById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductRecordsBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductRecords)
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
}

func TestGetQtyOfRecords(t *testing.T) {
	t.Run("success - GetQtyOfAllRecords", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockQtyOfRecords := utils.CreateRandomListQtyOfRecords()

		productsServiceMock.On("GetQtyOfAllRecords",
			mock.Anything,
		).Return(&mockQtyOfRecords, nil).Once()

		payload, err := json.Marshal(mockQtyOfRecords)
		assert.NoError(t, err)
		PATH := "/api/v1/products/reportRecords"
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("internal sever error - GetAllQtyOfSellers", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)

		productsServiceMock.On("GetQtyOfAllRecords",
			mock.Anything,
		).Return(nil, errors.New("couldn`t return a list")).Once()

		PATH := "/api/v1/products/reportRecords"
		req := httptest.NewRequest(http.MethodGet, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("success - GetQtyOfRecordsById", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockQtyOfRecords := utils.CreateRandomQtyOfRecords()

		productsServiceMock.On("GetQtyOfRecordsById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockQtyOfRecords, nil).Once()

		payload, err := json.Marshal(mockQtyOfRecords)
		assert.NoError(t, err)
		PATH := fmt.Sprintf("/api/v1/products/reportRecords?id=%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("not found - GetQtyOfRecordsById", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)

		productsServiceMock.On("GetQtyOfRecordsById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("id not found")).Once()

		PATH := fmt.Sprintf("/api/v1/products/reportRecords?id=%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportRecords", productController.GetQtyOfRecords())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestCreateProductBatches(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductBatches := utils.CreateRandomProductBatches()
		mockProductBatchesId := utils.RandomInt64()

		productsServiceMock.On("CreateProductBatches",
			mock.Anything,
			mock.Anything,
		).Return(mockProductBatchesId, nil).Once().
			On("GetProductBatchesById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(&mockProductBatches, nil).Once()

		payload, err := json.Marshal(mockProductBatches)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productBatches", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productBatches", productController.CreateProductBatches())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status internal server error", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductBatches := utils.CreateRandomProductBatches()
		mockProductBatchesId := utils.RandomInt64()
		mockProductBatchesBad := &domain.ProductBatches{}

		productsServiceMock.On("CreateProductBatches",
			mock.Anything,
			mock.Anything,
		).Return(mockProductBatchesId, nil).
			On("GetProductBatchesById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductBatchesBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductBatches)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productBatches", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productBatches", productController.CreateProductBatches())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status unprocessable entity", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductBatchesBad := &domain.ProductBatches{}
		mockProductBatchesId := utils.RandomInt64()

		productsServiceMock.On("CreateProductBatches",
			mock.Anything,
			mock.Anything,
		).Return(mockProductBatchesId, errors.New("bad request")).Maybe().
			On("GetProductBatchesById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductBatchesBad, nil).Maybe()

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productBatches", nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productBatches", productController.CreateProductBatches())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockProductBatchesBad := &domain.ProductBatches{}
		mockProductBatches := utils.CreateRandomProductBatches()

		productsServiceMock.On("CreateProductBatches",
			mock.Anything,
			mock.Anything,
		).Return(int64(0), errors.New("bad request")).
			On("GetProductBatchesById",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(mockProductBatchesBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockProductBatches)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productBatches", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.POST("/api/v1/productBatches", productController.CreateProductBatches())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}

func TestGetQtdProductsBySectionId(t *testing.T) {
	t.Run("success - GetQtdOfAllProducts", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockQtyOfReports := utils.CreateRandomListQtdOfProducts()

		productsServiceMock.On("GetQtdOfAllProducts",
			mock.Anything,
		).Return(&mockQtyOfReports, nil).Once()

		payload, err := json.Marshal(mockQtyOfReports)
		assert.NoError(t, err)
		PATH := "/api/v1/products/reportProducts"
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportProducts", productController.GetQtdProductsBySectionId())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("internal sever error - GetQtdOfAllProducts", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)

		productsServiceMock.On("GetQtdOfAllProducts",
			mock.Anything,
		).Return(nil, errors.New("couldn`t return a list")).Once()

		PATH := "/api/v1/products/reportProducts"
		req := httptest.NewRequest(http.MethodGet, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportProducts", productController.GetQtdProductsBySectionId())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("success - GetQtdProductsBySectionId", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)
		mockQtdOfReports := utils.CreateRandomQtdOfProducts()

		productsServiceMock.On("GetQtdProductsBySectionId",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockQtdOfReports, nil).Once()

		payload, err := json.Marshal(mockQtdOfReports)
		assert.NoError(t, err)
		PATH := fmt.Sprintf("/api/v1/products/reportProducts?id=%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportProducts", productController.GetQtdProductsBySectionId())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})

	t.Run("not found - GetQtdProductsBySectionId", func(t *testing.T) {
		productsServiceMock := mocks.NewService(t)

		productsServiceMock.On("GetQtdProductsBySectionId",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("id not found")).Once()

		PATH := fmt.Sprintf("/api/v1/products/reportProducts?id=%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		productController := Controller{service: productsServiceMock}

		engine.GET("/api/v1/products/reportProducts", productController.GetQtdProductsBySectionId())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		productsServiceMock.AssertExpectations(t)
	})
}
