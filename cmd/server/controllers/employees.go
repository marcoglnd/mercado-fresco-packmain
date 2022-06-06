package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
)

type request struct {
	ID           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

type Employee struct {
	service employees.Service
}

func NewEmployee(e employees.Service) *Employee {
	return &Employee{
		service: e,
	}
}

// @Summary List employees
// @Tags Employees
// @Description get all employees
// @Accept json
// @Produce json
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 404 {object} schemes.JSONSuccessResult{error=string}
// @Router /employees [get]
func (c *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": e})
	}
}

// @Summary Employee by id
// @Tags Employees
// @Description get employee by it's id
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees/{id} [get]
func (c *Employee) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		e, err := c.service.GetById(intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, e)
	}
}

// @Summary Create employees
// @Tags Employees
// @Description Add a new employee to the list
// @Accept json
// @Produce json
// @Param employee body request true "Employee to create"
// @Success 201 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees/{id} [post]
func (c *Employee) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.CardNumberId == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "CardNumberId is required"})
			return
		}
		if req.FirstName == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "FirstName is required"})
			return
		}
		if req.LastName == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "LastName is required"})
			return
		}
		if req.WarehouseId < 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "WarehouseId cannot be negative"})
			return
		}

		e, err := c.service.Create(req.ID, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, e)
	}
}

// @Summary Update employee
// @Tags Employees
// @Description Update existing employee in list
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body request true "Employee to update"
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees/{id} [patch]
func (c *Employee) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.CardNumberId == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "CardNumberId is required"})
			return
		}
		if req.FirstName == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "FirstName is required"})
			return
		}
		if req.LastName == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "LastName is required"})
			return
		}
		if req.WarehouseId < 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "WarehouseId cannot be negative"})
			return
		}

		e, err := c.service.Update(int(id), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, e)
	}
}

// @Summary Delete employee
// @Tags Employees
// @Description Delete existing employee in list
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 204 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees/{id} [delete]
func (c *Employee) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("Employee %d was deleted", id)})
	}
}
