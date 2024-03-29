package mariadb

const (
	sqlInsert = `INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) 
	VALUES (?, ?, ?, ?, ?)`

	sqlGetAll = `SELECT 
		id, 
		order_date, 
		order_number, 
		employee_id, 
		product_batch_id, 
		warehouse_id 
	FROM inbound_orders`
)
