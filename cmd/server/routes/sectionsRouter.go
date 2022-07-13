package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/service"
)

func sectionsRouter(superRouter *gin.RouterGroup, DBConnection *sql.DB) {
	repository := mariadb.NewMariaDBRepository(DBConnection)
	service := service.NewService(repository)
	sectionController := controller.NewSection(service)

	pr := superRouter.Group("/sections")
	{
		pr.GET("/", sectionController.GetAll())
		pr.GET("/:id", sectionController.GetById())
		pr.POST("/", sectionController.Create())
		pr.PATCH("/:id", sectionController.Update())
		pr.DELETE("/:id", sectionController.Delete())
	}
}
