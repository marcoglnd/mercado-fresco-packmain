package mariadb

const (
	sqlInsertSeller  = "INSERT INTO mercado_fresco.sellers (cid, company_name, address, telephone) VALUES (?, ?, ?, ?);"
	sqlGetAllSellers = "SELECT * FROM mercado_fresco.sellers;"
	sqlGetSellerById = "SELECT * FROM mercado_fresco.sellers WHERE ID = ?;"
	sqlUpdateSeller  = "UPDATE mercado_fresco.sellers SET cid=?, company_name=?, address=?, telephone=? WHERE id=?;"
	sqlDeleteSeller  = "DELETE FROM mercado_fresco.sellers WHERE id=?"
)
