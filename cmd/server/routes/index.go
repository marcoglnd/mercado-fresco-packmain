package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRouter *gin.RouterGroup, dbConnection *sql.DB) {
	productsRouter(superRouter)
	buyersRouter(superRouter)
	employeesRouter(superRouter, dbConnection)
	sectionsRouter(superRouter)
	warehousesRouter(superRouter)
	sellersRouter(superRouter)
}
