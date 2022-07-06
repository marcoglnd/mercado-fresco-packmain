package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
)

// Controller receives a service
type SellerController struct {
	service domain.SellerService
}

func NewSellerController(service domain.SellerService) (*SellerController, error) {

	if service == nil {
		return nil, errors.New("invalid service")
	}

	return &SellerController{
		service: service,
	}, nil
}

// @Summary List sellers
// @Tags Sellers
// @Description get all sellers
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Seller}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sellers [get]
func (c SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sellers, err := c.service.GetAll(ctx.Request.Context())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, sellers)
	}
}

// @Summary Seller by id
// @Tags Sellers
// @Description get Seller by it's id
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param token header string true "token"
// @Success 200 {object} schemes.Seller
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sellers/{id} [get]
func (c SellerController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		strId := ctx.Param("id")
		intId, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		seller, err := c.service.GetByID(ctx, intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, seller)
	}
}

type requestCreate struct {
	Cid          int64  `json:"cid" binding:"required"`
	Company_name string `json:"company_name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Telephone    string `json:"telephone" binding:"required"`
	LocalityID   int64  `json:"locality_id" binding:"required"`
}

// @Summary Create seller
// @Tags Sellers
// @Description Add a new Seller to the list
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param Seller body requestSellers true "seller to create"
// @Success 201 {object} schemes.Seller
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /sellers [post]
func (c SellerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestCreate

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		seller, err := c.service.Create(ctx, &domain.Seller{
			Cid:          req.Cid,
			Company_name: req.Company_name,
			Address:      req.Address,
			Telephone:    req.Telephone,
			LocalityID:   req.LocalityID,
		})
		if err != nil {
			if errors.Is(err, domain.ErrDuplicatedCID) {
				ctx.JSON(http.StatusConflict, gin.H{
					"message": err.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
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
		if req.Telephone == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O telefone da empresa é obrigatório"})
			return
		}
		if req.LocalityID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "O id da localidade é obrigatório"})
			return
		}

		ctx.JSON(http.StatusCreated, seller)
	}
}

type requestUpdate struct {
	Cid          int64  `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	LocalityID   int64  `json:"locality_id"`
}

// @Summary Update seller
// @Tags Sellers
// @Description Update existing Seller in list
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param token header string true "token"
// @Param seller body requestSellers true "Seller to update"
// @Success 200 {object} schemes.Seller
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sellers/{id} [patch]
func (c *SellerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strId := ctx.Param("id")
		intId, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var req requestUpdate

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		seller, err := c.service.Update(ctx, &domain.Seller{
			ID:           intId,
			Cid:          req.Cid,
			Company_name: req.Company_name,
			Address:      req.Address,
			Telephone:    req.Telephone,
			LocalityID:   req.LocalityID,
		})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, seller)
	}
}

// @Summary Delete seller
// @Tags Sellers
// @Description Delete existing seller in list
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param token header string true "token"
// @Success 204 {object} schemes.JSONSuccessResult{data=string}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /sellers/{id} [delete]
func (c *SellerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strId := ctx.Param("id")
		intId, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = c.service.Delete(ctx, intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
