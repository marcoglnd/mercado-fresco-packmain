package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
)

func productsRouter(superRouter *gin.RouterGroup) {

	var listOfProducts []products.Product
	repo := products.NewRepository(listOfProducts)
	service := products.NewService(repo)
	controller := controllers.NewProduct(service)

	pr := superRouter.Group("/products")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/debug", Debug)
	}
}

func Debug(ctx *gin.Context) {
	ctx.JSON(http.StatusTeapot, gin.H{
		"debug": "is running",
	})
}
