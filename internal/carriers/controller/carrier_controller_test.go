package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"
	mock "github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	carrierInput := utils.CreateRandomCarrier()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(carrierInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().IsCidAvailable(gomock.Any(), carrierInput.Cid).Return(nil)
	serviceMock.EXPECT().Create(gomock.Any(), &carrierInput).Return(&carrierInput, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", &buf)

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var objRes domain.Carrier
	err = json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.Equal(t, carrierInput.Cid, objRes.Cid)
}

func TestCreateConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	carrierInput := utils.CreateRandomCarrier()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(carrierInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().IsCidAvailable(gomock.Any(), carrierInput.Cid).Return(errors.New("duplicate"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", &buf)

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestCreateInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("invalid_data")))

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestCreateFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	carrierInput := utils.CreateRandomCarrier()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(carrierInput)
	assert.Nil(t, err)

	serviceMock.EXPECT().IsCidAvailable(gomock.Any(), carrierInput.Cid).Return(nil)
	serviceMock.EXPECT().Create(gomock.Any(), &carrierInput).Return(nil, errors.New("error saving"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodPost, "/", &buf)

	engine.POST("/", controller.Create())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestReportCarriersOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	serviceMock.EXPECT().GetAllCarriersReport(gomock.Any()).Return(&[]domain.CarrierReport{}, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/", nil)

	engine.GET("/", controller.ReportCarriers())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestReportCarriersFailValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/?id=notint", nil)

	engine.GET("/", controller.ReportCarriers())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestReportCarriersFailGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	serviceMock.EXPECT().GetAllCarriersReport(gomock.Any()).Return(nil, errors.New("some error"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/", nil)

	engine.GET("/", controller.ReportCarriers())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestReportCarriersByIdOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	serviceMock.EXPECT().GetCarriersReportById(gomock.Any(), int64(1)).Return(&domain.CarrierReport{}, nil)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/?id=1", nil)

	engine.GET("/", controller.ReportCarriers())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestReportCarriersByIdFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	serviceMock.EXPECT().GetCarriersReportById(
		gomock.Any(), int64(1),
	).Return(nil, errors.New("some error"))

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)

	req, err := http.NewRequest(http.MethodGet, "/?id=1", nil)

	engine.GET("/", controller.ReportCarriers())
	engine.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
