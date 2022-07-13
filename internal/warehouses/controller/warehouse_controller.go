package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/domain"
)

type WarehouseController struct {
	service domain.WarehouseService
}

func NewWarehouseController(ws domain.WarehouseService) *WarehouseController {
	return &WarehouseController{service: ws}
}

// @Summary Create warehouse
// @Tags Warehouses
// @Description Add a new warehouse checking for duplicate warehouses code before
// @Accept json
// @Produce json
// @Param warehouse body domain.CreateWarehouseInput true "Warehouse to create"
// @Success 201 {object} domain.Warehouse
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Failure 409 {object} schemas.JSONBadReqResult{error=string}
// @Router /warehouses [post]
func (wc *WarehouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var warehouseInput domain.CreateWarehouseInput
		if err := ctx.ShouldBindJSON(&warehouseInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		if err := wc.service.IsWarehouseCodeAvailable(ctx, warehouseInput.WarehouseCode); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusConflict,
				gin.H{"error": err.Error()},
			)
			return
		}

		warehouse := domain.Warehouse{
			WarehouseCode:      warehouseInput.WarehouseCode,
			Address:            warehouseInput.Address,
			Telephone:          warehouseInput.Telephone,
			MinimumCapacity:    warehouseInput.MinimumCapacity,
			MinimumTemperature: warehouseInput.MinimumTemperature,
		}

		createdWarehouse, err := wc.service.Create(ctx, &warehouse)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		ctx.JSON(
			http.StatusCreated, createdWarehouse,
		)
	}
}

// @Summary List warehouses
// @Tags Warehouses
// @Description get all warehouses
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JSONSuccessResult{data=[]domain.Warehouse}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Router /warehouses [get]
func (wc *WarehouseController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws, err := wc.service.GetAll(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		ctx.JSON(
			http.StatusOK, gin.H{
				"data": ws,
			},
		)
	}
}

// @Summary Warehouse by id
// @Tags Warehouses
// @Description get warehouse by it's id
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /warehouses/{id} [get]
func (wc *WarehouseController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "invalid id type"},
			)
			return
		}

		w, err := wc.service.FindById(ctx, warehouseId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(
			http.StatusOK, w,
		)
	}
}

// @Summary Update warehouse
// @Tags Warehouses
// @Description Update existing warehouse in list checking for duplicate warehouses code
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Param warehouse body domain.UpdateWarehouseInput true "Warehouse to update"
// @Success 200 {object} domain.Warehouse
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /warehouses/{id} [patch]
func (wc *WarehouseController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "invalid id type"},
			)
			return
		}

		var warehouseInput domain.UpdateWarehouseInput
		if err := ctx.ShouldBindJSON(&warehouseInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		if _, err := wc.service.FindById(ctx, warehouseId); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": "could not find warehouse"},
			)
			return
		}

		warehouse := domain.Warehouse{
			ID:                 warehouseId,
			WarehouseCode:      warehouseInput.WarehouseCode,
			Address:            warehouseInput.Address,
			Telephone:          warehouseInput.Telephone,
			MinimumCapacity:    warehouseInput.MinimumCapacity,
			MinimumTemperature: warehouseInput.MinimumTemperature,
		}

		updatedWarehouse, err := wc.service.Update(ctx, &warehouse)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(
			http.StatusOK, updatedWarehouse,
		)
	}
}

// @Summary Delete warehouse
// @Tags Warehouses
// @Description Delete existing warehouse in list
// @Accept json
// @Produce json
// @Param id path int true "warehouse ID"
// @Success 204 {object} schemas.JSONBadReqResult{error=string}
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /warehouses/{id} [delete]
func (wc *WarehouseController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "invalid id type"},
			)
			return
		}

		if err := wc.service.Delete(ctx, warehouseId); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
