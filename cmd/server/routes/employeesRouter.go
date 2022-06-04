package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
)

func employeesRouter(superRouter *gin.RouterGroup) {
	repository := employees.NewRepository()
	service := employees.NewService(repository)
	e := controllers.NewEmployee(service)

	pr := superRouter.Group("/employees")
	{
		pr.GET("/debug", func(ctx *gin.Context) {
			ctx.JSON(http.StatusTeapot, gin.H{
				"debug": "is running",
			})
		})
		pr.GET("/", e.GetAll()) // Getall
		pr.GET("/:id")          // Get one employee bi id
		pr.POST("/", e.Store()) // Store
		pr.PATCH("/:id", e.UpdateName())
		pr.PUT("/:id", e.Update()) // Update
		pr.DELETE("/:id", e.Delete())
	}
}
