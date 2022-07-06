package controller

import (
	"bytes"
	"encoding/json"
	"errors"
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

// func TestGetAll(t *testing.T) {

// 	t.Run("ok", func(t *testing.T) {
// 		r := createServer()

// 		req, rr := createRequestTest(http.MethodGet, getPathUrl("/sellers/"), "")

// 		defer req.Body.Close()

// 		r.ServeHTTP(rr, req)

// 		assert.Equal(t, http.StatusOK, rr.Code)

// 		objRes := struct {
// 			Code int
// 			Data []domain.Seller
// 		}{}

// 		err := json.Unmarshal(rr.Body.Bytes(), &objRes)

// 		assert.Nil(t, err)
// 		assert.True(t, len(objRes.Data) >= 0)
// 	})
// }

// func TestGetByID(t *testing.T) {

// 	t.Run("non existent", func(t *testing.T) {
// 		r := createServer()

// 		post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sellers/"), `{
// 			"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
// 		}`)

// 		get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/sellers/10"), "")

// 		defer post_req.Body.Close()
// 		defer get_req.Body.Close()

// 		r.ServeHTTP(post_rr, post_req)
// 		r.ServeHTTP(get_rr, get_req)

// 		assert.Equal(t, http.StatusNotFound, get_rr.Code)
// 	})

// 	t.Run("existent", func(t *testing.T) {
// 		seller := &domain.Seller{
// 			ID:           1,
// 			Cid:          402323,
// 			Company_name: "Jhon",
// 			Address:      "Doe",
// 			Telephone:    "1234",
// 		}
// 		mockService := new(mocks.Service)
// 		mockService.On("GetById", mock.AnythingOfType("int")).Return(*seller, nil)

// 		rr := httptest.NewRecorder()
// 		ctx, engine := gin.CreateTestContext(rr)
// 		ns := controller.NewSeller(mockService)

// 		engine.GET("/api/v1/sellers/:id", ns.GetById())
// 		request, err := http.NewRequest(http.MethodGet, "/api/v1/sellers/1", nil)
// 		assert.NoError(t, err)
// 		ctx.Request = request
// 		engine.ServeHTTP(rr, ctx.Request)

// 		assert.Equal(t, http.StatusOK, rr.Code)

// 		var objRes domain.Seller

// 		err = json.Unmarshal(rr.Body.Bytes(), &objRes)

// 		assert.Nil(t, err)
// 		assert.True(t, objRes.ID == 1)
// 		assert.True(t, objRes.Cid == 402323)
// 		assert.True(t, objRes.Company_name == "Jhon")
// 		assert.True(t, objRes.Address == "Doe")
// 		assert.True(t, objRes.Telephone == "1234")
// 	})
// }

// func TestUpdate(t *testing.T) {

// 	t.Run("ok", func(t *testing.T) {
// 		r := createServer()

// 		post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sellers/"), `{
// 			"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
// 		}`)

// 		patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/sellers/1"), `{
// 			"cid": 400000, "company_name": "Maria", "address": "Receba", "telephone": "4321"
// 		}`)

// 		defer post_req.Body.Close()
// 		defer patch_req.Body.Close()

// 		r.ServeHTTP(post_rr, post_req)
// 		r.ServeHTTP(patch_rr, patch_req)

// 		assert.Equal(t, http.StatusOK, patch_rr.Code)

// 		var objRes domain.Seller

// 		err := json.Unmarshal(patch_rr.Body.Bytes(), &objRes)

// 		assert.Nil(t, err)
// 		assert.True(t, objRes.ID == 1)
// 		assert.True(t, objRes.Cid == 400000)
// 		assert.True(t, objRes.Company_name == "Maria")
// 		assert.True(t, objRes.Address == "Receba")
// 		assert.True(t, objRes.Telephone == "4321")
// 	})

// 	t.Run("non existent", func(t *testing.T) {
// 		r := createServer()

// 		post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sellers/"), `{
// 			"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
// 		}`)

// 		patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/sellers/10"), `{
// 			"cid": 400000, "company_name": "Maria", "address": "Receba", "telephone": "4321"
// 		}`)

// 		defer post_req.Body.Close()
// 		defer patch_req.Body.Close()

// 		r.ServeHTTP(post_rr, post_req)
// 		r.ServeHTTP(patch_rr, patch_req)

// 		assert.Equal(t, http.StatusNotFound, patch_rr.Code)
// 	})
// }

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
