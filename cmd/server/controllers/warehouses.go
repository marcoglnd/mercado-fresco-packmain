package controllers

import (
	"net/http"

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
// @Success 201 {object} web.Response
// @Failure 422 {object} web.Response
// @Router /warehouses [post]
func (wc *WarehouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouse := &warehouses.Warehouse{}
		if err := ctx.ShouldBindJSON(&warehouse); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		warehouseDuplicated, err := wc.service.FindByWarehouseCode(warehouse.WarehouseCode)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if warehouseDuplicated != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "warehouseCode already exists"})
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
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusCreated, w,
		)
	}
}
