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
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {

	mockSeller := utils.CreateRandomListSeller()
	sellerServiceMock := mocks.NewSellerService(t)

	t.Run("ok", func(t *testing.T) {
		sellerServiceMock.On("GetAll",
			mock.Anything,
		).Return(&mockSeller, nil).Once()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.GET("/api/v1/sellers", sellerController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockSellerBad := &[]domain.Seller{}

		sellerServiceMock.On("GetAll",
			mock.Anything,
		).Return(mockSellerBad, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockSellerBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.GET("/api/v1/sellers", sellerController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {

	mockSeller := utils.CreateRandomSeller()
	sellerServiceMock := mocks.NewSellerService(t)

	t.Run("existent", func(t *testing.T) {
		sellerServiceMock.On("GetByID",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockSeller, nil).Once()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sellers/%v", mockSeller.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.GET("/api/v1/sellers/:id", sellerController.GetByID())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("non existent", func(t *testing.T) {
		sellerServiceMock.On("GetByID",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sellers/%v", utils.RandomInt(0, 999))
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.GET("/api/v1/sellers/:id", sellerController.GetByID())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		mockSellerBad := &domain.Seller{}

		sellerServiceMock.On("GetByID",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockSellerBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockSellerBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sellers/%v", "a")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.GET("/api/v1/sellers/:id", sellerController.GetByID())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		mockSeller := utils.CreateRandomSeller()
		sellerServiceMock := mocks.NewSellerService(t)

		sellerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSeller, nil).Once()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.POST("/api/v1/sellers", sellerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		sellerServiceMock := mocks.NewSellerService(t)
		mockSellerBad := &domain.Seller{}

		sellerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(mockSellerBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockSellerBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.POST("/api/v1/sellers", sellerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockSeller := utils.CreateRandomSeller()
		sellerServiceMock := mocks.NewSellerService(t)

		sellerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("internal error")).Maybe()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.POST("/api/v1/sellers", sellerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("conflict", func(t *testing.T) {
		mockSeller := utils.CreateRandomSeller()
		sellerServiceMock := mocks.NewSellerService(t)

		sellerServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, domain.ErrDuplicatedCID).Maybe()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.POST("/api/v1/sellers", sellerController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {

	mockSeller := utils.CreateRandomSeller()
	sellerServiceMock := mocks.NewSellerService(t)

	t.Run("ok", func(t *testing.T) {
		sellerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSeller, nil).Once()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sellers/%v", mockSeller.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.PATCH("/api/v1/sellers/:id", sellerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("non existent", func(t *testing.T) {
		sellerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("id not found error")).Maybe()

		payload, err := json.Marshal(mockSeller)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sellers/%v", utils.RandomInt(0, 999))
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.PATCH("/api/v1/sellers/:id", sellerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		sellerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSeller, errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/sellers/%v", mockSeller.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.PATCH("/api/v1/sellers/:id", sellerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})

	t.Run("invalid id", func(t *testing.T) {
		sellerServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockSeller, errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/sellers/%v", "a")
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sellerController := SellerController{service: sellerServiceMock}

		engine.PATCH("/api/v1/sellers/:id", sellerController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sellerServiceMock.AssertExpectations(t)
	})
}

// func TestDelete(t *testing.T) {

// 	t.Run("ok", func(t *testing.T) {
// 		r := createServer()

// 		post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sellers/"), `{
// 			"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
// 		}`)

// 		get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/sellers/"), "")

// 		r.ServeHTTP(post_rr, post_req)
// 		r.ServeHTTP(get_rr, get_req)

// 		objRes := struct {
// 			Code int
// 			Data []domain.Seller
// 		}{}

// 		err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

// 		buyersLen := len(objRes.Data)

// 		assert.Nil(t, err)
// 		assert.True(t, buyersLen > 0)

// 		delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/sellers/1"), "")

// 		defer post_req.Body.Close()
// 		defer delete_req.Body.Close()

// 		r.ServeHTTP(delete_rr, delete_req)

// 		secondGet_req, secondGet_rr := createRequestTest(http.MethodGet, getPathUrl("/sellers/"), "")

// 		r.ServeHTTP(secondGet_rr, secondGet_req)

// 		secondObjRes := struct {
// 			Code int
// 			Data []domain.Seller
// 		}{}

// 		json.Unmarshal(secondGet_rr.Body.Bytes(), &secondObjRes)

// 		secondBuyersLen := len(secondObjRes.Data)

// 		assert.Equal(t, http.StatusNoContent, delete_rr.Code)
// 		assert.True(t, buyersLen-1 == secondBuyersLen)
// 	})

// 	t.Run("non existent", func(t *testing.T) {
// 		r := createServer()

// 		post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sellers/"), `{
// 			"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
// 		}`)

// 		r.ServeHTTP(post_rr, post_req)

// 		delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/sellers/10"), "")

// 		defer post_req.Body.Close()
// 		defer delete_req.Body.Close()

// 		r.ServeHTTP(delete_rr, delete_req)

// 		assert.Equal(t, http.StatusNotFound, delete_rr.Code)
// 	})
// }
