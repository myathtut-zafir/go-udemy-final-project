package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, _ = sql.Open("sqlite3", "api.db")
	// if _ != nil {
	// 	panic("DB Error!")
	// }
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL 
	);`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		print("Error creating table")
	}

}
