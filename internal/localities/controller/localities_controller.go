package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
)

// Controller receives a service
type LocalityController struct {
	service domain.LocalityService
}

func NewLocalityController(service domain.LocalityService) (*LocalityController, error) {

	if service == nil {
		return nil, errors.New("invalid service")
	}

	return &LocalityController{
		service: service,
	}, nil
}

type requestCreateLocality struct {
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceID   int64  `json:"province_id" binding:"required"`
}

func (c *LocalityController) CreateLocality() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestCreateLocality
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		localId, err := c.service.CreateLocality(
			ctx.Request.Context(),
			&domain.Locality{
				LocalityName: req.LocalityName,
				ProvinceID:   req.ProvinceID,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }

		locality, err := c.service.GetLocalityByID(ctx, localId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, locality)
	}
}

func (c LocalityController) GetQtyOfSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		strId := ctx.Query("id")
		intId, _ := strconv.ParseInt(strId, 10, 64)
		if intId == 0 {
			listsOfSellers, err := c.service.GetQtyOfSellers(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, listsOfSellers)
			return
		}

		sellersByLocality, err := c.service.GetQtyOfSellersByLocalityId(ctx, intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, sellersByLocality)
	}
}
