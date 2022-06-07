package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
)

type requestEmployee struct {
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
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees [get]
func (c *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// @Summary Employee by id
// @Tags Employees
// @Description get employee by id
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
		employee, err := c.service.GetById(intId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
// @Param employee body requestEmployees true "Employee to create"
// @Success 201 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees [post]
func (c *Employee) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestEmployee
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

		employee, err := c.service.Create(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
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
// @Param employee body requestEmployees true "Employee to update"
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

		var req requestEmployee
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

		employee, err := c.service.Update(int(id), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
