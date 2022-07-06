package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateSections_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateSections_bad_request(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"warehouse_id": 1
	  }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_CreateSections_conflict(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }`)

	second_req, second_rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)

	assert.Equal(t, http.StatusConflict, second_rr.Code)
}

func Test_GetSections_OK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, getPathUrl("/sections/"), "")

	defer req.Body.Close()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []sections.Section
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) >= 0)
}

func Test_GetSectionsById_OK(t *testing.T) {

	section := &sections.Section{
		ID:                 1,
		CurrentCapacity:    1,
		CurrentTemperature: 1,
		MaximumCapacity:    1,
		MinimumCapacity:    1,
		MinimumTemperature: 1,
		ProductTypeId:      1,
		SectionNumber:      1,
		WarehouseId:        1,
	}

	mockService := new(mocks.Service)
	mockService.On("GetById", mock.AnythingOfType("int")).Return(*section, nil)

	res := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(res)
	newSection := controllers.NewSection(mockService)

	engine.GET("/api/v1/sections/:id", newSection.GetById())
	request, err := http.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
	assert.NoError(t, err)
	ctx.Request = request
	engine.ServeHTTP(res, ctx.Request)

	assert.Equal(t, http.StatusOK, res.Code)

	var objRes sections.Section

	err = json.Unmarshal(res.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
	assert.True(t, objRes.CurrentCapacity == 1)
	assert.True(t, objRes.CurrentTemperature == 1)
	assert.True(t, objRes.MaximumCapacity == 1)
	assert.True(t, objRes.MinimumCapacity == 1)
	assert.True(t, objRes.MinimumTemperature == 1)
	assert.True(t, objRes.ProductTypeId == 1)
	assert.True(t, objRes.SectionNumber == 1)
	assert.True(t, objRes.WarehouseId == 1)
}

func Test_GetSectionsById_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }}`)

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/sections/10"), "")

	defer post_req.Body.Close()
	defer get_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	assert.Equal(t, http.StatusNotFound, get_rr.Code)
}

func Test_UpdateSections_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/sections/1"), `{
		"current_capacity": 3,
		"current_temperature": 3,
		"maximum_capacity": 3,
		"minimum_capacity": 3,
		"minimum_temperature": 3,
		"product_type_id": 3,
		"section_number": 3,
		"warehouse_id": 3
	  }`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusOK, patch_rr.Code)

	var objRes sections.Section

	err := json.Unmarshal(patch_rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, objRes.ID == 1)
}

func Test_UpdateSections_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }}`)

	patch_req, patch_rr := createRequestTest(http.MethodPatch, getPathUrl("/sections/10"), `{
		"current_capacity": 4,
		"current_temperature": 4,
		"maximum_capacity": 4,
		"minimum_capacity": 4,
		"minimum_temperature": 4,
		"product_type_id": 4,
		"section_number": 4,
		"warehouse_id": 4
	  }`)

	defer post_req.Body.Close()
	defer patch_req.Body.Close()

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(patch_rr, patch_req)

	assert.Equal(t, http.StatusNotFound, patch_rr.Code)
}

func Test_DeleteSections_OK(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }`)

	get_req, get_rr := createRequestTest(http.MethodGet, getPathUrl("/sections/"), "")

	r.ServeHTTP(post_rr, post_req)
	r.ServeHTTP(get_rr, get_req)

	objRes := struct {
		Code int
		Data []sections.Section
	}{}

	err := json.Unmarshal(get_rr.Body.Bytes(), &objRes)

	sectionsLen := len(objRes.Data)

	assert.Nil(t, err)
	assert.True(t, sectionsLen > 0)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/sections/1"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	secondGet_req, secondGet_rr := createRequestTest(http.MethodGet, getPathUrl("/sections/"), "")

	r.ServeHTTP(secondGet_rr, secondGet_req)

	secondObjRes := struct {
		Code int
		Data []sections.Section
	}{}

	json.Unmarshal(secondGet_rr.Body.Bytes(), &secondObjRes)

	secondsectionsLen := len(secondObjRes.Data)

	assert.Equal(t, http.StatusNoContent, delete_rr.Code)
	assert.True(t, sectionsLen-1 == secondsectionsLen)
}

func Test_DeleteSections_fail(t *testing.T) {
	r := createServer()

	post_req, post_rr := createRequestTest(http.MethodPost, getPathUrl("/sections/"), `{
		"current_capacity": 1,
		"current_temperature": 1,
		"maximum_capacity": 1,
		"minimum_capacity": 1,
		"minimum_temperature": 1,
		"product_type_id": 1,
		"section_number": 1,
		"warehouse_id": 1
	  }`)

	r.ServeHTTP(post_rr, post_req)

	delete_req, delete_rr := createRequestTest(http.MethodDelete, getPathUrl("/sections/10"), "")

	defer post_req.Body.Close()
	defer delete_req.Body.Close()

	r.ServeHTTP(delete_rr, delete_req)

	assert.Equal(t, http.StatusNotFound, delete_rr.Code)
}
