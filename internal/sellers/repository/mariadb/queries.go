package mariadb

const (
	sqlInsertSeller  = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES(?, ?, ?, ?, ?);"
	sqlGetAllSellers = "SELECT * FROM sellers;"
	sqlGetSellerById = "SELECT * FROM sellers WHERE ID = ?;"
	sqlUpdateSeller  = "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?;"
	sqlDeleteSeller  = "DELETE FROM sellers WHERE id=?"
)
