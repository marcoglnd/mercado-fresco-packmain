package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := buyers.NewRepository()
	service := buyers.NewService(repo)

	buyerController := controllers.NewBuyer(service)

	router := gin.Default()

	pr := router.Group("/buyers")
	{
		pr.GET("/", buyerController.GetAll())
		pr.GET("/:id", buyerController.GetById())
		pr.POST("/", buyerController.Create())
		pr.PATCH("/:id", buyerController.Update())
		pr.DELETE("/:id", buyerController.Delete())
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

func Test_CreateBuyer_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateBuyer_fail(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"first_name": "Jhon", "last_name": "Doe"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_CreateBuyer_conflict(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	second_req, second_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Maria", "last_name": "Silva"
	}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)

	assert.Equal(t, http.StatusConflict, second_rr.Code)
}

func Test_GetBuyers_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/buyers/", "")

	defer req.Body.Close()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []buyers.Buyer
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func Test_GetBuyerById_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/buyers/1", "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusOK, get_rr.Code)

	var objRes buyers.Buyer

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	assert.True(t, objRes.CardNumberID == "402323")
	assert.True(t, objRes.FirstName == "Jhon")
	assert.True(t, objRes.LastName == "Doe")
}

func Test_GetBuyerById_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/buyers/10", "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusNotFound, get_rr.Code)
}

func Test_UpdateBuyer_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, "/buyers/1", `{
		"card_number_id": "400000", "first_name": "Maria", "last_name": "Silva"
	}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusOK, patch_rr.Code)

	var objRes buyers.Buyer

	err := json.Unmarshal(patch_rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	assert.True(t, objRes.CardNumberID == "400000")
	assert.True(t, objRes.FirstName == "Maria")
	assert.True(t, objRes.LastName == "Silva")
}

func Test_UpdateBuyer_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, "/buyers/10", `{
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

	post_req, post_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/buyers/", "")

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	objRes := struct {
		Code int
		Data []buyers.Buyer
	}{}

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	buyersLen := len(objRes.Data)

	assert.Nil(t, err)
	assert.True(t, buyersLen > 0)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, "/buyers/1", "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	secondGet_req, secondGet_rr := createRequestTest(http.MethodGet, "/buyers/", "")

	r.ServeHTTP(secondGet_rr, secondGet_req)

	secondObjRes := struct {
		Code int
		Data []buyers.Buyer
	}{}

	json.Unmarshal(secondGet_rr.Body.Bytes(), &secondObjRes)

	secondBuyersLen := len(secondObjRes.Data)

	assert.Equal(t, http.StatusNoContent, delete_rr.Code)
	assert.True(t, buyersLen-1 == secondBuyersLen)
}

func Test_DeleteBuyer_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	r.ServeHTTP(post_rr, post_req)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, "/buyers/10", "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_rr.Code)
}
