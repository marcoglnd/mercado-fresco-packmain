package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBConnection *sql.DB

func GetDBConnection() *sql.DB {
	if DBConnection != nil {
		return DBConnection
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(localhost:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	DBConnection, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("failed to connect to mariadb")
	}
	return DBConnection
}
