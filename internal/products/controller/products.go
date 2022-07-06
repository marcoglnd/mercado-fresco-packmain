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
// @Success 200 {object} schemes.JSONSuccessResult{data=domain.Product}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /products [get]
func (c *Controller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
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
// @Success 200 {object} domain.Product
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
// @Param product body domain.RequestProductsUpdated true "Product to create"
// @Success 201 {object} domain.Product
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
// @Param product body domain.RequestProductsUpdated true "Product to update"
// @Success 200 {object} domain.Product
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

// @Summary Create product records
// @Tags Products
// @Description Create a new product records
// @Accept json
// @Produce json
// @Param product body domain.RequestProductRecords true "Product record to create"
// @Success 201 {object} domain.ProductRecords
// @Failure 409 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Failure 500 {object} schemes.JSONBadReqResult{error=string}
// @Router /productRecords [post]
func (c *Controller) CreateProductRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProductRecords
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		recordId, err := c.service.CreateProductRecords(
			ctx.Request.Context(),
			&domain.ProductRecords{
				PurchasePrice: req.PurchasePrice,
				SalePrice:     req.SalePrice,
				ProductId:     req.ProductId,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		record, err := c.service.GetProductRecordsById(
			ctx.Request.Context(),
			recordId,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, record)
	}
}

// @Summary Quantity of records
// @Tags Products
// @Description Get quantity of product records
// @Accept json
// @Produce json
// @Param id query int true "records ID"
// @Success 201 {object} domain.ProductRecords
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /products/reportRecords [get]
func (c *Controller) GetQtyOfRecordsById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProductRecordId
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product, err := c.service.GetQtyOfRecordsById(ctx.Request.Context(), req.Id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

// @Summary Create product records
// @Tags Products
// @Description Create a new product records
// @Accept json
// @Produce json
// @Param product body domain.RequestProductBatches true "Product record to create"
// @Success 201 {object} domain.ProductBatches
// @Failure 409 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Failure 500 {object} schemes.JSONBadReqResult{error=string}
// @Router /productRecords [post]
func (c *Controller) CreateProductBatches() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.RequestProductBatches
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid inputs"})
			return
		}
		batchId, err := c.service.CreateProductBatches(
			ctx.Request.Context(),
			&domain.ProductBatches{
				BatchNumber:        req.BatchNumber,
				CurrentQuantity:    req.CurrentQuantity,
				CurrentTemperature: req.CurrentTemperature,
				InitialQuantity:    req.InitialQuantity,
				MinumumTemperature: req.MinumumTemperature,
				ProductId:          req.ProductId,
				SectionId:          req.SectionId,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		batch, err := c.service.GetProductBatchesById(
			ctx.Request.Context(),
			batchId,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, batch)
	}
}
