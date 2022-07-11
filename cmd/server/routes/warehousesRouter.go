package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/controller"
	repository "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/service"
)

func warehousesRouter(superRouter *gin.RouterGroup, dbConnection *sql.DB) {
	repository := repository.NewWarehouseRepository(dbConnection)
	service := service.NewWarehouseService(repository)
	warehouseController := controller.NewWarehouseController(service)

	pr := superRouter.Group("/warehouses")
	{
		pr.POST("/", warehouseController.Create())
		pr.GET("/", warehouseController.GetAll())
		pr.GET("/:id", warehouseController.GetById())
		pr.PATCH("/:id", warehouseController.Update())
		pr.DELETE("/:id", warehouseController.Delete())
	}
}
