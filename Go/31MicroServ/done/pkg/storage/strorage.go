package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	//tables, err := db.Query("SELECT name FROM users.sqlite_master WHERE type='table'")
	log.Printf("Подключение к базе данных: %v", db.Ping())
	//for tables.Next() {
	//	log.Printf(">>%v", tables)
	//}
	return db
}
