package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := sellers.NewRepository()
	service := sellers.NewService(repo)

	sellerController := controllers.NewSeller(service)

	router := gin.Default()

	pr := router.Group("/sellers")
	{
		pr.GET("/", sellerController.GetAll())
		pr.GET("/:id", sellerController.GetById())
		pr.POST("/", sellerController.Create())
		pr.PATCH("/:id", sellerController.Update())
		pr.DELETE("/:id", sellerController.Delete())
	}

	return router
}

func createRequestTest(
	method string,
	url string,
	body string,
) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func Test_CreateSeller_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 123, "company_name": "John", "address": "Rua Meli", "telephone": "1234"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateSeller_bad_request(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": "123", "company_name": "Jhon", "address": "Doe", "telephone": "123456"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_CreateSeller_fail(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"company_name": "Jhon", "address": "Doe"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_CreateSeller_conflict(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
	}`)

	second_req, second_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Maria", "address": "Silva", "telephone": "4321"
	}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)

	assert.Equal(t, http.StatusConflict, second_rr.Code)
}

func Test_GetSellers_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/sellers/", "")

	defer req.Body.Close()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []sellers.Seller
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func Test_GetBuyerById_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/sellers/1", "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusOK, get_rr.Code)

	var objRes sellers.Seller

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	// assert.True(t, objRes.CardNumberID == "402323")
	// assert.True(t, objRes.FirstName == "Jhon")
	// assert.True(t, objRes.LastName == "Doe")
}

func Test_GetBuyerById_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/sellers/10", "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusNotFound, get_rr.Code)
}

func Test_UpdateBuyer_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, "/sellers/1", `{
		"card_number_id": "400000", "first_name": "Maria", "last_name": "Silva"
	}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusOK, patch_rr.Code)

	var objRes sellers.Seller

	err := json.Unmarshal(patch_rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	// assert.True(t, objRes.CardNumberID == "400000")
	// assert.True(t, objRes.FirstName == "Maria")
	// assert.True(t, objRes.LastName == "Silva")
}

func Test_UpdateBuyer_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, "/sellers/10", `{
		"card_number_id": "400000", "first_name": "Maria", "last_name": "Silva"
	}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusNotFound, patch_rr.Code)
}

func Test_DeleteBuyer_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/sellers/", "")

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	objRes := struct {
		Code int
		Data []sellers.Seller
	}{}

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	sellersLen := len(objRes.Data)

	assert.Nil(t, err)
	assert.True(t, sellersLen > 0)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, "/sellers/1", "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	secondGet_req, secondGet_rr := createRequestTest(http.MethodGet, "/sellers/", "")

	r.ServeHTTP(secondGet_rr, secondGet_req)

	secondObjRes := struct {
		Code int
		Data []sellers.Seller
	}{}

	json.Unmarshal(secondGet_rr.Body.Bytes(), &secondObjRes)

	secondsellersLen := len(secondObjRes.Data)

	assert.Equal(t, http.StatusNoContent, delete_rr.Code)
	assert.True(t, sellersLen-1 == secondsellersLen)
}

func Test_DeleteBuyer_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	r.ServeHTTP(post_rr, post_req)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, "/sellers/10", "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_rr.Code)
}
