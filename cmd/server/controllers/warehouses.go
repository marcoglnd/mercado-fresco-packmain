package controllers

import "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses"

type WarehouseController struct {
	service warehouses.Service
}

func NewWarehouse(w warehouses.Service) *WarehouseController {
	return &WarehouseController{service: w}
}
