package controller_test

import (
	"bytes"
	"encoding/json"
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
