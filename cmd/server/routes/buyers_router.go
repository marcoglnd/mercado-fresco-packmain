package routes

import (
	"database/sql"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/service"

	"github.com/gin-gonic/gin"
)

func buyersRouter(superRouter *gin.RouterGroup, DBConnection *sql.DB) {
	repository := mariadb.NewMariaDBRepository(DBConnection)

	buyerService := service.NewBuyerService(repository)

	buyerController, _ := controller.NewBuyerController(buyerService)
	pr := superRouter.Group("/buyers")
	{
		pr.GET("/", buyerController.GetAll())
		pr.GET("/:id", buyerController.GetById())
		pr.POST("/", buyerController.Create())
		pr.PATCH("/:id", buyerController.Update())
		pr.DELETE("/:id", buyerController.Delete())
		pr.GET("/reportPurchaseOrders", buyerController.ReportPurchaseOrders())
	}
}
