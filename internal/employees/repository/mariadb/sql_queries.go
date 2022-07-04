package mariadb

const (
	sqlGetAll = `SELECT * FROM employees`

	sqlGetById = `SELECT * FROM employees WHERE ID = ?`

	sqlInsert = `INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)`

	sqlUpdate = `UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE ID = ?`

	sqlDelete = `DELETE FROM employees WHERE id=?`
)
