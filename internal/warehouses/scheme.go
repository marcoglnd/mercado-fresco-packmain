package warehouses

type UpdateWarehouseInput struct {
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"gte=0"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"gte=0"`
}

type CreateWarehouseInput struct {
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required,gte=1"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"required,gte=1"`
}
