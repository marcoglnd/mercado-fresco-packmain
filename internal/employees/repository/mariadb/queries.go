package mariadb

const (
	sqlGetAll = `SELECT 
		id,
		card_number_id,
		first_name,
		last_name,
		warehouse_id
	FROM employees`

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
