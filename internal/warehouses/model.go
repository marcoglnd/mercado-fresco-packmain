package warehouses

type Warehouse struct {
	ID                 int     `json:"id" gorm:"primaryKey" binding:"omitempty"`
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required,gte=1"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"required,gte=1"`
	LocalityId         int     `json:"locality_id"`
}
