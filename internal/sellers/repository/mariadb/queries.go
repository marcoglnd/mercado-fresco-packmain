package mariadb

const (
	sqlInsertSeller    = "INSERT INTO mercado_fresco.sellers (cid, company_name, address, telephone, locality_id) VALUES(?, ?, ?, ?, ?);"
	sqlGetAllSellers   = "SELECT * FROM mercado_fresco.sellers;"
	sqlGetSellerById   = "SELECT * FROM mercado_fresco.sellers WHERE ID = ?;"
	sqlUpdateSeller    = "UPDATE mercado_fresco.sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?;"
	sqlDeleteSeller    = "DELETE FROM mercado_fresco.sellers WHERE id=?"
	sqlCreateLocality  = "INSERT INTO mercado_fresco.localities (locality_name, province_id) VALUES(?, ?);"
	sqlGetLocalityById = `SELECT localities.id, localities.locality_name, provinces.province_name, countries.country_name 
	FROM mercado_fresco.countries countries, mercado_fresco.localities localities, mercado_fresco.provinces provinces
	WHERE 
		provinces.id = localities.province_id 
		AND countries.id = provinces.id_country_fk
		AND localities.id = ?;`
)
