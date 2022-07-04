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
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mock.NewMockCarrierService(ctrl)
	controller := controller.NewCarrierController(serviceMock)

	carrierInput := domain.Carrier{
		Cid:         "CID#1",
		CompanyName: "Some Company",
		Address:     "Rua Sao Paulo 200",
		Telephone:   "30021025",
		LocalityId:  1,
	}
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

	carrierInput := domain.Carrier{
		Cid:         "CID#1",
		CompanyName: "Some Company",
		Address:     "Rua Sao Paulo 200",
		Telephone:   "30021025",
		LocalityId:  1,
	}
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

	carrierInput := domain.Carrier{
		Cid:         "CID#1",
		CompanyName: "Some Company",
		Address:     "Rua Sao Paulo 200",
		Telephone:   "30021025",
		LocalityId:  1,
	}
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
