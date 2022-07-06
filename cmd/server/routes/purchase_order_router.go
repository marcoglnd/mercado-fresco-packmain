package routes

import (
	"database/sql"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/controller"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/repository/mariadb"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/service"

	"github.com/gin-gonic/gin"
)

func purchaseOrdersRouter(superRouter *gin.RouterGroup, DBConnection *sql.DB) {
	repository := mariadb.NewMariaDBRepository(DBConnection)

	purchaseOrderService := service.NewPurchaseOrderService(repository)

	purchaseOrderController, _ := controller.NewPurchaseOrderController(purchaseOrderService)
	pr := superRouter.Group("/purchaseOrders")
	{
		pr.POST("/", purchaseOrderController.Create())
	}
}
