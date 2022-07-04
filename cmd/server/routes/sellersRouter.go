package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/service"
)

func sellersRouter(superRouter *gin.RouterGroup, DBConnection *sql.DB) {
	//1. repositório
	repository := mariadb.NewMariaDBRepository(DBConnection)

	//2. serviço (regra de negócio)
	sellerService := service.NewService(repository)

	//3. controller
	sellerController, _ := controller.NewSellerController(sellerService)

	pr := superRouter.Group("/sellers")
	{
		pr.GET("/", sellerController.GetAll())
		pr.GET("/:id", sellerController.GetByID())
		pr.POST("/", sellerController.Create())
		pr.PATCH("/:id", sellerController.Update())
		pr.DELETE("/:id", sellerController.Delete())
	}
}
