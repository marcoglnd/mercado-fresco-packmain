package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRouter *gin.RouterGroup) {
	productsRouter(superRouter)
	buyersRouter(superRouter)
	employeesRouter(superRouter)
	sectionsRouter(superRouter)
	warehousesRouter(superRouter)
	sellersRouter(superRouter)
}
