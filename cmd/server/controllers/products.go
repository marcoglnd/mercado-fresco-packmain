package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
)

type Controller struct {
	service products.Service
}

func NewProduct(p products.Service) *Controller {
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
		data, err := c.service.GetAll()
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
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		product, err := c.service.GetById(intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

type requestProducts struct {
	Description                    string  `json:"description" binding:"required"`
	ExpirationRate                 int     `json:"expiration_rate" binding:"required"`
	FreezingRate                   int     `json:"freezing_rate" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"netweight" binding:"required"`
	ProductCode                    string  `json:"product_code" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id" binding:"required"`
}

// @Summary Create product
// @Tags Products
// @Description Add a new product to the list
// @Accept json
// @Produce json
// @Param product body requestProducts true "Product to create"
// @Success 201 {object} products.Product
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /products [post]
func (c *Controller) CreateNewProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestProducts
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		product, err := c.service.CreateNewProduct(
			req.Description,
			req.ExpirationRate,
			req.FreezingRate,
			req.Height,
			req.Length,
			req.NetWeight,
			req.ProductCode,
			req.RecommendedFreezingTemperature,
			req.Width,
			req.ProductTypeId,
			req.SellerId)
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
// @Router /products/{id} [patch]
func (c *Controller) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestProducts
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		product, err := c.service.Update(intId,
			req.Description, req.ExpirationRate, req.FreezingRate,
			req.Height, req.Length, req.NetWeight, req.ProductCode,
			req.RecommendedFreezingTemperature, req.Width, req.ProductTypeId, req.SellerId)
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
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("product %d removed", id)})
	}
}
