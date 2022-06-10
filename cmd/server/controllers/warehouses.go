package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses"
)

type WarehouseController struct {
	service warehouses.Service
}

func NewWarehouse(w warehouses.Service) *WarehouseController {
	return &WarehouseController{service: w}
}

// @Summary Create warehouse
// @Tags Warehouses
// @Description Add a new warehouse checking for duplicate warehouses code before
// @Accept json
// @Produce json
// @Param warehouse body warehouses.CreateWarehouseInput true "Warehouse to create"
// @Success 201 {object} schemes.JSONSuccessResult{data=warehouses.Warehouse}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /warehouses [post]
func (wc *WarehouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var warehouseInput warehouses.CreateWarehouseInput
		if err := ctx.ShouldBindJSON(&warehouseInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		w, err := wc.service.Create(
			warehouseInput.WarehouseCode,
			warehouseInput.Address,
			warehouseInput.Telephone,
			warehouseInput.MinimumCapacity,
			warehouseInput.MinimumTemperature,
		)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		ctx.JSON(
			http.StatusCreated, w,
		)
	}
}

// @Summary List warehouses
// @Tags Warehouses
// @Description get all warehouses
// @Accept json
// @Produce json
// @Success 200 {object} schemes.JSONSuccessResult{data=warehouses.Warehouse}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /warehouses [get]
func (wc *WarehouseController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws, err := wc.service.GetAll()

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
// @Success 200 {object} schemes.JSONSuccessResult{data=warehouses.Warehouse}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /warehouses/{id} [get]
func (wc *WarehouseController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "invalid id type"},
			)
			return
		}

		w, err := wc.service.FindById(warehouseId)
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
// @Param warehouse body warehouses.UpdateWarehouseInput true "Warehouse to update"
// @Success 200 {object} schemes.JSONSuccessResult{data=warehouses.Warehouse}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /warehouses/{id} [patch]
func (wc *WarehouseController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "invalid id type"},
			)
			return
		}

		var warehouseInput warehouses.UpdateWarehouseInput
		if err := ctx.ShouldBindJSON(&warehouseInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		currentWarehouse, err := wc.service.FindById(warehouseId)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": "could not find warehouse"},
			)
			return
		}

		updatedWarehouse, err := wc.service.Update(
			*currentWarehouse,
			warehouseInput.WarehouseCode,
			warehouseInput.Address,
			warehouseInput.Telephone,
			warehouseInput.MinimumCapacity,
			warehouseInput.MinimumTemperature,
		)
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
// @Success 204 {object} schemes.JSONSuccessResult{data=string}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /warehouses/{id} [delete]
func (wc *WarehouseController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "invalid id type"},
			)
			return
		}

		if err := wc.service.Delete(warehouseId); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
