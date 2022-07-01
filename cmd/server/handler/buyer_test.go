package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateBuyer_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateBuyer_conflict(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	second_req, second_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Maria", "last_name": "Silva"
	}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)

	assert.Equal(t, http.StatusConflict, second_rr.Code)
}

func Test_CreateBuyer_unprocessable(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"first_name": "Jhon", "last_name": "Doe"
	}`)

	second_req, second_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "401010", "last_name": "Silva"
	}`)

	third_req, third_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "401010", "first_name": "Jhon"
	}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)
	r.ServeHTTP(third_rr, third_req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, second_rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, third_rr.Code)
}

func Test_GetBuyers_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, getPathUrl("/buyers/"), "")

	defer req.Body.Close()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []buyers.Buyer
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) >= 0)
}

func Test_GetBuyerById_OK(t *testing.T) {
	buyer := &buyers.Buyer{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
	}
	mockService := new(mocks.Service)
	mockService.On("GetById", mock.AnythingOfType("int")).Return(*buyer, nil)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	ns := controllers.NewBuyer(mockService)

	engine.GET("/api/v1/buyers/:id", ns.GetById())
	request, err := http.NewRequest(http.MethodGet, "/api/v1/buyers/1", nil)
	assert.NoError(t, err)
	ctx.Request = request
	engine.ServeHTTP(rr, ctx.Request)

	assert.Equal(t, http.StatusOK, rr.Code)

	var objRes buyers.Buyer

	err = json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	assert.True(t, objRes.CardNumberID == "402323")
	assert.True(t, objRes.FirstName == "Jhon")
	assert.True(t, objRes.LastName == "Doe")
}

func Test_GetBuyerById_notFound(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/buyers/10"), "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusNotFound, get_rr.Code)
}

func Test_GetBuyerById_badRequest(t *testing.T) {
	r := createServer()

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/buyers/abc"), "")

	defer get_req.Body.Close()

	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusBadRequest, get_rr.Code)
}

func Test_UpdateBuyer_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/buyers/1"), `{
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

func Test_UpdateBuyer_notFound(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/buyers/10"), `{
		"card_number_id": "400000", "first_name": "Maria", "last_name": "Silva"
	}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusNotFound, patch_rr.Code)
}

func Test_UpdateBuyer_badRequest(t *testing.T) {
	r := createServer()

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/buyers/abc"), `{
		"card_number_id": "400000", "first_name": "Maria", "last_name": "Silva"
	}`)

	defer patch_req.Body.Close()

	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusBadRequest, patch_rr.Code)
}

func Test_UpdateBuyer_unprocessable(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/buyers/1"), `{
		"first_name": "Maria", "last_name": "Silva"
	}`)

	secondPatch_req, secondPatch_rr := createRequestTest(http.MethodPatch, getPathUrl("/buyers/1"), `{
		"card_number_id": "400000", "last_name": "Silva"
	}`)

	thirdPatch_req, thirdPatch_rr := createRequestTest(http.MethodPatch, getPathUrl("/buyers/1"), `{
		"card_number_id": "400000", "first_name": "Maria"
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

func Test_DeleteBuyer_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/buyers/"), "")

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

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/buyers/1"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	secondGet_req, secondGet_rr := createRequestTest(http.MethodGet, getPathUrl("/buyers/"), "")

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

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/buyers/"), `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	r.ServeHTTP(post_rr, post_req)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/buyers/10"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_rr.Code)
}
