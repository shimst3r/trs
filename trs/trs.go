// Package trs implements the trs state machine.
package trs

import (
	"fmt"
	"os"
	"time"

	"github.com/shimst3r/trs/db"
)

// Help prints the help and usage message.
func Help() {
	fmt.Println(`trs is a terminal-based time-recording system.
		
Usage:
	trs <command>
	
The commands are:

	help	print this help message
	init	initialise the trs database (stored at $HOME/.trs.db)
	start	start a time entry
	stop	stop the currently running time entry
	today	print the amount of time worked today`)
}

// Init sets up the database file and schemas.
func Init() {
	notExist, err := db.DoesNotExist()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !notExist {
		fmt.Println("error: database file already exists")
		os.Exit(1)
	}
	db.InitDB()
}

// Start adds a new time entry to the database.
func Start() {
	notExist, err := db.DoesNotExist()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if notExist {
		fmt.Println("error: database file does not exist. Run \"trs init\" before starting a time entry.")
		os.Exit(1)
	}
	err = db.StartRow(time.Now())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("success: Started new time entry")
}

// Stop sets the "end" field on the latest non-ended time entry.
func Stop() {
	notExist, err := db.DoesNotExist()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if notExist {
		fmt.Println("error: database file does not exist. Run \"trs init\" before using trs.")
		os.Exit(1)
	}
	err = db.StopRow(time.Now())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("success: Stopped time entry")
}

// Today prints the time worked at the current day.
func Today() {
	notExist, err := db.DoesNotExist()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if notExist {
		fmt.Println("error: database file does not exist. Run \"trs init\" before using trs.")
		os.Exit(1)
	}
	midnight, tomorrow := getStartAndEndTimes(time.Now())
	entries, err := db.GetTimeEntries(midnight, tomorrow)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var duration time.Duration
	for _, entry := range entries {
		duration += entry.End.Sub(entry.Start)
	}
	fmt.Printf("You have been working %v today.\n", duration)
}

func getStartAndEndTimes(timestamp time.Time) (time.Time, time.Time) {
	start := time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.FixedZone(timestamp.Zone()))
	end := start.AddDate(0, 0, 1)

	return start, end
}
