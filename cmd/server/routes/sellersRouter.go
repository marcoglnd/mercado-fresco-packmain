package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"
)

func sellersRouter(superRouter *gin.RouterGroup) {

	repo := sellers.NewRepository()
	service := sellers.NewService(repo)
	controller := controllers.NewSeller(service)

	pr := superRouter.Group("/sellers")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.CreateNewSeller())
		//pr.PATCH("/:id", controller.Update())
		//pr.DELETE("/:id", controller.Delete())
		pr.GET("/debug", Debug)
	}
}

func Debug(ctx *gin.Context) {
	ctx.JSON(http.StatusTeapot, gin.H{
		"debug": "is running",
	})
}
