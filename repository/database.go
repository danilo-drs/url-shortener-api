package repository

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

// Connect to the database
func Connect() error {
	// Get the database connection parameters from the environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	connString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname

	if sslMode != "" {
		connString += " sslmode=" + sslMode
	}

	// Connect to the database

	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}
	// Set the database singleton instance
	DB = conn

	// Create tables if they do not exist
	err = ExecMigrations(conn)
	if err != nil {
		return err
	}

	// Defer the connection close to the end of the program
	// defer conn.Close()

	// Return nil if the connection is successful
	return nil
}
