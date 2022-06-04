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
		warehouse := &warehouses.Warehouse{}
		if err := ctx.ShouldBindJSON(&warehouse); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		w, err := wc.service.Create(
			warehouse.WarehouseCode,
			warehouse.Address,
			warehouse.Telephone,
			warehouse.MinimumCapacity,
			warehouse.MinimumTemperature,
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
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(
			http.StatusOK, w,
		)
	}
}
