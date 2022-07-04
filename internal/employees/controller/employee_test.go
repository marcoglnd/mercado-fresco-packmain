package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatEmployeeOk(t *testing.T) {
	routes := createServer()

	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestCreateEmployeeConflict(t *testing.T) {
	routes := createServer()

	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
	}`)

	second_req, second_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Maria",
		"last_name": "Joana",
		"warehouse_id": 10
	}`)

	routes.ServeHTTP(res, req)
	routes.ServeHTTP(second_res, second_req)

	assert.Equal(t, http.StatusConflict, second_res.Code)

}

func TestCreateEmployeeUnprocessable(t *testing.T) {
	routes := createServer()

	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
	}`)

	second_req, second_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"last_name": "Joana",
		"warehouse_id": 10
	}`)

	third_req, third_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Maria",
		"last_name": "Joana",
	}`)

	routes.ServeHTTP(res, req)
	routes.ServeHTTP(second_res, second_req)
	routes.ServeHTTP(third_res, third_req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, second_res.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, third_res.Code)

}

func TestGetAllOk(t *testing.T) {
	routes := createServer()

	req, res := createRequestTest(http.MethodGet, getPathUrl("/employees/"), "")

	defer req.Body.Close()

	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	objRes := struct {
		Code int
		Data []employees.Employee
	}{}

	err := json.Unmarshal(res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) >= 0)
}

func TestGetEmployeeByIdOk(t *testing.T) {
	employee := &employees.Employee{
		ID:           1,
		CardNumberId: "1234",
		FirstName:    "Luiza",
		LastName:     "Maria",
		WarehouseId:  123,
	}

	mockService := new(mocks.Service)
	mockService.On("GetById", mock.AnythingOfType("int")).Return(*employee, nil)

	res := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(res)
	newEmployee := controllers.NewEmployee(mockService)

	engine.GET("/api/v1/employees/:id", newEmployee.GetById())
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
	assert.NoError(t, err)
	ctx.Request = request
	engine.ServeHTTP(res, ctx.Request)

	assert.Equal(t, http.StatusOK, res.Code)

	var objRes employees.Employee

	err = json.Unmarshal(res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	assert.True(t, objRes.CardNumberId == "1234")
	assert.True(t, objRes.FirstName == "Luiza")
	assert.True(t, objRes.LastName == "Maria")
	assert.True(t, objRes.WarehouseId == 123)

}

func TestGetEmployeeByIdNotFound(t *testing.T) {
	routes := createServer()

	post_req, post_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
	}`)

	get_req, get_res := createRequestTest(http.MethodGet, getPathUrl("/employees/42"), "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	routes.ServeHTTP(post_res, post_req)
	routes.ServeHTTP(get_res, get_req)

	assert.Equal(t, http.StatusNotFound, get_res.Code)
}

func TestGetEmployeeByIdBadRequest(t *testing.T) {
	routes := createServer()

	get_req, get_res := createRequestTest(http.MethodGet, getPathUrl("/employees/abc"), "")

	defer get_req.Body.Close()

	routes.ServeHTTP(get_res, get_req)

	assert.Equal(t, http.StatusBadRequest, get_res.Code)
}

func TestUpdateEmployeeOK(t *testing.T) {
	routes := createServer()

	post_req, post_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	patch_req, patch_res := createRequestTest(http.MethodPatch, getPathUrl("/employees/1"), `{
		"card_number_id": "1234",
		"first_name": "Maria",
		"last_name": "Silva",
		"warehouse_id": 4
		}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	routes.ServeHTTP(post_res, post_req)
	routes.ServeHTTP(patch_res, patch_req)

	assert.Equal(t, http.StatusOK, patch_res.Code)

	var objRes employees.Employee

	err := json.Unmarshal(patch_res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	assert.True(t, objRes.CardNumberId == "1234")
	assert.True(t, objRes.FirstName == "Maria")
	assert.True(t, objRes.LastName == "Silva")
	assert.True(t, objRes.WarehouseId == 4)
}

func TestUpdateEmployeeNotFound(t *testing.T) {
	routes := createServer()

	post_req, post_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	patch_req, patch_res := createRequestTest(http.MethodPatch, getPathUrl("/employees/10"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	routes.ServeHTTP(post_res, post_req)
	routes.ServeHTTP(patch_res, patch_req)

	assert.Equal(t, http.StatusNotFound, patch_res.Code)
}

func TestUpdateEmployeeBadRequest(t *testing.T) {
	routes := createServer()

	patch_req, patch_res := createRequestTest(http.MethodPatch, getPathUrl("/employees/abc"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	defer patch_req.Body.Close()

	routes.ServeHTTP(patch_res, patch_req)

	assert.Equal(t, http.StatusBadRequest, patch_res.Code)
}

//
func TestUpdateEmployeeUnprocessable(t *testing.T) {
	routes := createServer()

	post_req, post_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	patch_req, patch_res := createRequestTest(http.MethodPatch, getPathUrl("/employees/1"), `{
		"card_number_id": "1234",
		"last_name": "Rosas",
		"warehouse_id": 4
		}`)

	secondPatch_req, secondPatch_res := createRequestTest(http.MethodPatch, getPathUrl("/employees/1"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"warehouse_id": 3
		}`)

	thirdPatch_req, thirdPatch_res := createRequestTest(http.MethodPatch, getPathUrl("/employees/1"), `{
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()
	defer secondPatch_req.Body.Close()
	defer thirdPatch_req.Body.Close()

	routes.ServeHTTP(post_res, post_req)
	routes.ServeHTTP(patch_res, patch_req)
	routes.ServeHTTP(secondPatch_res, secondPatch_req)
	routes.ServeHTTP(thirdPatch_res, thirdPatch_req)

	assert.Equal(t, http.StatusUnprocessableEntity, patch_res.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, secondPatch_res.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, thirdPatch_res.Code)
}

func TestDeleteEmployeeOK(t *testing.T) {
	routes := createServer()

	post_req, post_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	get_req, get_res := createRequestTest(http.MethodGet, getPathUrl("/employees/"), "")

	routes.ServeHTTP(post_res, post_req)
	routes.ServeHTTP(get_res, get_req)

	objRes := struct {
		Code int
		Data []employees.Employee
	}{}

	err := json.Unmarshal(get_res.Body.Bytes(), &objRes)

	employeesLen := len(objRes.Data)

	assert.Nil(t, err)
	assert.True(t, employeesLen > 0)

	delete_req, delete_res := createRequestTest(http.MethodDelete, getPathUrl("/employees/1"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	routes.ServeHTTP(delete_res, delete_req)

	secondGet_req, secondGet_res := createRequestTest(http.MethodGet, getPathUrl("/employees/"), "")

	routes.ServeHTTP(secondGet_res, secondGet_req)

	secondObjRes := struct {
		Code int
		Data []employees.Employee
	}{}

	json.Unmarshal(secondGet_res.Body.Bytes(), &secondObjRes)

	secondEmployeesLen := len(secondObjRes.Data)

	assert.Equal(t, http.StatusNoContent, delete_res.Code)
	assert.True(t, employeesLen-1 == secondEmployeesLen)
}

func TestDeleteEmployeesFail(t *testing.T) {
	routes := createServer()

	post_req, post_res := createRequestTest(http.MethodPost, getPathUrl("/employees/"), `{
		"card_number_id": "1234",
		"first_name": "Julia",
		"last_name": "Rosas",
		"warehouse_id": 3
		}`)

	routes.ServeHTTP(post_res, post_req)

	delete_req, delete_res := createRequestTest(http.MethodDelete, getPathUrl("/employees/10"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	routes.ServeHTTP(delete_res, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_res.Code)
}

func TestDeleteEmployeesBadRequest(t *testing.T) {
	routes := createServer()

	delete_req, delete_res := createRequestTest(http.MethodDelete, getPathUrl("/employees/abc"), "")

	defer delete_req.Body.Close()

	routes.ServeHTTP(delete_res, delete_req)

	assert.Equal(t, http.StatusBadRequest, delete_res.Code)
}
