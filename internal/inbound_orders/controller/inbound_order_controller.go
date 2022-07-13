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

// @Summary List Inbound Orders
// @Tags Inbound Orders
// @Description get all inbound orders
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JSONSuccessResult{data=domain.InboundOrder}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /inboundOrders [get]
func (c InboundOrderController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		inboundOrders, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": inboundOrders})
	}
}

// @Summary Create inbound order
// @Tags Inbound Orders
// @Description Add a new inbound order to the list
// @Accept json
// @Produce json
// @Param inbound order body requestInboundOrderCreate true "Inbound Order to create"
// @Success 201 {object} domain.InboundOrder
// @Failure 409 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Router /inboundOrders [post]
func (c InboundOrderController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestInboundOrderCreate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		inboundOrder, err := c.service.Create(ctx, &domain.InboundOrder{
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

		ctx.JSON(http.StatusCreated, gin.H{"data": inboundOrder})
	}
}
