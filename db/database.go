package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

// var (
// 	DBUser   = os.Getenv("DB_USER")
// 	DBPass   = os.Getenv("DB_PASS")
// 	DBServer = os.Getenv("DB_SERVER")
// 	DBPort   = os.Getenv("DB_PORT")
// )

func init() {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatal("failed to load .env")
	// }

	// dataSource := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/mercado_fresco?parseTime=true",
	// 	DBUser,
	// 	DBPass,
	// 	DBServer,
	// 	DBPort,
	// )
	dataSource := "root:secret@tcp(localhost:3306)/mercado_fresco?parseTime=true"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")
}
