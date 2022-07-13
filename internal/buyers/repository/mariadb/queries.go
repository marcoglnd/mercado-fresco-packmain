package mariadb

const (
	sqlInsert                     = "INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?, ?, ?);"
	sqlGetAll                     = "SELECT * FROM buyers;"
	sqlGetById                    = "SELECT * FROM buyers WHERE ID = ?;"
	sqlGetByCardNumberId          = "SELECT * FROM buyers WHERE card_number_id = ?;"
	sqlUpdate                     = "UPDATE buyers SET card_number_id=?, first_name=?, last_name=? WHERE id=?;"
	sqlDelete                     = "DELETE FROM buyers WHERE id=?"
	sqlFindAllPurchaseOrders      = "SELECT b.*, COUNT(p.id) AS `purchase_order_count` FROM buyers b INNER JOIN purchase_orders p ON b.id = p.buyer_id GROUP BY b.id;"
	sqlFindPurchaseOrderByBuyerId = "SELECT b.*, COUNT(p.id) AS `purchase_order_count` FROM buyers b INNER JOIN purchase_orders p ON b.id = p.buyer_id WHERE b.id = ? GROUP BY b.id;"
)
