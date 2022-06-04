package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses"
)

func warehousesRouter(superRouter *gin.RouterGroup) {
	repository := warehouses.NewRepository()
	service := warehouses.NewService(repository)
	w := controllers.NewWarehouse(service)

	pr := superRouter.Group("/warehouses")
	{
		pr.GET("/debug", func(ctx *gin.Context) {
			ctx.JSON(http.StatusTeapot, gin.H{
				"debug": "is running",
			})
		})
		pr.POST("/", w.Create())
		pr.GET("/", w.GetAll())
		pr.GET("/:id", w.GetById())
		pr.PATCH("/:id", w.Update())
	}
}
