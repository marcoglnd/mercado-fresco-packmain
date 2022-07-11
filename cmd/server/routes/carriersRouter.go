package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/controller"
	repository "github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/service"
)

func carriersRouter(superRouter *gin.RouterGroup, dbConnection *sql.DB) {
	repository := repository.NewCarrierRepository(dbConnection)
	service := service.NewCarrierService(repository)
	carrierController := controller.NewCarrierController(service)

	pr := superRouter.Group("/carriers")
	{
		pr.POST("/", carrierController.Create())
		pr.GET("/reportLocalities", carrierController.ReportCarriers())
	}
}
