package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
)

type Controller struct {
	service products.Service
}

func (c *Controller) GetAll() gin.HandlerFunc {
	data, err := c.service.GetAll()
	return func(ctx *gin.Context) {
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"data": data,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func (c *Controller) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		p, err := c.service.GetById(intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": p})
	}
}

func NewProduct(p products.Service) *Controller {
	return &Controller{
		service: p,
	}
}
