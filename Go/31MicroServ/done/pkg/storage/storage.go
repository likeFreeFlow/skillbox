package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/miekg/pkcs11"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("sqlite3", "../../users.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Подключение к базе данных: %v", db.Ping())
	sts :=
		//DROP TABLE IF EXISTS users;
		`CREATE TABLE users(
		"id" Integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" Text NOT NULL,
		"age" Integer );
	`
	db.Exec(`CREATE TABLE "friends"(
	   		"user1id" Integer NOT NULL,
	   		"user2id" Integer NOT NULL,
	   		CONSTRAINT "friends1" FOREIGN KEY ( "user1id" ) REFERENCES "users"( "id" )
	   			ON DELETE Cascade
	   			ON UPDATE Cascade,
	   		CONSTRAINT "friends2" FOREIGN KEY ( "user2id" ) REFERENCES "users"( "id" )
	   			ON DELETE Cascade
	   			ON UPDATE Cascade
	   	 )`)

	db.Exec(sts)
	return db
}
