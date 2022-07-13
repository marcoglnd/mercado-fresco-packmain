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

	sl := superRouter.Group("/sellers")
	{
		sl.GET("/", sellerController.GetAll())
		sl.GET("/:id", sellerController.GetByID())
		sl.POST("/", sellerController.Create())
		sl.PATCH("/:id", sellerController.Update())
		sl.DELETE("/:id", sellerController.Delete())
	}
}
