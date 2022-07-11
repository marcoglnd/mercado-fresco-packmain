package mariadb

const (
	sqlGetAll = `SELECT 
		id,
		card_number_id,
		first_name,
		last_name,
		warehouse_id
	FROM employees`

	sqlFindAllInboundOrders = `SELECT e.*, 
	COUNT(i.id) AS inbound_orders_count 
	FROM employees e 
	INNER JOIN inbound_orders i 
	ON e.id = i.employee_id 
	GROUP BY e.id`

	sqlFindInboundOrdersByEmployeeId = `SELECT e.*, 
	COUNT(p.id) AS inbound_orders_count 
	FROM employees e 
	INNER JOIN inbound_orders i 
	ON e.id = i.employee_id 
	WHERE e.id = ? GROUP BY e.id`

	sqlGetById = `SELECT 
		id,
		card_number_id,
		first_name,
		last_name,
		warehouse_id
	FROM employees
	WHERE ID = ?`

	sqlGetByCardNumberId = `SELECT
		id,
		card_number_id,
		first_name,
		last_name,
		warehouse_id
	FROM employees
	WHERE card_number_id = ?`

	sqlInsert = `INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) 
	VALUES (?, ?, ?, ?)`

	sqlUpdate = `UPDATE employees SET
		card_number_id=?, 
		first_name=?, 
		last_name=?, 
		warehouse_id=?
	WHERE ID =?`

	sqlDelete = `DELETE FROM employees WHERE id=?`
)
