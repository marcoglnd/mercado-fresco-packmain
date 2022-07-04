package mariadb

const (
	sqlInsert = "INSERT INTO mercado_fresco.products (`description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetAll = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM mercado_fresco.products;"
	sqlGetById = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM mercado_fresco.products WHERE ID = ?;"
	sqlUpdate = "UPDATE mercado_fresco.products SET `description` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `height` = ?, `length` = ?, `net_weight` = ?, `product_code` = ?, `recommended_freezing_temperature` = ?, `width` = ?, `product_type_id` = ?, `seller_id` = ? WHERE ID = ?;"
	sqlDelete = "DELETE FROM mercado_fresco.products WHERE id=?"
	sqlCreateRecord = "INSERT INTO `mercado_fresco`.`product_records` (`purchase_price`, `sale_price`, `product_id`) VALUES (?, ?, ?);"
	sqlGetRecord = "SELECT `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM `mercado_fresco`.`product_records` WHERE ID = ?;"
)