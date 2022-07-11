package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRouter *gin.RouterGroup, dbConnection *sql.DB) {
	buyersRouter(superRouter, dbConnection)
	purchaseOrdersRouter(superRouter, dbConnection)
	productsRouter(superRouter, dbConnection)
	employeesRouter(superRouter, dbConnection)
	sectionsRouter(superRouter)
	warehousesRouter(superRouter, dbConnection)
	sellersRouter(superRouter)
}
