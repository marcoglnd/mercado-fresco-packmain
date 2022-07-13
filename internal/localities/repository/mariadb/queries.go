package mariadb

const (
	sqlCreateLocality  = "INSERT INTO localities (locality_name, province_id) VALUES(?, ?);"
	sqlGetLocalityById = `SELECT localities.id, localities.locality_name, provinces.province_name, countries.country_name 
	FROM countries countries, localities localities, provinces provinces
	WHERE 
		provinces.id = localities.province_id 
		AND countries.id = provinces.id_country_fk
		AND localities.id = ?;`
	sqlGetQtyOfSellersByLocalityId = `SELECT localities.id, localities.locality_name , COUNT(localities.id) sellers_count 
	FROM localities localities, sellers sellers
	WHERE 
		localities.id = sellers.locality_id 
		AND localities.id = ?
			GROUP BY localities.id `
	sqlGetQtyOfSellersByLocality = `SELECT localities.id, localities.locality_name , COUNT(localities.id) sellers_count 
			FROM localities localities, sellers sellers
			WHERE 
				localities.id = sellers.locality_id 
					GROUP BY localities.id `
)
