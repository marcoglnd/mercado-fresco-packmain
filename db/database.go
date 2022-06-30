package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
)

var (
	StorageDB *sql.DB
)

func init() {
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
	// dataSource := "root:secret@tcp(localhost:3306)/mercado_fresco?parseTime=true"
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")
}
