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
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

// type requestUpdate struct {
// 	SectionNumber      int64 `json:"section_number"`
// 	CurrentTemperature int16 `json:"current_temperature"`
// 	MinimumTemperature int16 `json:"minimum_temperature"`
// 	CurrentCapacity    int64 `json:"current_capacity"`
// 	MinimumCapacity    int64 `json:"minimum_capacity"`
// 	MaximumCapacity    int64 `json:"maximum_capacity"`
// 	WarehouseID        int64 `json:"warehouse_id"`
// 	ProductTypeID      int64 `json:"product_type_id"`
// }

func NewBuyerController(buyer domain.BuyerService) (*BuyerController, error) {

	if buyer == nil {
		return nil, errors.New("invalid buyer")
	}

	return &BuyerController{
		buyer: buyer,
	}, nil
}

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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, buyer)
	}
}

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
