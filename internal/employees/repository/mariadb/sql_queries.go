package mariadb

const (
	sqlGetAll = `SELECT * FROM mercado_fresco.employees;`

	sqlGetById = `SELECT * FROM mercado_fresco.employees WHERE ID = ?;`

	sqlInsert = "INSERT INTO mercado_fresco.employees (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?);"

	sqlUpdate = `UPDATE mercado_fresco.employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE ID = ?;`

	sqlDelete = `DELETE FROM mercado_fresco.employees WHERE id = ?;`
)
