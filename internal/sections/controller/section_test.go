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
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewSections(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, nil).Once()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.POST("/api/v1/sections", sectionController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with unprocessable entity", func(t *testing.T) {
		sectionsServiceMock := mocks.NewService(t)
		mockSectionBad := &domain.Section{}

		sectionsServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(mockSectionBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockSectionBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.POST("/api/v1/sections", sectionController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with status conflict", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.POST("/api/v1/sections", sectionController.Create())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockSection := utils.CreateRandomListSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("GetAll",
			mock.Anything,
		).Return(&mockSection, nil).Once()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.GET("/api/v1/sections", sectionController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of internal server error", func(t *testing.T) {
		sectionsServiceMock := mocks.NewService(t)
		mockSectionBad := &[]domain.Section{}

		sectionsServiceMock.On("GetAll",
			mock.Anything,
		).Return(mockSectionBad, errors.New("Internal server error")).Maybe()

		payload, err := json.Marshal(mockSectionBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.GET("/api/v1/sections", sectionController.GetAll())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})
}

func TestGetByI(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(&mockSection, nil).Once()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sections/%v", mockSection.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.GET("/api/v1/sections/:id", sectionController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid section id", func(t *testing.T) {
		sectionsServiceMock := mocks.NewService(t)
		mockSectionBad := &domain.Section{}

		sectionsServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockSectionBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockSectionBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sections/%v", "a")
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.GET("/api/v1/sections/:id", sectionController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting section", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("GetById",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil, errors.New("expected conflict error")).Maybe()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sections/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.GET("/api/v1/sections/:id", sectionController.GetById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, nil).Once()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sections/%v", mockSection.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.PATCH("/api/v1/sections/:id", sectionController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of unprocessable entity", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(&mockSection, errors.New("unprocessable entity")).Maybe()

		PATH := fmt.Sprintf("/api/v1/sections/%v", mockSection.ID)
		req := httptest.NewRequest(http.MethodPatch, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.PATCH("/api/v1/sections/:id", sectionController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid section id", func(t *testing.T) {
		sectionsServiceMock := mocks.NewService(t)
		mockSectionBad := &domain.Section{}

		sectionsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(mockSectionBad, errors.New("bad request")).Maybe()

		payload, err := json.Marshal(mockSectionBad)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sections/%v", "a")
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.PATCH("/api/v1/sections/:id", sectionController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting section", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Update",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("expected not found error")).Maybe()

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/api/v1/sections/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodPatch, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.PATCH("/api/v1/sections/:id", sectionController.Update())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockSection := utils.CreateRandomSection()
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(nil).Once()

		PATH := fmt.Sprintf("/api/v1/sections/%v", mockSection.ID)
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionsController := SectionsController{service: sectionsServiceMock}

		engine.DELETE("/api/v1/sections/:id", sectionsController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of invalid section id", func(t *testing.T) {
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Delete",
			mock.Anything,
			mock.Anything,
		).Return(errors.New("bad request")).Maybe()

		PATH := fmt.Sprintf("/api/v1/sections/%v", "a")
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.DELETE("/api/v1/sections/:id", sectionController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("In case of nonexisting section", func(t *testing.T) {
		sectionsServiceMock := mocks.NewService(t)

		sectionsServiceMock.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int64"),
		).Return(errors.New("expected conflict error")).Maybe()

		PATH := fmt.Sprintf("/api/v1/sections/%v", utils.RandomInt64())
		req := httptest.NewRequest(http.MethodDelete, PATH, nil)
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		sectionController := SectionsController{service: sectionsServiceMock}

		engine.DELETE("/api/v1/sections/:id", sectionController.Delete())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

}
