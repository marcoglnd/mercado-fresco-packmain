package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/db"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/repository"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/service"
)

func warehousesRouter(superRouter *gin.RouterGroup) {
	repository := repository.NewWarehouseRepository(db.GetDBConnection())
	service := service.NewWarehouseService(repository)
	warehouseController := controller.NewWarehouseController(service)

	pr := superRouter.Group("/warehouses")
	{
		pr.GET("/debug", func(ctx *gin.Context) {
			ctx.JSON(http.StatusTeapot, gin.H{
				"debug": "is running",
			})
		})
		pr.POST("/", warehouseController.Create())
		pr.GET("/", warehouseController.GetAll())
		pr.GET("/:id", warehouseController.GetById())
		pr.PATCH("/:id", warehouseController.Update())
		pr.DELETE("/:id", warehouseController.Delete())
	}
}
