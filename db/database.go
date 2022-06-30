package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
)

func InitDB() (*sql.DB) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dataSource := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/mercado_fresco?parseTime=true",
		config.DBUser,
		config.DBPass,
		config.DBServer,
		config.DBPort,
	)
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")
	return conn
}
