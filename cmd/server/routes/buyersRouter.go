package routes

import (
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"

	"github.com/gin-gonic/gin"
)

func buyersRouter(superRouter *gin.RouterGroup) {
	//1. repositório
	repository := buyers.NewRepository()

	//2. serviço (regra de negócio)
	service := buyers.NewService(repository)

	//3. controller
	buyerController := controllers.NewBuyer(service)
	pr := superRouter.Group("/buyers")
	{
		pr.GET("/", buyerController.GetAll())
		pr.POST("/", buyerController.Create())
		pr.PATCH("/:id", buyerController.Update())
		pr.DELETE("/:id", buyerController.Delete())
	}
}
