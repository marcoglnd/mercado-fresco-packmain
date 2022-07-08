package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/service"
)

func inboundOrderRouter(superRouter *gin.RouterGroup, conn *sql.DB) {
	repository := mariadb.NewMariaDBRepository(conn)
	service := service.NewInboundOrderService(repository)
	controller, _ := controller.NewInboundOrderController(service)

	pr := superRouter.Group("/inboundOrders")
	{
		pr.POST("/", controller.Create())
		//pr.GET("/reportInboundOrders", controller.GetOrdersByEmployeeId())
	}
}
