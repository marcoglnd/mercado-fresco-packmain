package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
)

type Controller struct {
	service domain.Service
}

func NewProduct(p domain.Service) *Controller {
	return &Controller{
		service: p,
	}
}

// @Summary List products
// @Tags Products
// @Description get all products
// @Accept json
// @Produce json
// @Success 200 {object} schemes.JSONSuccessResult{data=products.Product}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /products [get]
func (c *Controller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// @Summary Product by id
// @Tags Products
// @Description get product by it's id
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} products.Product
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /products/{id} [get]



func (c *Controller) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProductId
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		product, err := c.service.GetById(ctx.Request.Context(), req.Id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

// @Summary Create product
// @Tags Products
// @Description Add a new product to the list
// @Accept json
// @Produce json
// @Param product body requestProducts true "Product to create"
// @Success 201 {object} products.Product
// @Failure 409 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /products [post]
func (c *Controller) CreateNewProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProducts
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		product, err := c.service.CreateNewProduct(
			ctx.Request.Context(),
			&domain.Product{
				Description:                    req.Description,
				ExpirationRate:                 req.ExpirationRate,
				FreezingRate:                   req.FreezingRate,
				Height:                         req.Height,
				Length:                         req.Length,
				NetWeight:                      req.NetWeight,
				ProductCode:                    req.ProductCode,
				RecommendedFreezingTemperature: req.RecommendedFreezingTemperature,
				Width:                          req.Width,
				ProductTypeId:                  req.ProductTypeId,
				SellerId:                       req.SellerId,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, product)
	}
}

// @Summary Update product
// @Tags Products
// @Description Update existing product in list
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body requestProducts true "Product to update"
// @Success 200 {object} products.Product
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /products/{id} [patch]

func (c *Controller) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProductsUpdated
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		var reqId domain.RequestProductId
		if err := ctx.ShouldBindUri(&reqId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		product, err := c.service.Update(
			ctx.Request.Context(),
			&domain.Product{
				Id:                             reqId.Id,
				Description:                    req.Description,
				ExpirationRate:                 req.ExpirationRate,
				FreezingRate:                   req.FreezingRate,
				Height:                         req.Height,
				Length:                         req.Length,
				NetWeight:                      req.NetWeight,
				ProductCode:                    req.ProductCode,
				RecommendedFreezingTemperature: req.RecommendedFreezingTemperature,
				Width:                          req.Width,
				ProductTypeId:                  req.ProductTypeId,
				SellerId:                       req.SellerId,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

// @Summary Delete product
// @Tags Products
// @Description Delete existing product in list
// @Accept json
// @Produce json
// @Param id path int true "product ID"
// @Success 204 {object} schemes.JSONSuccessResult{data=string}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /products/{id} [delete]
func (c *Controller) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProductId
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err := c.service.Delete(ctx.Request.Context(), req.Id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("product %d removed", req.Id)})
	}
}
