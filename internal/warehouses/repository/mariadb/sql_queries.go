package repository

const (
	sqlGetAll = "SELECT * FROM warehouses"

	sqlGetById = "SELECT * FROM warehouses WHERE id = ?"

	sqlGetByWarehouseCode = "SELECT * FROM warehouses WHERE warehouse_code = ?"

	sqlStore = `
	INSERT INTO warehouses 
	(
		address, 
		telephone, 
		warehouse_code, 
		minimum_capacity, 
		minimum_temperature, 
		locality_id
	) 
	VALUES (?, ?, ?, ?, ?, 1)`

	sqlUpdate = `
	UPDATE warehouses SET
		warehouse_code=?, 
		address=?, 
		telephone=?, 
		minimum_capacity=?, 
		minimum_temperature=? 
	WHERE id=?
	`

	sqlDelete = "DELETE FROM warehouses WHERE id=?"
)
