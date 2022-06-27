package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

var (
	DBUser   = os.Getenv("DB_USER")
	DBPass   = os.Getenv("DB_PASS")
	DBServer = os.Getenv("DB_SERVER")
	DBPort   = os.Getenv("DB_PORT")
	DBName   = os.Getenv("DB_NAME")
)

func init() {
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)%s?parseTime=true",
		DBUser,
		DBPass,
		DBServer,
		DBPort,
		DBName,
	)
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
