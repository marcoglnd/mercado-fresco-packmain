package routes

/*
import (
	"log"
	"net/http"
	"strconv"

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
		// lista todos os vendedores existentes
		pr.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"data": &sellers})
		})

		// retorna as informacoes do vendedor de acordo com o id informado (param)
		pr.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			parsedID, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "erro interno, tente novamente",
				})

				log.Println(err.Error())

				return
			}

			if parsedID != 10 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "id não encontrado",
				})

				log.Println("id não encontrado")

				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": id,
			})

		})

		// insercao de dados
		pr.POST("/", func(ctx *gin.Context) {

			// validação de token -> header
			token := ctx.GetHeader("token")
			if token != "123456" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "token inválido",
				})
			}

			var seller Seller
			// tratando erro caso os dados não estiverem compatíveis com a estrutura esperada
			if err := ctx.ShouldBindJSON(&seller); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// incrementando o ID a cada post
			lastID++
			seller.ID = lastID

			// adicionando na lista de sellers
			sellers = append(sellers, seller)

			//		if seller.Company_name != "frutaria" {
			//		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "não autorizado"})
			//		return
			//}

			ctx.JSON(http.StatusOK, gin.H{
				"data": seller,
			})
		})
	}
}
*/
