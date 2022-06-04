package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
)

type request struct {
	ID           int    `json:"id"`
	CardNumberId int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

type Employee struct {
	service employees.Service
}

func NewEmployee(e employees.Service) *Employee {
	return &Employee{
		service: e,
	}
}

func (c *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)
	}
}

func (c *Employee) GetEmployee() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		e, err := c.service.GetEmployee(intId)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)
	}
}

func (c *Employee) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		e, err := c.service.Create(req.ID, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)
	}
}

func (c *Employee) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		e, err := c.service.Update(req.ID, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)
	}
}

func (c *Employee) Delete() gin.HandlerFunc {
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
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("The employee %d was deleted", id)})
	}
}
