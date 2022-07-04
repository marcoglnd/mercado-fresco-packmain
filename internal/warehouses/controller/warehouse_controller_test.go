package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/domain"
	mock "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseInput := domain.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(warehouseInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().IsWarehouseCodeAvailable(gomock.Any(), warehouseInput.WarehouseCode).Return(nil)
	serviceMock.EXPECT().Create(gomock.Any(), &warehouseInput).Return(&warehouseInput, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", &buf)

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var objRes domain.Warehouse
	err = json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.Equal(t, warehouseInput.WarehouseCode, objRes.WarehouseCode)
}

func TestCreateInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("invalid_data")))

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestCreateConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseInput := domain.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(warehouseInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().IsWarehouseCodeAvailable(gomock.Any(), warehouseInput.WarehouseCode).Return(errors.New("duplicate"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", &buf)

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestCreateFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseInput := domain.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(warehouseInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().IsWarehouseCodeAvailable(gomock.Any(), warehouseInput.WarehouseCode).Return(nil)
	serviceMock.EXPECT().Create(gomock.Any(), &warehouseInput).Return(nil, errors.New("error saving"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", &buf)

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestGetAllFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	serviceMock.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("error getting all"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/", nil)

	engine.GET("/", controller.GetAll())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestGetAllOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	serviceMock.EXPECT().GetAll(gomock.Any()).Return(&[]domain.Warehouse{}, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/", nil)

	engine.GET("/", controller.GetAll())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetByIdInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/invalid_string_id", nil)

	engine.GET("/:id", controller.GetById())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetByIdFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseFakeInput := 1
	serviceMock.EXPECT().FindById(gomock.Any(), int64(warehouseFakeInput)).Return(nil, errors.New("fail finding id"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%d", warehouseFakeInput), nil)

	engine.GET("/:id", controller.GetById())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGetByIdOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseFakeInput := 1
	serviceMock.EXPECT().FindById(gomock.Any(), int64(warehouseFakeInput)).Return(&domain.Warehouse{}, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%d", warehouseFakeInput), nil)

	engine.GET("/:id", controller.GetById())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateIdInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPatch, "/invalid_string_id", nil)

	engine.PATCH("/:id", controller.Update())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPatch, "/1", bytes.NewBuffer([]byte("invalid_data")))

	engine.PATCH("/:id", controller.Update())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestUpdateNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseInput := domain.Warehouse{
		ID:                 int64(1),
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(warehouseInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().FindById(gomock.Any(), warehouseInput.ID).Return(nil, errors.New("fail finding id"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/%d", warehouseInput.ID), &buf)

	engine.PATCH("/:id", controller.Update())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUpdateFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseInput := domain.Warehouse{
		ID:                 int64(1),
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(warehouseInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().FindById(gomock.Any(), warehouseInput.ID).Return(&warehouseInput, nil)
	serviceMock.EXPECT().Update(gomock.Any(), &warehouseInput).Return(nil, errors.New("fail update"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/%d", warehouseInput.ID), &buf)

	engine.PATCH("/:id", controller.Update())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseInput := domain.Warehouse{
		ID:                 int64(1),
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(warehouseInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().FindById(gomock.Any(), warehouseInput.ID).Return(&warehouseInput, nil)
	serviceMock.EXPECT().Update(gomock.Any(), &warehouseInput).Return(&warehouseInput, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/%d", warehouseInput.ID), &buf)

	engine.PATCH("/:id", controller.Update())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodDelete, "/invalid_string_id", nil)

	engine.DELETE("/:id", controller.Delete())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseFakeInput := 1
	serviceMock.EXPECT().Delete(gomock.Any(), int64(warehouseFakeInput)).Return(errors.New("fail finding id"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/%d", warehouseFakeInput), nil)

	engine.DELETE("/:id", controller.Delete())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestDeleteOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockWarehouseService(ctrl)
	controller := controller.NewWarehouseController(serviceMock)

	warehouseFakeInput := 1
	serviceMock.EXPECT().Delete(gomock.Any(), int64(warehouseFakeInput)).Return(nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/%d", warehouseFakeInput), nil)

	engine.DELETE("/:id", controller.Delete())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
