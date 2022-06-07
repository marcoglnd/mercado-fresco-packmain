package controllers

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

// @Summary List buyers
// @Tags Buyers
// @Description get all buyers
// @Accept json
// @Produce json
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Buyer}
// @Failure 404 {object} schemes.JSONBadReqResult{}
// @Router /buyers [get]
func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": b,
		})
	}
}

// @Summary Buyer by id
// @Tags Buyers
// @Description get buyer by its id
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Buyer}
// @Failure 400 {object} schemes.JSONBadReqResult{}
// @Failure 404 {object} schemes.JSONBadReqResult{}
// @Router /buyers/{id} [get]
func (c *BuyerController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		b, err := c.service.GetById(buyerId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}

// @Summary Create buyer
// @Tags Buyers
// @Description Add a new buyer to the list
// @Accept json
// @Produce json
// @Param buyer body request true "Buyer to create"
// @Success 201 {object} schemes.JSONSuccessResult{data=schemes.Buyer}
// @Failure 404 {object} schemes.JSONBadReqResult{}
// @Failure 422 {object} schemes.JSONBadReqResult{}
// @Router /buyers/ [post]
func (c *BuyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.CardNumberID == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O CardNumberID do buyer é obrigatório"})
			return
		}
		if req.FirstName == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O FirstName do buyer é obrigatório"})
			return
		}
		if req.LastName == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O LastName do buyer é obrigatório"})
			return
		}
		b, err := c.service.Create(req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, b)
	}
}

// @Summary Update buyer
// @Tags Buyers
// @Description Update existing buyer in list
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body request true "Buyer to update"
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Buyer,}
// @Failure 400 {object} schemes.JSONBadReqResult{}
// @Failure 404 {object} schemes.JSONBadReqResult{}
// @Router /buyers/{id} [patch]
func (c *BuyerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
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

// @Summary Delete buyer
// @Tags Buyers
// @Description Delete existing buyer in list
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 204 {object} schemes.JSONSuccessResult{data=schemes.Buyer,}
// @Failure 400 {object} schemes.JSONBadReqResult{}
// @Failure 404 {object} schemes.JSONBadReqResult{}
// @Router /buyers/{id} [delete]
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

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("O buyer %d foi removido", id)})
	}
}

type request struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
