package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
)

type BuyerController struct {
	buyer domain.BuyerService
}

type AppError struct {
	Message string
	Code    int
}

type request struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

func NewBuyerController(buyer domain.BuyerService) (*BuyerController, error) {

	if buyer == nil {
		return nil, errors.New("invalid buyer")
	}

	return &BuyerController{
		buyer: buyer,
	}, nil
}

// @Summary List buyers
// @Tags Buyers
// @Description get all buyers
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JSONSuccessResult{data=domain.Buyer}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /buyers [get]
func (c BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyers, err := c.buyer.GetAll(ctx.Request.Context())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, buyers)
	}
}

// @Summary Buyer by id
// @Tags Buyers
// @Description get buyer by it's id
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} domain.Buyer
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /buyer/{id} [get]
func (c BuyerController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		buyer, err := c.buyer.GetById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, buyer)
	}
}

// @Summary Create buyer
// @Tags Buyer
// @Description Add a new buyer to the list
// @Accept json
// @Produce json
// @Param buyer body domain.RequestBuyer true "Buyer to create"
// @Success 201 {object} domain.Buyer
// @Failure 409 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Router /buyers [post]
func (c BuyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		buyer, err := c.buyer.Create(ctx, req.CardNumberID, req.FirstName, req.LastName)

		if err != nil {
			if errors.Is(err, domain.ErrDuplicatedID) {
				ctx.JSON(http.StatusConflict, gin.H{
					"message": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, buyer)
	}
}

// @Summary Update buyer
// @Tags Buyers
// @Description Update existing buyer in list
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body domain.RequestBuyer true "Buyer to update"
// @Success 200 {object} domain.Buyer
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Router /buyers/{id} [patch]
func (c BuyerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		buyer, err := c.buyer.Update(ctx, id, req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, buyer)
	}
}

// @Summary Delete buyer
// @Tags Buyers
// @Description Delete existing buyer in list
// @Accept json
// @Produce json
// @Param id path int true "buyer ID"
// @Success 204 {object} schemas.JSONSuccessResult{data=string}
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /buyers/{id} [delete]
func (c BuyerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = c.buyer.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}

// @Summary Report purchase orders
// @Tags Buyers
// @Description Get quantity of purchase orders for buyer
// @Accept json
// @Produce json
// @Param id query int true "buyer ID"
// @Success 201 {object} domain.PurchaseOrdersResponse
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /buyers/reportPurchaseOrders [get]
func (c BuyerController) ReportPurchaseOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)

		if id == 0 {
			report, err := c.buyer.ReportAllPurchaseOrders(ctx)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data": report,
			})
			return
		}

		report, err := c.buyer.ReportPurchaseOrders(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": report,
		})
	}
}
