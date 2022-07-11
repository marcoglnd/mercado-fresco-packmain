package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/service"
)

func localitiesRouter(superRouter *gin.RouterGroup, DBConnection *sql.DB) {
	//1. repositório
	repository := mariadb.NewMariaDBRepository(DBConnection)

	//2. serviço (regra de negócio)
	localityService := service.NewService(repository)

	//3. controller
	localityController, _ := controller.NewLocalityController(localityService)

	ll := superRouter.Group("/localities")
	{
		ll.POST("/", localityController.CreateLocality())
		ll.GET("/reportSellers", localityController.GetAllQtyOfSellers())
	}
}
