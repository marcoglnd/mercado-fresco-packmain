package controllers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductOK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestCreateProductUnprocessable(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	second_req, second_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	third_req, third_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)
	r.ServeHTTP(third_rr, third_req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, second_rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, third_rr.Code)
}

func TestCreateProductConflict(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	second_req, second_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Queijo",
		"expiration_rate": 2,
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)

	assert.Equal(t, http.StatusConflict, second_rr.Code)
}

func TestGetAllOK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, getPathUrl("/products/"), "")

	defer req.Body.Close()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []products.Product
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) >= 0)
}

func TestGetProductByIdOK(t *testing.T) {
	seller := &products.Product{
		Id: 1,
		Description:                    "Yogurt",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "PROD01",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	mockService := new(mocks.Service)
	mockService.On("GetById", mock.AnythingOfType("int")).Return(*seller, nil)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	ns := controllers.NewProduct(mockService)

	engine.GET("/api/v1/products/:id", ns.GetById())
	request, err := http.NewRequest(http.MethodGet, "/api/v1/products/1", nil)
	assert.NoError(t, err)
	ctx.Request = request
	engine.ServeHTTP(rr, ctx.Request)

	assert.Equal(t, http.StatusOK, rr.Code)

	var objRes products.Product

	err = json.Unmarshal(rr.Body.Bytes(), &objRes)
	log.Print(objRes)
	assert.Nil(t, err)
	assert.True(t, objRes.Id == 1)
	assert.True(t, objRes.Description == "Yogurt")
	assert.True(t, objRes.ExpirationRate == 1)
	assert.True(t, objRes.FreezingRate == 2)
	assert.True(t, objRes.Height == 6.4)
	assert.True(t, objRes.Length == 4.5)
	assert.True(t, objRes.NetWeight == 3.4)
	assert.True(t, objRes.ProductCode == "PROD01")
	assert.True(t, objRes.RecommendedFreezingTemperature == 1.3)
	assert.True(t, objRes.Width == 1.2)
	assert.True(t, objRes.ProductTypeId == 2)
	assert.True(t, objRes.SellerId == 2)
}

func TestGetProductByIdNotFound(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/products/10"), "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusNotFound, get_rr.Code)
}

func TestGetProductByIdBadRequest(t *testing.T) {
	r := createServer()

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/products/abc"), "")

	defer get_req.Body.Close()

	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusBadRequest, get_rr.Code)
}

func TestUpdateProductOK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/products/1"), `{
		"description": "Queijo",
		"expiration_rate": 2,
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD02",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusOK, patch_rr.Code)

	var objRes products.Product

	err := json.Unmarshal(patch_rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.Id == 1)
	assert.True(t, objRes.Description == "Queijo")
	assert.True(t, objRes.ExpirationRate == 2)
	assert.True(t, objRes.FreezingRate == 3)
	assert.True(t, objRes.Height == 8.6)
	assert.True(t, objRes.Length == 2.4)
	assert.True(t, objRes.NetWeight == 5.7)
	assert.True(t, objRes.ProductCode == "PROD02")
	assert.True(t, objRes.RecommendedFreezingTemperature == 4.5)
	assert.True(t, objRes.Width == 2.5)
	assert.True(t, objRes.ProductTypeId == 54)
	assert.True(t, objRes.SellerId == 1)
}

func TestUpdateProductNotFound(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/products/10"), `{
		"description": "Queijo",
		"expiration_rate": 2,
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD02",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusNotFound, patch_rr.Code)
}

func TestUpdateProductBadRequest(t *testing.T) {
	r := createServer()

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/products/abc"), `{
		"description": "Queijo",
		"expiration_rate": 2,
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD02",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	defer patch_req.Body.Close()

	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusBadRequest, patch_rr.Code)
}

func TestUpdateProductUnprocessable(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/products/1"), `{
		"expiration_rate": 2,
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD02",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	secondPatch_req, secondPatch_rr := createRequestTest(http.MethodPatch, getPathUrl("/products/1"), `{
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD02",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	thirdPatch_req, thirdPatch_rr := createRequestTest(http.MethodPatch, getPathUrl("/products/1"), `{
		"description": "Queijo",
		"expiration_rate": 2,
		"freezing_rate": 3,
		"product_code": "PROD02",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()
	defer secondPatch_req.Body.Close()
	defer thirdPatch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)
	r.ServeHTTP(secondPatch_rr, secondPatch_req)
	r.ServeHTTP(thirdPatch_rr, thirdPatch_req)

	assert.Equal(t, http.StatusUnprocessableEntity, patch_rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, secondPatch_rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, thirdPatch_rr.Code)
}

func TestDeleteProductOK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/products/"), "")

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	objRes := struct {
		Code int
		Data []products.Product
	}{}

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	productsLen := len(objRes.Data)

	assert.Nil(t, err)
	assert.True(t, productsLen > 0)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/products/1"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	secondGet_req, secondGet_rr := createRequestTest(http.MethodGet, getPathUrl("/products/"), "")

	r.ServeHTTP(secondGet_rr, secondGet_req)

	secondObjRes := struct {
		Code int
		Data []products.Product
	}{}

	json.Unmarshal(secondGet_rr.Body.Bytes(), &secondObjRes)

	secondProductsLen := len(secondObjRes.Data)

	assert.Equal(t, http.StatusNoContent, delete_rr.Code)
	assert.True(t, productsLen-1 == secondProductsLen)
}

func TestDeleteProductsFail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/products/"), `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	r.ServeHTTP(post_rr, post_req)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/products/10"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_rr.Code)
}

func TestDeleteProductsBadRequest(t *testing.T) {
	r := createServer()

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/products/abc"), "")

	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusBadRequest, delete_rr.Code)
}
