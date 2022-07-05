package mariadb

const (
	sqlInsert = "INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)"

	sqlGetAll = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees"

	sqlGetById = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?"

	sqlUpdate = "UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=? WHERE id=?"

	sqlDelete = "DELETE FROM employees WHERE id=?"
)
