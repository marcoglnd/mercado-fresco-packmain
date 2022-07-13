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

// @Summary Create locality
// @Tags Localities
// @Description Add a new Locality to the list
// @Accept json
// @Produce json
// @Param Locality body requestCreateLocality true "locality to create"
// @Success 201 {object} schemas.JSONSuccessResult{data=domain.GetLocality}
// @Failure 409 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /localities [post]
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

		locality, err := c.service.GetLocalityByID(ctx, localId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"data": locality})
	}
}

// @Summary List of reports AllQtyOfSellers
// @Tags Localities
// @Description get AllQtyOfSellers
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JSONSuccessResult{data=domain.QtyOfSellers}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /localities/reportSellers [get]
func (c LocalityController) GetAllQtyOfSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		strId := ctx.Query("id")
		intId, _ := strconv.ParseInt(strId, 10, 64)
		if intId == 0 {
			listsOfSellers, err := c.service.GetAllQtyOfSellers(ctx)
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
		ctx.JSON(http.StatusOK, gin.H{"data": sellersByLocality})
	}
}
