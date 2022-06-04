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

// CreateWarehouse godoc
// @Summary Creates a warehouse
// @Tags Warehouses
// @Description You can choose to create a warehouse with custom attributes, a unique and valid warehouse code should be defined
// @Accept json
// @Produce json
// @Param warehouse body request true "Warehouse to create"
// @Success 201 {object} gin.H
// @Failure 422 {object} gin.H
// @Router /warehouses [post]
func (wc *WarehouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type CreateWarehouseInput struct {
			WarehouseCode      string `json:"warehouse_code" binding:"required,len=3"`
			Address            string `json:"address" binding:"required"`
			Telephone          string `json:"telephone" binding:"required"`
			MinimumCapacity    int    `json:"minimum_capacity" binding:"required,gte=1"`
			MinimumTemperature int    `json:"minimum_temperature" binding:"required,gte=1"`
		}

		var warehouseInput CreateWarehouseInput
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
			http.StatusOK, ws,
		)
	}
}

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

		type UpdateWarehouseInput struct {
			WarehouseCode      string `json:"warehouse_code" binding:"len=3"`
			Address            string `json:"address"`
			Telephone          string `json:"telephone"`
			MinimumCapacity    int    `json:"minimum_capacity" binding:"gte=1"`
			MinimumTemperature int    `json:"minimum_temperature" binding:"gte=1"`
		}

		var warehouseInput UpdateWarehouseInput
		if err := ctx.ShouldBindJSON(&warehouseInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		updatedWarehouse, err := wc.service.Update(
			warehouseId,
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
