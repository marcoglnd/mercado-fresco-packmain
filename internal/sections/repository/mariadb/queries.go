package mariadb

const (
	sqlInsertSection  = "INSERT INTO sections (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetAllSections = "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM sections"
	sqlGetSectionById = "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM sections WHERE id = ?;"
	sqlUpdateSection  = "UPDATE sections SET `section_number`=?, `current_temperature`=?, `minimum_temperature`=?, `current_capacity`=?, `minimum_capacity`=?, `maximum_capacity`=?, `warehouse_id`=?, `product_type_id`=? WHERE id=?;"
	sqlDeleteSection  = "DELETE FROM sections WHERE id=?"
)
