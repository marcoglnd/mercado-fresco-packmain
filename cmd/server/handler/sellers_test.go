package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerForSeller() *gin.Engine {
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

func Test_CreateSeller_OK(t *testing.T) {
	r := createServerForSeller()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 123, "company_name": "John", "address": "Rua Meli", "telephone": "1234"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateSeller_bad_request(t *testing.T) {
	r := createServerForSeller()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": "123", "company_name": "Jhon", "address": "Doe", "telephone": "123456"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_CreateSeller_fail(t *testing.T) {
	r := createServerForSeller()

	req, rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"company_name": "Jhon", "address": "Doe"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_CreateSeller_conflict(t *testing.T) {
	r := createServerForSeller()

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

func Test_GetAllSellers_OK(t *testing.T) {
	r := createServerForSeller()

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
	assert.True(t, len(objRes.Data) >= 0)
}

func Test_GetSellerById_existent(t *testing.T) {
	r := createServerForSeller()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
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
	assert.True(t, objRes.Cid == 402323)
	assert.True(t, objRes.Company_name == "Jhon")
	assert.True(t, objRes.Address == "Doe")
	assert.True(t, objRes.Telephone == "1234")
}

func Test_GetSellerById_non_existent(t *testing.T) {
	r := createServerForSeller()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/sellers/10", "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusNotFound, get_rr.Code)
}

func Test_UpdateSeller_OK(t *testing.T) {
	r := createServerForSeller()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, "/sellers/1", `{
		"cid": 400000, "company_name": "Maria", "address": "Receba", "telephone": "4321"
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
	assert.True(t, objRes.Cid == 400000)
	assert.True(t, objRes.Company_name == "Maria")
	assert.True(t, objRes.Address == "Receba")
	assert.True(t, objRes.Telephone == "4321")
}

func Test_UpdateSeller_non_existent(t *testing.T) {
	r := createServerForSeller()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, "/sellers/10", `{
		"cid": 400000, "company_name": "Maria", "address": "Receba", "telephone": "4321"
	}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusNotFound, patch_rr.Code)
}

func Test_DeleteSeller_OK(t *testing.T) {
	r := createServerForSeller()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, "/sellers/", "")

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	objRes := struct {
		Code int
		Data []sellers.Seller
	}{}

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	buyersLen := len(objRes.Data)

	assert.Nil(t, err)
	assert.True(t, buyersLen > 0)

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

	secondBuyersLen := len(secondObjRes.Data)

	assert.Equal(t, http.StatusNoContent, delete_rr.Code)
	assert.True(t, buyersLen-1 == secondBuyersLen)
}

func Test_DeleteSeller_non_existent(t *testing.T) {
	r := createServerForSeller()

	post_req, post_rr := createRequestTest(http.MethodPost, "/sellers/", `{
		"cid": 402323, "company_name": "Jhon", "address": "Doe", "telephone": "1234"
	}`)

	r.ServeHTTP(post_rr, post_req)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, "/sellers/10", "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_rr.Code)
}
