package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLocality(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)
		mockLocality := utils.CreateRandomGetLocality()
		mockLocalityId := utils.RandomInt64()
		params := requestCreateLocality{
			LocalityName: mockLocality.LocalityName,
			ProvinceID:   utils.RandomInt64(),
		}
		localityServiceMock.On("GetLocalityByID",
			mock.Anything,
			mock.Anything,
		).Return(&mockLocality, nil).Once()
		localityServiceMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(mockLocalityId, nil).Maybe()

		payload, err := json.Marshal(params)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.POST("/api/v1/localities", localityController.CreateLocality())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)
		mockBadLocality := &domain.Locality{}

		localityServiceMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(mockBadLocality, errors.New("error: unprocessable entity")).Maybe()

		payload, err := json.Marshal(mockBadLocality)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.POST("/api/v1/localities", localityController.CreateLocality())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})

	t.Run("conflict", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)
		mockLocality := utils.CreateRandomLocality()

		localityServiceMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(int64(0), domain.ErrIDNotFound).Maybe()

		payload, err := json.Marshal(mockLocality)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.POST("/api/v1/localities", localityController.CreateLocality())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})

	t.Run("internal error", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)
		mockLocality := utils.CreateRandomLocality()

		localityServiceMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(utils.RandomInt64(), nil).Maybe()
		localityServiceMock.On("GetLocalityByID",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("internal error")).Once()

		payload, err := json.Marshal(mockLocality)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.POST("/api/v1/localities", localityController.CreateLocality())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})
}

func TestGetAllQtyOfSellers(t *testing.T) {
	t.Run("success - GetAllQtyOfSellers", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)
		mockQtyOfSellers := utils.CreateRandomListQtyOfSellers()

		localityServiceMock.On("GetAllQtyOfSellers",
			mock.Anything,
		).Return(&mockQtyOfSellers, nil).Once()

		payload, err := json.Marshal(mockQtyOfSellers)
		assert.NoError(t, err)
		PATH := "/api/v1/localities/reportSellers"
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.GET("/api/v1/localities/reportSellers", localityController.GetAllQtyOfSellers())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})

	t.Run("internal sever error - GetAllQtyOfSellers", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)

		localityServiceMock.On("GetAllQtyOfSellers",
			mock.Anything,
		).Return(nil, errors.New("couldn`t return a list")).Once()

		PATH := "/api/v1/localities/reportSellers"
		req := httptest.NewRequest(http.MethodGet, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.GET("/api/v1/localities/reportSellers", localityController.GetAllQtyOfSellers())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})

	t.Run("success - GetQtyOfSellersByLocalityId", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)
		mockQtyOfSellers := utils.CreateRandomQtyOfSellers()

		localityServiceMock.On("GetQtyOfSellersByLocalityId",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockQtyOfSellers, nil).Once()

		payload, err := json.Marshal(mockQtyOfSellers)
		assert.NoError(t, err)
		PATH := fmt.Sprintf("/api/v1/localities/reportSellers?id=%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.GET("/api/v1/localities/reportSellers", localityController.GetAllQtyOfSellers())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})

	t.Run("not found - GetQtyOfSellersByLocalityId", func(t *testing.T) {
		localityServiceMock := mocks.NewLocalityService(t)

		localityServiceMock.On("GetQtyOfSellersByLocalityId",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("id not found")).Once()

		PATH := fmt.Sprintf("/api/v1/localities/reportSellers?id=%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		localityController := LocalityController{service: localityServiceMock}

		engine.GET("/api/v1/localities/reportSellers", localityController.GetAllQtyOfSellers())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		localityServiceMock.AssertExpectations(t)
	})
}
