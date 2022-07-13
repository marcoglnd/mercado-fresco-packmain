package domain

type InboundOrder struct {
	ID             int64  `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int64  `json:"employee_id"`
	ProductBatchId int64  `json:"product_batch_id"`
	WarehouseId    int64  `json:"warehouse_id"`
}
