// Package db implements database-related utility and maintenance functions.
package db

import "database/sql"

// doMigrations creates the time_entry table and a corresponding indices on its primary key and the start/end columns.
func doMigrations(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS time_entry(
	"id" INTEGER PRIMARY KEY,
	"start" INTEGER,
	"end" INTEGER
);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	query = "CREATE UNIQUE INDEX IF NOT EXISTS idx_time_entry_id ON time_entry(id);"
	stmt, err = db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	query = "CREATE UNIQUE INDEX IF NOT EXISTS idx_time_entry_start_asc_end_asc ON time_entry(start ASC, end ASC);"
	stmt, err = db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
