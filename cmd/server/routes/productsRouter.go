package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/controller"
)

func productsRouter(superRouter *gin.RouterGroup, conn *sql.DB) {
	repo := products.NewRepository(conn)
	service := products.NewService(repo)
	controller := controller.NewProduct(service)

	pr := superRouter.Group("/products")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.CreateNewProduct())
		pr.PATCH("/:id", controller.Update())
		pr.DELETE("/:id", controller.Delete())
	}
}
