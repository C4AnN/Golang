package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	// Format: username:password@tcp(host:port)/dbname
	dsn := "root:@tcp(127.0.0.1:3308)/go_products?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Optional: check if the connection is actually established
	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database connected successfully")
	DB = db
}
