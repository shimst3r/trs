// Package db implements database-related utility and maintenance functions.
package db

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"time"

	// this package needs a blank import
	_ "github.com/mattn/go-sqlite3"
)

// TimeEntry models a row in the time_entry table.
type TimeEntry struct {
	ID         int64
	Start, End time.Time
}

// DoesNotExist checks whether the database file exists.
func DoesNotExist() (bool, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return false, err
	}
	_, err = os.Stat(dbPath)
	return os.IsNotExist(err), nil
}

// InitDB ensures that the trs database exists and contains all necessary schemas.
func InitDB() error {
	dbPath, err := getDBPath()
	if err != nil {
		return err
	}
	createDBFile(dbPath)
	db, err := openDB(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	err = doMigrations(db)
	if err != nil {
		return err
	}
	return nil
}

// createDBFile checks whether the trs database file exists and creates it if not.
func createDBFile(path string) error {
	notExists, err := DoesNotExist()
	if err != nil {
		return err
	}
	if notExists {
		os.Create(path)
		fmt.Printf("trs.db created at %s\n", path)
	}
	return nil
}

// getDBPath returns the path to .trs.db based on the user's home directory.
func getDBPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir + "/.trs.db", nil
}

// openDB opens the database at the expected path.
func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}
