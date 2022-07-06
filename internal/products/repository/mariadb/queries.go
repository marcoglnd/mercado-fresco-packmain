package mariadb

const (
	sqlInsertProduct       = "INSERT INTO mercado_fresco.products (`description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetAllProducts      = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM mercado_fresco.products;"
	sqlGetProductById      = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM mercado_fresco.products WHERE ID = ?;"
	sqlUpdateProduct       = "UPDATE mercado_fresco.products SET `description` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `height` = ?, `length` = ?, `net_weight` = ?, `product_code` = ?, `recommended_freezing_temperature` = ?, `width` = ?, `product_type_id` = ?, `seller_id` = ? WHERE ID = ?;"
	sqlDeleteProduct       = "DELETE FROM mercado_fresco.products WHERE id=?"
	
	sqlCreateRecord        = "INSERT INTO `mercado_fresco`.`product_records` (`purchase_price`, `sale_price`, `product_id`) VALUES (?, ?, ?);"
	sqlGetRecord           = "SELECT `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM `mercado_fresco`.`product_records` WHERE ID = ?;"
	
	sqlGetQtyOfRecordsById = "SELECT p.id, p.description, COUNT(r.id) records_count FROM mercado_fresco.products p INNER JOIN product_records r ON p.id = r.product_id WHERE p.id = ? GROUP BY p.id;"

	sqlCreateBatch        = "INSERT INTO `mercado_fresco`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `initial_quantity`, `minumum_temperature`, `product_id`, `section_id`) VALUES (?, ?, ?);"
	sqlGetBatch           = "SELECT batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id FROM mercado_fresco.product_batches;"
)
