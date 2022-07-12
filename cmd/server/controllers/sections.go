package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections"

	"github.com/gin-gonic/gin"
)

type SectionsController struct {
	service sections.Service
}

func NewSection(b sections.Service) *SectionsController {
	return &SectionsController{
		service: b,
	}
}

// @Summary List sections
// @Tags Sections
// @Description get all sections
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JSONSuccessResult{data=schemas.Section}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /sections [get]
func (c *SectionsController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": b,
		})
	}
}

// @Summary Section by id
// @Tags Sections
// @Description get section by its id
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 200 {object} schemas.Section
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /sections/{id} [get]
func (c *SectionsController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sectionId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		b, err := c.service.GetById(sectionId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}

// @Summary Create section
// @Tags Sections
// @Description Add a new section to the list
// @Accept json
// @Produce json
// @Param section body requestSection true "Section to create"
// @Success 201 {object} schemas.Section
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Router /sections [post]
func (c *SectionsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSection
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}
		b, err := c.service.Create(
			req.SectionNumber, req.CurrentTemperature,
			req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity,
			req.WarehouseId, req.ProductTypeId)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, b)
	}
}

// @Summary Update section
// @Tags Sections
// @Description Update existing section in list
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Param section body requestSection true "Section to update"
// @Success 200 {object} schemas.Section
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /sections/{id} [patch]
func (c *SectionsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var req requestSection
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		b, err := c.service.Update(int(id),
			req.SectionNumber, req.CurrentTemperature,
			req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity,
			req.WarehouseId, req.ProductTypeId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}

// @Summary Delete section
// @Tags Sections
// @Description Delete existing sections in list
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 204 {object} schemas.JSONSuccessResult{data=string}
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /sections/{id} [delete]
func (c *SectionsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("A section %d foi removido", id)})
	}
}

type requestSection struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseId        int `json:"warehouse_id" binding:"required"`
	ProductTypeId      int `json:"product_type_id" binding:"required"`
}
