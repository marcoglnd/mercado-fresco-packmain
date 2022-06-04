package warehouses

type UpdateWarehouseInput struct {
	WarehouseCode      string `json:"warehouse_code" binding:"len=3"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimumCapacity    int    `json:"minimum_capacity" binding:"gte=1"`
	MinimumTemperature int    `json:"minimum_temperature" binding:"gte=1"`
}

type CreateWarehouseInput struct {
	WarehouseCode      string `json:"warehouse_code" binding:"required,len=3"`
	Address            string `json:"address" binding:"required"`
	Telephone          string `json:"telephone" binding:"required"`
	MinimumCapacity    int    `json:"minimum_capacity" binding:"required,gte=1"`
	MinimumTemperature int    `json:"minimum_temperature" binding:"required,gte=1"`
}
