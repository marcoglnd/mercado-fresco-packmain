package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses"
	mock_warehouses "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodPost, getPathUrl("/warehouses/"),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 22",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestCreateFail(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodPost, getPathUrl("/warehouses/"),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 22",
			"telephone": "1130304040",
			"minimum_capacity": 10
		}
		`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}

func TestCreateConflict(t *testing.T) {
	routes := createServer()

	reqFake, resFake := createRequestTest(http.MethodPost, getPathUrl("/warehouses/"),
		`
	{
		"warehouse_code": "BRO",
		"address": "Rua Sao Paulo 22",
		"telephone": "1130304040",
		"minimum_capacity": 10,
		"minimum_temperature": 20
	}
	`,
	)

	defer reqFake.Body.Close()
	routes.ServeHTTP(resFake, reqFake)

	req, res := createRequestTest(http.MethodPost, getPathUrl("/warehouses/"),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 22",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusConflict, res.Code)
}

func TestFindAll(t *testing.T) {
	routes := createServer()
	reqFake, resFake := createRequestTest(
		http.MethodPost,
		getPathUrl("/warehouses/"),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 22",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer reqFake.Body.Close()
	routes.ServeHTTP(resFake, reqFake)
	req, res := createRequestTest(http.MethodGet, getPathUrl("/warehouses/"), "")

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	objRes := struct {
		Data []warehouses.Warehouse
	}{}

	err := json.Unmarshal(res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindByIdNonExistent(t *testing.T) {
	routes := createServer()
	inexistentId := 10
	req, res := createRequestTest(
		http.MethodGet,
		getPathUrl(fmt.Sprintf("/warehouses/%d", inexistentId)),
		"",
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)
}

func TestFindByIdExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	fakeWarehouse := &warehouses.Warehouse{
		ID:                 1,
		WarehouseCode:      "UBS",
		Address:            "Rua Sao Joao",
		Telephone:          "3120104030",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	serviceMock := mock_warehouses.NewMockService(ctrl)
	serviceMock.EXPECT().FindById(fakeWarehouse.ID).Return(fakeWarehouse, nil)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	wc := controllers.NewWarehouse(serviceMock)

	engine.GET("/:id", wc.GetById())
	request, err := http.NewRequest(http.MethodGet, "/1", nil)
	ctx.Request = request
	engine.ServeHTTP(rr, ctx.Request)
	assert.NoError(t, err)

	var objRes warehouses.Warehouse
	err = json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, *fakeWarehouse, objRes)
}

func TestUpdateOk(t *testing.T) {
	routes := createServer()
	reqFakePost, resFakePost := createRequestTest(
		http.MethodPost,
		getPathUrl("/warehouses/"),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 22",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer reqFakePost.Body.Close()
	routes.ServeHTTP(resFakePost, reqFakePost)
	reqFake, resFake := createRequestTest(http.MethodGet, getPathUrl("/warehouses/1"), "")
	defer reqFake.Body.Close()
	routes.ServeHTTP(resFake, reqFake)

	assert.Equal(t, http.StatusOK, resFake.Code)
	type createdWarehouse struct {
		ID int `json:"id"`
		warehouses.Warehouse
	}

	var oldObjRes createdWarehouse
	errFake := json.Unmarshal(resFake.Body.Bytes(), &oldObjRes)
	assert.Nil(t, errFake)

	req, res := createRequestTest(
		http.MethodPatch,
		getPathUrl(fmt.Sprintf("/warehouses/%d", oldObjRes.ID)),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 21",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)
	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	var newObjRes createdWarehouse
	err := json.Unmarshal(res.Body.Bytes(), &newObjRes)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotEqual(t, oldObjRes.Address, newObjRes.Address)
}

func TestUpdateInvalidSchema(t *testing.T) {
	routes := createServer()

	req, res := createRequestTest(
		http.MethodPatch,
		getPathUrl(fmt.Sprintf("/warehouses/%d", 1)),
		`
		{
			"warehouse": "BRO",
			"address": "Rua Sao Paulo 21",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}

func TestUpdateNonExistent(t *testing.T) {
	routes := createServer()

	reqValidateCheck, resValidateCheck := createRequestTest(
		http.MethodGet,
		getPathUrl("/warehouses/notint"),
		"",
	)

	defer reqValidateCheck.Body.Close()
	routes.ServeHTTP(resValidateCheck, reqValidateCheck)
	assert.Equal(t, http.StatusBadRequest, resValidateCheck.Code)

	inexistentId := 10
	req, res := createRequestTest(
		http.MethodPatch,
		getPathUrl(fmt.Sprintf("/warehouses/%d", inexistentId)),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 21",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)
}

func TestDeleteNonExistent(t *testing.T) {
	routes := createServer()
	inexistentId := 10
	req, res := createRequestTest(
		http.MethodDelete,
		getPathUrl(fmt.Sprintf("/warehouses/%d", inexistentId)), "")

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)
}

func TestDeleteOk(t *testing.T) {
	routes := createServer()
	reqFakePost, resFakePost := createRequestTest(http.MethodPost, getPathUrl("/warehouses/"),
		`
		{
			"warehouse_code": "BRO",
			"address": "Rua Sao Paulo 22",
			"telephone": "1130304040",
			"minimum_capacity": 10,
			"minimum_temperature": 20
		}
		`,
	)

	defer reqFakePost.Body.Close()
	routes.ServeHTTP(resFakePost, reqFakePost)
	existentId := 1
	req, res := createRequestTest(
		http.MethodDelete,
		getPathUrl(fmt.Sprintf("/warehouses/%d", existentId)), "")

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}

func TestDeleteInvalidParam(t *testing.T) {
	routes := createServer()
	reqValidateCheck, resValidateCheck := createRequestTest(
		http.MethodGet,
		getPathUrl("/warehouses/notint"),
		"",
	)

	defer reqValidateCheck.Body.Close()
	routes.ServeHTTP(resValidateCheck, reqValidateCheck)
	assert.Equal(t, http.StatusBadRequest, resValidateCheck.Code)
}
