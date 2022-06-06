package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections"
)

func sectionsRouter(superRouter *gin.RouterGroup) {
	//1. repositório
	repository := sections.NewRepository()

	//2. serviço (regra de negócio)
	service := sections.NewService(repository)

	//3. controller
	sectionController := controllers.NewSection(service)

	pr := superRouter.Group("/sections")
	{
		pr.GET("/debug", func(ctx *gin.Context) {
			ctx.JSON(http.StatusTeapot, gin.H{
				"debug": "is running",
			})
		})
		pr.GET("/", sectionController.GetAll())
		pr.GET("/:id", sectionController.GetById())
		pr.POST("/", sectionController.Create())
		pr.PATCH("/:id", sectionController.Update())
		pr.DELETE("/:id", sectionController.Delete())
	}
}
