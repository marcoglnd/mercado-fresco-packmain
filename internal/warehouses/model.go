package warehouses

type Warehouse struct {
	ID                 int    `json:"id"`
	WarehouseCode      string `json:"warehouse_code" binding:"required"`
	Address            string `json:"address" binding:"required"`
	Telephone          string `json:"telephone" binding:"required"`
	MinimumCapacity    int    `json:"minimum_capacity" binding:"required,gte=1"`
	MinimumTemperature int    `json:"minimum_temperature" binding:"required,gte=1"`
}
