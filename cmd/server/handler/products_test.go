package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := controllers.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.CreateNewProduct())
		pr.PATCH("/:id", controller.Update())
		pr.DELETE("/:id", controller.Delete())
	}

	return router
}
