package domain

type Employee struct {
	ID           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int64  `json:"warehouse_id"`
}

type InboundOrderResponse struct {
	ID                 int64  `json:"id"`
	CardNumberId       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WarehouseId        int64  `json:"warehouse_id"`
	InboundOrdersCount int64  `json:"inbound_orders_count"`
}
