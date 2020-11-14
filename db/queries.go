// Package db implements database-related utility and maintenance functions.
package db

import (
	"database/sql"
	"errors"
	"time"
)

// GetTimeEntries returns all time_entry rows within start and end timestamps.
func GetTimeEntries(start, end time.Time) ([]TimeEntry, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}
	db, err := openDB(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	entries, err := getTimeEntries(start.Unix(), end.Unix(), db)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// StartRow inserts a new row into time_entry
func StartRow(timestamp time.Time) error {
	dbPath, err := getDBPath()
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	// first check if there's currently no active time entry
	exists, err := existsOpenTimeEntry(db)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("error: Found unfinished time entry. Run \"trs stop\" before starting a new time entry")
	}
	err = createNewTimeEntry(timestamp, db)
	if err != nil {
		return err
	}
	return nil
}

// StopRow sets the end field of an open time entry row if it exists.
func StopRow(timestamp time.Time) error {
	dbPath, err := getDBPath()
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	// first check if there's currently an active time entry
	exists, err := existsOpenTimeEntry(db)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("error: Found no unfinished time entry. Run \"trs start\" to start a new time entry")
	}
	err = setEndOnLatestTimeEntry(timestamp, db)
	if err != nil {
		return err
	}
	return nil
}

// createNewTimeEntry adds a new row to the table time_entry.
func createNewTimeEntry(timestamp time.Time, db *sql.DB) error {
	query := "INSERT INTO time_entry (start) VALUES (?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(timestamp.Unix())
	if err != nil {
		return err
	}
	return nil
}

// existsOpenTimeEntry checks if there is already an open time entry in the database.
func existsOpenTimeEntry(db *sql.DB) (bool, error) {
	query := "SELECT id FROM time_entry WHERE end IS NULL"
	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

// getTimeEntries returns all rows from time_entry within two timestamps.
func getTimeEntries(start, end int64, db *sql.DB) ([]TimeEntry, error) {
	query := `SELECT id, start, end
	FROM time_entry
	WHERE start >= ?
	AND end < ?
	AND end IS NOT NULL`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []TimeEntry
	for rows.Next() {
		var id, start, end int64
		rows.Scan(&id, &start, &end)
		timeEntry := TimeEntry{id, time.Unix(start, 0), time.Unix(end, 0)}
		result = append(result, timeEntry)
	}
	return result, nil
}

// setEndOnLatestTimeEntry adds a new row to the table time_entry.
func setEndOnLatestTimeEntry(timestamp time.Time, db *sql.DB) error {
	query := "UPDATE time_entry SET end = ? WHERE end IS NULL"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(timestamp.Unix())
	if err != nil {
		return err
	}
	return nil
}
