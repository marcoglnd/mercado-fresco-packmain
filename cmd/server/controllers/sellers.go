package controllers

import (
	"fmt"
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
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		s, err := c.service.GetById(intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, s)
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
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Cid == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "A identificação da empresa (cid) é obrigatória"})
			return
		}
		if req.Company_name == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O nome da empresa é obrigatório"})
			return
		}
		if req.Address == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O endereço da empresa é obrigatório"})
			return
		}
		if req.Telephone == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O telefone da empresa é obrigatório"})
			return
		}

		s, err := c.service.Store(req.Cid, req.Company_name, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, s)
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
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "A identificação da empresa (cid) é obrigatória"})
			return
		}
		if req.Company_name == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O nome da empresa é obrigatório"})
			return
		}
		if req.Address == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O endereço da empresa é obrigatório"})
			return
		}
		if req.Telephone == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O telefone da empresa é obrigatório"})
			return
		}

		p, err := c.service.Update(int(id), req.Cid, req.Company_name, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *SellerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("O seller %d foi removido", id)})
	}
}
