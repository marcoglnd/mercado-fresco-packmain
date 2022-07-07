package repository

const (
	sqlGetAll = "SELECT * FROM carriers"

	sqlGetById = "SELECT * FROM carriers WHERE id = ?"

	sqlGetByCid = "SELECT * FROM carriers WHERE cid = ?"

	sqlStore = `
	INSERT INTO carriers 
	(
		cid, 
		company_name, 
		address, 
		telephone, 
		locality_id
	) 
	VALUES (?, ?, ?, ?, 1)`

	sqlCarriersCount = `
	SELECT 
		COUNT(c.locality_id) AS carriers_count
	FROM 
		localities l
	LEFT JOIN
        carriers c
	ON  c.locality_id = l.id
    GROUP BY
        l.id
	`
)
