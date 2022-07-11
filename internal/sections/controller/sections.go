package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
)

type SectionsController struct {
	service domain.Service
}

func NewSection(s domain.Service) *SectionsController {
	return &SectionsController{
		service: s,
	}
}

// @Summary List sections
// @Tags Sections
// @Description get all sections
// @Accept json
// @Produce json
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Section}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sections [get]
func (c *SectionsController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// @Summary Section by id
// @Tags Sections
// @Description get section by its id
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 200 {object} schemes.Section
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sections/{id} [get]
func (c *SectionsController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// sectionId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		var req domain.RequestSectionId
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		section, err := c.service.GetById(ctx.Request.Context(), req.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, section)

		// vraw
		// sectionId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		// 	return
		// }
		// b, err := c.service.GetById(ctx, sectionId)
		// if err != nil {
		// 	ctx.JSON(http.StatusNotFound, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }
		// ctx.JSON(http.StatusOK, b)
		// vraw
	}
}

// @Summary Create section
// @Tags Sections
// @Description Add a new section to the list
// @Accept json
// @Produce json
// @Param section body requestSection true "Section to create"
// @Success 201 {object} schemes.Section
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /sections [post]
func (c *SectionsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestSections
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		section, err := c.service.Create(
			ctx.Request.Context(),
			&domain.Section{
				SectionNumber:      req.SectionNumber,
				CurrentCapacity:    req.CurrentCapacity,
				MinimumCapacity:    req.MinimumCapacity,
				MaximumCapacity:    req.MaximumCapacity,
				WarehouseId:        req.WarehouseId,
				ProductTypeId:      req.ProductTypeId,
				CurrentTemperature: req.CurrentTemperature,
				MinimumTemperature: req.MinimumTemperature,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, section)

		// vraw
		// if err := ctx.ShouldBindJSON(&req); err != nil {
		// 	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }
		// b, err := c.service.Create(
		// 	ctx,
		// 	req.SectionNumber, req.CurrentTemperature,
		// 	req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity,
		// 	req.WarehouseId, req.ProductTypeId)
		// if err != nil {
		// 	ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		// 	return
		// }
		// ctx.JSON(http.StatusCreated, b)
		// vraw
	}
}

// @Summary Update section
// @Tags Sections
// @Description Update existing section in list
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Param section body requestSection true "Section to update"
// @Success 200 {object} schemes.Section
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sections/{id} [patch]
func (c *SectionsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestSectionsUpdated

		// vraw if duvidoso
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		var reqId domain.RequestSectionId
		if err := ctx.ShouldBindUri(&reqId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		section, err := c.service.Update(
			ctx.Request.Context(),
			&domain.Section{
				ID:                 reqId.ID,
				SectionNumber:      req.SectionNumber,
				CurrentCapacity:    req.CurrentCapacity,
				MinimumCapacity:    req.MinimumCapacity,
				MaximumCapacity:    req.MaximumCapacity,
				WarehouseId:        req.WarehouseId,
				ProductTypeId:      req.ProductTypeId,
				CurrentTemperature: req.CurrentTemperature,
				MinimumTemperature: req.MinimumTemperature,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, section)

		// vraw
		// id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		// 	return
		// }

		// var req requestSection
		// if err := ctx.ShouldBindJSON(&req); err != nil {
		// 	ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// 	return
		// }

		// b, err := c.service.Update(
		// 	ctx,
		// 	int(id),
		// 	req.SectionNumber, req.CurrentTemperature,
		// 	req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity,
		// 	req.WarehouseId, req.ProductTypeId)
		// if err != nil {
		// 	ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// 	return
		// }
		// ctx.JSON(http.StatusOK, b)
		// vraw
	}
}

// @Summary Delete section
// @Tags Sections
// @Description Delete existing sections in list
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 204 {object} schemes.JSONSuccessResult{data=string}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sections/{id} [delete]
func (c *SectionsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestSectionId
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		}

		err := c.service.Delete(ctx.Request.Context(), req.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("A section %d foi removida", req)})

		// vraw
		// id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		// 	return
		// }
		// err = c.service.Delete(ctx, int(id))
		// if err != nil {
		// 	ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// 	return
		// }
		// ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("A section %d foi removido", id)})
		// vraw
	}
}

// vraw
// type requestSection struct {
// 	SectionNumber      int64 `json:"section_number" binding:"required"`
// 	CurrentTemperature int64 `json:"current_temperature" binding:"required"`
// 	MinimumTemperature int64 `json:"minimum_temperature" binding:"required"`
// 	CurrentCapacity    int64 `json:"current_capacity" binding:"required"`
// 	MinimumCapacity    int64 `json:"minimum_capacity" binding:"required"`
// 	MaximumCapacity    int64 `json:"maximum_capacity" binding:"required"`
// 	WarehouseId        int64 `json:"warehouse_id" binding:"required"`
// 	ProductTypeId      int64 `json:"product_type_id" binding:"required"`
// }
// vraw
