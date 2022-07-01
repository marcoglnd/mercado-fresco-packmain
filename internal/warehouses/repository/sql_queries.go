package repository

const (
	sqlGetAll = "SELECT * FROM warehouses"

	sqlStore = `INSERT INTO warehouses 
	(address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) 
	VALUES (?, ?, ?, ?, ?, 1)`

	sqlLastID = "SELECT MAX(id) as last_id FROM warehouses"

	sqlUpdate = `
	UPDATE warehouses SET
	warehouse_code=?, 
	address=?, 
	telephone=?, 
	minimum_capacity=? 
	minimum_temperature=? 
	WHERE id=?
	`

	sqlUpdateWarehouseCode = "UPDATE warehouses SET warehouse_code=? WHERE id=?"

	sqlDelete = "DELETE FROM warehouses WHERE id=?"
)
