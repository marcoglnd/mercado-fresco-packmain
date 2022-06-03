package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
)

func productsRouter(superRouter *gin.RouterGroup) {

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := controllers.NewProduct(service)

	pr := superRouter.Group("/products")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.CreateNewProduct())
		pr.GET("/debug", Debug)
	}
}

func Debug(ctx *gin.Context) {
	ctx.JSON(http.StatusTeapot, gin.H{
		"debug": "is running",
	})
}
