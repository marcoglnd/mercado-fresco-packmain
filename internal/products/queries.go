package products

const (
	sqlStore = "INSERT INTO mercado_fresco.products (`description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetAll = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM mercado_fresco.products;"
	sqlGetById = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM mercado_fresco.products WHERE ID = ?;"
	sqlUpdate = "UPDATE mercado_fresco.products SET `description` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `height` = ?, `length` = ?, `net_weight` = ?, `product_code` = ?, `recommended_freezing_temperature` = ?, `width` = ?, `product_type_id` = ?, `seller_id` = ? WHERE ID = ?;"
)