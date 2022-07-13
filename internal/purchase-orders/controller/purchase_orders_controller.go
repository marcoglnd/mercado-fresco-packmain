package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/domain"
)

type PurchaseOrderController struct {
	purchaseOrder domain.PurchaseOrderService
}

type AppError struct {
	Message string
	Code    int
}

type request struct {
	OrderNumber   string `json:"order_number" binding:"required"`
	OrderDate     string `json:"order_date" binding:"required"`
	TrackingCode  string `json:"tracking_code" binding:"required"`
	BuyerId       int64  `json:"buyer_id" binding:"required"`
	CarrierId     int64  `json:"carrier_id" binding:"required"`
	OrderStatusId int64  `json:"order_status_id" binding:"required"`
	WarehouseId   int64  `json:"warehouse_id" binding:"required"`
}

func NewPurchaseOrderController(purchaseOrder domain.PurchaseOrderService) (*PurchaseOrderController, error) {
	if purchaseOrder == nil {
		return nil, errors.New("invalid purchase Order")
	}

	return &PurchaseOrderController{
		purchaseOrder: purchaseOrder,
	}, nil
}

// @Summary Create purchase order
// @Tags Purchase Orders
// @Description Create a new purchase order
// @Accept json
// @Produce json
// @Success 201 {object} domain.PurchaseOrder
// @Failure 409 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /purchaseOrders [post]
func (c PurchaseOrderController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		purchaseOrder, err := c.purchaseOrder.Create(
			ctx,
			req.OrderNumber,
			req.OrderDate,
			req.TrackingCode,
			req.BuyerId,
			req.CarrierId,
			req.OrderStatusId,
			req.WarehouseId,
		)

		if err != nil {
			if errors.Is(err, domain.ErrDuplicatedOrderNumber) {
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

		ctx.JSON(http.StatusCreated, gin.H{
			"data": purchaseOrder,
		})
	}
}
