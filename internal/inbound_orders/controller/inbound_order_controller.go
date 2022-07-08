package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain"
)

type requestInboundOrderCreate struct {
	OrderDate      string `json:"order_date" binding:"required"`
	OrderNumber    string `json:"order_number" binding:"required"`
	EmployeeId     int64  `json:"employee_id" binding:"required"`
	ProductBatchId int64  `json:"product_batch_id" binding:"required"`
	WarehouseId    int64  `json:"warehouse_id" binding:"required"`
}

type InboundOrderController struct {
	service domain.InboundOrderService
}

func NewInboundOrderController(service domain.InboundOrderService) (*InboundOrderController, error) {
	if service == nil {
		return nil, domain.ErrInvalidService
	}

	return &InboundOrderController{
		service: service,
	}, nil
}

func (i InboundOrderController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestInboundOrderCreate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		inboundOrder, err := i.service.Create(ctx, &domain.InboundOrder{
			OrderDate:      req.OrderDate,
			OrderNumber:    req.OrderNumber,
			EmployeeId:     req.EmployeeId,
			ProductBatchId: req.ProductBatchId,
			WarehouseId:    req.WarehouseId,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, inboundOrder)
	}
}
