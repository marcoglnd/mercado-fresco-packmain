package buyers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"

	"github.com/gin-gonic/gin"
)

type BuyerController struct {
	service buyers.Service
}

func NewBuyer(b buyers.Service) *BuyerController {
	return &BuyerController{
		service: b,
	}
}

func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}

func (c *BuyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		b, err := c.service.Create(req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}

func (c *BuyerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if req.CardNumberID == "" {
			ctx.JSON(400, gin.H{"error": "O CardNumberID do buyer é obrigatório"})
			return
		}
		if req.FirstName == "" {
			ctx.JSON(400, gin.H{"error": "O FirstName do buyer é obrigatório"})
			return
		}
		if req.LastName == "" {
			ctx.JSON(400, gin.H{"error": "O LastName do buyer é obrigatório"})
			return
		}

		b, err := c.service.Update(int(id), req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}

func (c *BuyerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O buyer %d foi removido", id)})
	}
}

type request struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
