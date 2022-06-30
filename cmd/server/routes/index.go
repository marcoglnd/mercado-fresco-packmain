package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRouter *gin.RouterGroup, conn *sql.DB) {
	productsRouter(superRouter, conn)
	buyersRouter(superRouter)
	employeesRouter(superRouter)
	sectionsRouter(superRouter)
	warehousesRouter(superRouter)
	sellersRouter(superRouter)
}
