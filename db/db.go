package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var DB *sql.DB
var err error

func InitDB() {
	_ = godotenv.Load()
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "api.db"
	}
	DB, err = sql.Open("sqlite3", dbPath)

	if err != nil {
		panic("Could not connect to DataBase.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not Create Users Table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
        user_id INTEGER,
        FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not Create Event Table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not Create Registrations Table")
	}
}
