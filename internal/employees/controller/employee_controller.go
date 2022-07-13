package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"
)

type requestEmployeeCreate struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseId  int64  `json:"warehouse_id" binding:"required"`
}

type requestEmployeeUpdate struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int64  `json:"warehouse_id"`
}

type EmployeeController struct {
	service domain.EmployeeService
}

func NewEmployeeController(service domain.EmployeeService) (*EmployeeController, error) {
	if service == nil {
		return nil, domain.ErrInvalidService
	}

	return &EmployeeController{
		service: service,
	}, nil
}

// @Summary List employees
// @Tags Employees
// @Description get all employees
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JSONSuccessResult{data=domain.Employee}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /employees [get]
func (c EmployeeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		employees, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": employees})
	}
}

// @Summary Employee by id
// @Tags Employees
// @Description get employee by id
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} domain.Employee
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /employees/{id} [get]
func (c EmployeeController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		employee, err := c.service.GetById(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, employee)
	}
}

// @Summary Create employee
// @Tags Employees
// @Description Add a new employee to the list
// @Accept json
// @Produce json
// @Param employee body requestEmployeeCreate true "Employee to create"
// @Success 201 {object} domain.Employee
// @Failure 409 {object} schemas.JSONBadReqResult{error=string}
// @Failure 422 {object} schemas.JSONBadReqResult{error=string}
// @Router /employees [post]
func (c EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestEmployeeCreate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		employee, err := c.service.Create(ctx, &domain.Employee{
			CardNumberId: req.CardNumberId,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			WarehouseId:  req.WarehouseId,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, employee)
	}
}

// @Summary Update employee
// @Tags Employees
// @Description Update existing employee in list
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body requestEmployeeUpdate true "Employee to update"
// @Success 200 {object} domain.Employee
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /employees/{id} [patch]
func (c EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req requestEmployeeUpdate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		employee, err := c.service.Update(ctx, &domain.Employee{
			ID:           id,
			CardNumberId: req.CardNumberId,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			WarehouseId:  req.WarehouseId,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, employee)
	}
}

// @Summary Delete employee
// @Tags Employees
// @Description Delete existing employee in list
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 204 {object} schemas.JSONSuccessResult{data=string}
// @Failure 400 {object} schemas.JSONBadReqResult{error=string}
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Router /employees/{id} [delete]
func (c EmployeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}

// @Summary Quantity of inbound orders
// @Tags Employees
// @Description Get quantity of inbound orders
// @Accept json
// @Produce json
// @Param id query int true "employee ID"
// @Success 200 {object} domain.InboundOrderResponse
// @Failure 404 {object} schemas.JSONBadReqResult{error=string}
// @Failure 500 {object} schemas.JSONBadReqResult{error=string}
// @Router /employees/reportInboundOrders [get]
func (c EmployeeController) ReportInboundOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)

		if id == 0 {
			inboundOrders, err := c.service.ReportAllInboundOrders(ctx)

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"data": inboundOrders})
			return
		}

		inboundOrder, err := c.service.ReportInboundOrders(ctx, id)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": inboundOrder})
	}
}
