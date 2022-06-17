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
