package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/routes"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
	"github.com/stretchr/testify/assert"
)

func getPathUrl(url string) string {
	PATH := "/api/v1"
	return fmt.Sprintf("%s%s", PATH, url)
}

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routerGroup := router.Group(getPathUrl(""))
	routes.AddRoutes(routerGroup)

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

func TestCreateOk(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"),
		` 
		{
			"card_number_id": "1234",
			"first_name": "Paloma",
			"last_name": "Ribeiro",
			"warehouse_id": 2
		}
		`,
	)
	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestCreateFail(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"),
		` 
		{
			"card_number_id": "1234",
			"first_name": "Paloma",
			"last_name": "Ribeiro",
		}
		`,
	)
	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}

func TestCreateConflict(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"),
		` 
		{
			"card_number_id": "1234",
			"first_name": "Paloma",
			"last_name": "Ribeiro",
			"warehouse_id": 2
		}
		`)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusConflict, res.Code)

}

func TestFindAll(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodGet, getPathUrl("/employees/"), "")

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	objRes := struct {
		Data []employees.Employee
	}{}

	err := json.Unmarshal(res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) >= 0)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestFindByIdNonExistent(t *testing.T) {
	routes := createServer()
	inexistentId := 10
	req, res := createRequestTest(http.MethodGet, getPathUrl(fmt.Sprintf("/employees/%d", inexistentId)), "")

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)

}

func TestFindByIdExistent(t *testing.T) {
	routes := createServer()

	reqValidateCheck, resValidateCheck := createRequestTest(http.MethodGet, getPathUrl("/employees/notint"), "")

	defer reqValidateCheck.Body.Close()
	routes.ServeHTTP(resValidateCheck, reqValidateCheck)
	assert.Equal(t, http.StatusBadRequest, resValidateCheck.Code)

	existentId := 1
	req, res := createRequestTest(http.MethodGet, getPathUrl(fmt.Sprintf("/employees/%d", existentId)), "")

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	var objRes employees.Employee

	err := json.Unmarshal(res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestUpdateOk(t *testing.T) {
	routes := createServer()
	reqFake, resFake := createRequestTest(http.MethodGet, getPathUrl("/employees/1"), "")

	defer reqFake.Body.Close()
	routes.ServeHTTP(resFake, reqFake)

	assert.Equal(t, http.StatusOK, resFake.Code)

	type createdEmployee struct {
		ID int `json:"id"`
		employees.Employee
	}

	var oldObjRes createdEmployee

	errFake := json.Unmarshal(resFake.Body.Bytes(), &oldObjRes)
	assert.Nil(t, errFake)

	req, res := createRequestTest(http.MethodPatch, getPathUrl(fmt.Sprintf("/employees/%d", oldObjRes.ID)),
		`
		{
			"card_number_id": "1234",
			"first_name": "Paloma",
			"last_name": "Ribeiro",
			"warehouse_id": 2
		}`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	var newObjRes createdEmployee

	err := json.Unmarshal(res.Body.Bytes(), &newObjRes)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotEqual(t, oldObjRes.CardNumberId, newObjRes.CardNumberId)
}

func TestUpdateNonExistent(t *testing.T) {
	routes := createServer()

	reqValidateCheck, resValidateCheck := createRequestTest(http.MethodGet, getPathUrl("/employees/notint"), "")

	defer reqValidateCheck.Body.Close()
	routes.ServeHTTP(resValidateCheck, reqValidateCheck)
	assert.Equal(t, http.StatusBadRequest, resValidateCheck.Code)

	inexistentId := 10
	req, res := createRequestTest(http.MethodPatch, getPathUrl(fmt.Sprintf("/employees/%d", inexistentId)),

		`{
			"card_number_id": "1234",
			"first_name": "Paloma",
			"last_name": "Ribeiro",
			"warehouse_id": 2
		}
		`,
	)

	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)

}
