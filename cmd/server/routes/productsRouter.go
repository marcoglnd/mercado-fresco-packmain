package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func productsRouter(superRouter *gin.RouterGroup) {
	pr := superRouter.Group("/products")
	{
		pr.GET("/debug", func(ctx *gin.Context) {
			ctx.JSON(http.StatusTeapot, gin.H{
				"debug": "is running",
			})
		})
	}
}
