package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/routes"
	"github.com/marcoglnd/mercado-fresco-packmain/db"
	"github.com/marcoglnd/mercado-fresco-packmain/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MERCADO FRESCOS
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// BasePath /api/v1
// @query.collection.format multi

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbConnection := db.GetDBConnection()
	PATH := "/api/v1"
	router := gin.Default()
	routerGroup := router.Group(PATH)
	routes.AddRoutes(routerGroup, dbConnection)
	docs.SwaggerInfo.BasePath = PATH

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
