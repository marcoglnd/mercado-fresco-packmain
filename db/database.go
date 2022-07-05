package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var dbConnection *sql.DB

func GetDBConnection() *sql.DB {
	if dbConnection != nil {
		return dbConnection
	}
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(localhost:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	dbConnection, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("failed to connect to mariadb")
	}
	return dbConnection
}
