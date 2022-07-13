package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/service"
)

func productsRouter(superRouter *gin.RouterGroup, conn *sql.DB) {
	repo := mariadb.NewMariaDBRepository(conn)
	service := service.NewService(repo)
	controller := controller.NewProduct(service)

	superRouter.POST("/productRecords", controller.CreateProductRecords())
	superRouter.POST("/productBatches", controller.CreateProductBatches())

	pr := superRouter.Group("/products")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.CreateNewProduct())
		pr.PATCH("/:id", controller.Update())
		pr.DELETE("/:id", controller.Delete())
		pr.GET("/reportRecords", controller.GetQtyOfRecordsById())
		pr.GET("/reportProducts", controller.GetQtdProductsBySectionId())
	}
}
