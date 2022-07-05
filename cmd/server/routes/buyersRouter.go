package routes

import (
	"database/sql"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/service"

	"github.com/gin-gonic/gin"
)

func buyersRouter(superRouter *gin.RouterGroup, DBConnection *sql.DB) {
	//1. repositório
	repository := mariadb.NewMariaDBRepository(DBConnection)

	//2. serviço (regra de negócio)
	buyerService := service.NewBuyerService(repository)

	//3. controller
	buyerController, _ := controller.NewBuyerController(buyerService)
	pr := superRouter.Group("/buyers")
	{
		pr.GET("/", buyerController.GetAll())
		pr.GET("/:id", buyerController.GetById())
		pr.POST("/", buyerController.Create())
		pr.PATCH("/:id", buyerController.Update())
		pr.DELETE("/:id", buyerController.Delete())
	}
}
