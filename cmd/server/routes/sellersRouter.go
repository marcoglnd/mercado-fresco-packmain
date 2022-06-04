package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Seller struct {
	ID           int    `json:"id"`
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Adress       string `json:"adress"`
	Telephone    int    `json:"telephone"`
}

type Sellers []Seller

var lastID int
var sellers Sellers

func sellersRouter(superRouter *gin.RouterGroup) {
	pr := superRouter.Group("/sellers")
	{
		pr.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"data": &sellers})
		})

		pr.GET("/id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "is running",
			})
		})

		pr.POST("/", func(ctx *gin.Context) {

			var seller Seller
			// tratando erro caso os dados não estiverem compatíveis com a estrutura esperada
			if err := ctx.ShouldBindJSON(&seller); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			lastID++
			seller.ID = lastID

			sellers = append(sellers, seller)

			/*
				if seller.Company_name != "frutaria" {
					ctx.JSON(http.StatusUnauthorized, gin.H{"status": "não autorizado"})
					return
				}
			*/
			ctx.JSON(http.StatusOK, gin.H{
				"data": seller,
			})
		})
	}
}
