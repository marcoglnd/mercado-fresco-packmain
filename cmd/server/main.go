package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/routes"
)

func main() {
	router := gin.Default()
	routerGroup := router.Group("/api/v1")
	routes.AddRoutes(routerGroup)

	router.Run()
}
