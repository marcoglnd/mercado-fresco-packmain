package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"
)

// Controller recebe um serviço
type SellerController struct {
	service sellers.Service
}

func NewSeller(s sellers.Service) *SellerController {
	return &SellerController{
		service: s,
	}
}

func (c *SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		s, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, s)
	}
}

func (c *SellerController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
	}
}

type request struct {
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    int    `json:"telephone"`
}

func (c *SellerController) CreateNewSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		s, err := c.service.Store(req.Cid, req.Company_name, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, s)
	}
}

func (c *SellerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Cid == 0 {
			ctx.JSON(400, gin.H{"error": "A identificação da empresa é obrigatória"})
			return
		}
		if req.Company_name == "" {
			ctx.JSON(400, gin.H{"error": "O nome da empresa é obrigatório"})
			return
		}
		if req.Address == "" {
			ctx.JSON(400, gin.H{"error": "O endereço é obrigatório"})
			return
		}
		if req.Telephone == 0 {
			ctx.JSON(400, gin.H{"error": "O telefone é obrigatório"})
			return
		}

		p, err := c.service.Update(int(id), req.Cid, req.Company_name, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *SellerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": fmt.Sprintf("O seller %d foi removido", id)})
	}
}

/*
type Sellers []Seller

var lastID int
var sellers Sellers
*/

/*
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
*/
