package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func warehousesRouter(superRouter *gin.RouterGroup) {
	pr := superRouter.Group("/warehouses")
	{
		pr.GET("/debug", func(ctx *gin.Context) {
			ctx.JSON(http.StatusTeapot, gin.H{
				"debug": "is running",
			})
		})
	}
}
