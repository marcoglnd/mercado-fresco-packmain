package controller

import (
	"bytes"
	"encoding/json"
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
		mockLocality := utils.CreateRandomLocality()
		mockLocalityId := utils.RandomInt64()
		params := domain.Locality{
			LocalityName: mockLocality.LocalityName,
			ProvinceID:   mockLocality.ID,
		}
		localityServiceMock.On("CreateLocality",
			mock.Anything,
			mock.Anything,
		).Return(mockLocalityId, nil).Once().
			On("GetLocalityByID",
				mock.Anything,
				mock.AnythingOfType("int64"),
			).Return(&mockLocality, nil).Once()

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

	// t.Run("fail", func(t *testing.T) {
	// 	localityServiceMock := mocks.NewLocalityService(t)
	// 	mockBadLocality := &domain.Locality{}

	// 	localityServiceMock.On("CreateLocality",
	// 		mock.Anything,
	// 		mock.Anything,
	// 	).Return(mockBadLocality, errors.New("error: unprocessable entity")).Maybe()

	// 	payload, err := json.Marshal(mockBadLocality)
	// 	assert.NoError(t, err)

	// 	req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", bytes.NewBuffer(payload))
	// 	rec := httptest.NewRecorder()

	// 	_, engine := gin.CreateTestContext(rec)

	// 	localityController := LocalityController{service: localityServiceMock}

	// 	engine.POST("/api/v1/localities", localityController.CreateLocality())

	// 	engine.ServeHTTP(rec, req)

	// 	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	// 	localityServiceMock.AssertExpectations(t)
	// })

	// t.Run("conflict", func(t *testing.T) {
	// 	localityServiceMock := mocks.NewLocalityService(t)
	// 	mockLocality := utils.CreateRandomLocality()

	// 	localityServiceMock.On("CreateLocality",
	// 		mock.Anything,
	// 		mock.Anything,
	// 	).Return(nil, domain.ErrIDNotFound).Maybe()

	// 	payload, err := json.Marshal(mockLocality)
	// 	assert.NoError(t, err)

	// 	req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", bytes.NewBuffer(payload))
	// 	rec := httptest.NewRecorder()

	// 	_, engine := gin.CreateTestContext(rec)

	// 	localityController := LocalityController{service: localityServiceMock}

	// 	engine.POST("/api/v1/localities", localityController.CreateLocality())

	// 	engine.ServeHTTP(rec, req)

	// 	assert.Equal(t, http.StatusConflict, rec.Code)

	// 	localityServiceMock.AssertExpectations(t)
	// })
}

// func TestGetQtyOfSellers(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		localityServiceMock := mocks.NewLocalityService(t)
// 		mockQtyOfSellers := utils.CreateRandomQtyOfSellers()
// 		mockQtyOfRecordsId := utils.RandomInt64()

// 		localityServiceMock.On("GetQtyOfSellers",
// 			mock.Anything,
// 			mock.AnythingOfType("int64"),
// 		).Return(&mockQtyOfSellers, nil).Once()

// 		payload, err := json.Marshal(mockQtyOfSellers)
// 		assert.NoError(t, err)
// 		PATH := fmt.Sprintf("/api/v1/localities/reportSellers?id=%v", mockQtyOfRecordsId)
// 		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
// 		rec := httptest.NewRecorder()

// 		_, engine := gin.CreateTestContext(rec)

// 		localityController := LocalityController{service: localityServiceMock}

// 		engine.GET("/api/v1/localities/reportSellers", localityController.service.GetLocalityByID())

// 		engine.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusOK, rec.Code)

// 		localityServiceMock.AssertExpectations(t)
// 	})
// }
