package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/k-tsurumaki/todo-app/config"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	DB, err = sql.Open(config.Config.SLQDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	
	DB.Exec(cmdU)
}