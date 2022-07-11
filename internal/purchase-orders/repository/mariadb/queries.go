package mariadb

const (
	sqlInsert           = "INSERT INTO mercado_fresco.purchase_orders (order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?);"
	sqlGetByOrderNumber = "SELECT * FROM mercado_fresco.purchase_orders WHERE order_number = ?;"
)
