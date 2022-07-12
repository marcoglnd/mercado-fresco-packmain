package mariadb

const (
	sqlInsert                     = "INSERT INTO mercado_fresco.buyers (card_number_id, first_name, last_name) VALUES (?, ?, ?);"
	sqlGetAll                     = "SELECT * FROM mercado_fresco.buyers;"
	sqlGetById                    = "SELECT * FROM mercado_fresco.buyers WHERE ID = ?;"
	sqlGetByCardNumberId          = "SELECT * FROM mercado_fresco.buyers WHERE card_number_id = ?;"
	sqlUpdate                     = "UPDATE mercado_fresco.buyers SET card_number_id=?, first_name=?, last_name=? WHERE id=?;"
	sqlDelete                     = "DELETE FROM mercado_fresco.buyers WHERE id=?"
	sqlFindAllPurchaseOrders      = "SELECT b.*, COUNT(p.id) AS `purchase_order_count` FROM mercado_fresco.buyers b INNER JOIN mercado_fresco.purchase_orders p ON b.id = p.buyer_id GROUP BY b.id;"
	sqlFindPurchaseOrderByBuyerId = "SELECT b.*, COUNT(p.id) AS `purchase_order_count` FROM mercado_fresco.buyers b INNER JOIN mercado_fresco.purchase_orders p ON b.id = p.buyer_id WHERE b.id = ? GROUP BY b.id;"
)
