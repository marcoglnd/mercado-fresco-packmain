package mariadb

const (
	sqlInsertSection  = "INSERT INTO mercado_fresco.sections (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetAllSections = "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM mercado_fresco.sections"
	sqlGetSectionById = "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM mercado_fresco.sections WHERE id = ?;"
	// sqlUpdateSection  = "UPDATE mercado_fresco.sections SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity`, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE id = ?;"
	// sqlUpdateSection = "UPDATE mercado_fresco.sections SET `section_number`=?, `current_temperature`=?, `minimum_temperature`=?, `current_capacity`=?, `minimum_capacity`=?, `maximum_capacity`=?, `warehouse_id`=?, `product_type_id`=? WHERE ID=?;"
	sqlUpdateSection = "UPDATE mercado_fresco.sections SET `section_number`=?, `current_temperature`=?, `minimum_temperature`=?, `current_capacity`=?, `minimum_capacity`=?, `maximum_capacity`=?, `warehouse_id`=?, `product_type_id`=? WHERE id=?;"
	


	sqlDeleteSection = "DELETE FROM mercado_fresco.sections WHERE id=?"
)
