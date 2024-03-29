package mariadb

const (
	sqlInsertProduct  = "INSERT INTO products (`description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetAllProducts = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM products;"
	sqlGetProductById = "SELECT `id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id` FROM products WHERE ID = ?;"
	sqlUpdateProduct  = "UPDATE products SET `description` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `height` = ?, `length` = ?, `net_weight` = ?, `product_code` = ?, `recommended_freezing_temperature` = ?, `width` = ?, `product_type_id` = ?, `seller_id` = ? WHERE ID = ?;"
	sqlDeleteProduct  = "DELETE FROM products WHERE id=?"

	sqlCreateRecord = "INSERT INTO `product_records` (`purchase_price`, `sale_price`, `product_id`) VALUES (?, ?, ?);"
	sqlGetRecord    = "SELECT `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM `product_records` WHERE ID = ?;"

	sqlGetQtyOfRecordsById = "SELECT p.id, p.description, COUNT(r.id) records_count FROM products p INNER JOIN product_records r ON p.id = r.product_id WHERE p.id = ? GROUP BY p.id;"
	sqlGetQtyOfRecords     = "SELECT p.id, p.description, COUNT(r.id) records_count FROM products p INNER JOIN product_records r ON p.id = r.product_id GROUP BY p.id;"

	sqlCreateBatch = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	sqlGetBatch = "SELECT `batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id` FROM `product_batches`  WHERE ID=?;"

	sqlGetQtdProductsBySectionId = "SELECT  b.section_id, SUM(b.current_quantity) AS products_count, s.section_number FROM product_batches b INNER JOIN sections s ON b.section_id = s.id WHERE b.section_id = ?;"
	sqlGetQtdProductsInSection = "SELECT  b.section_id, SUM(b.current_quantity) AS products_count, s.section_number	FROM product_batches b INNER JOIN sections s ON b.section_id = s.id GROUP BY b.section_id;"
)
