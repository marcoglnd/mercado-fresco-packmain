package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/service"
)

func employeesRouter(superRouter *gin.RouterGroup, conn *sql.DB) {
	repository := mariadb.NewMariaDBRepository(conn)
	service := service.NewEmployeeService(repository)
	controller, _ := controller.NewEmployeeController(service)

	pr := superRouter.Group("/employees")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.Create())
		pr.PATCH("/:id", controller.Update())
		pr.DELETE("/:id", controller.Delete())
		pr.GET("/reportInboundOrders", controller.ReportInboundOrders())
	}
}
