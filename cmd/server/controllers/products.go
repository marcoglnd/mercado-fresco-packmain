package controllers

import (
	"net/http"

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

func NewProduct(p products.Service) *Controller {
	return &Controller{
		service: p,
	}
}