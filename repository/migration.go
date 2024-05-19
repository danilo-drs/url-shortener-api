package repository

import "database/sql"

// ExecMigrations creates the necessary tables in the database
func ExecMigrations(conn *sql.DB) error {

	// Create the short_url table if it does not exist
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS short_url (
		key TEXT PRIMARY KEY, 
		url TEXT, 
		short_url TEXT,
		create_at TEXT, 
		created_by TEXT, 
		update_at TEXT, 
		updated_by TEXT
		)`)
	return err

}
