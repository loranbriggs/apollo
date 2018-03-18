package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func dbConn() *sql.DB {
	db, err := sql.Open("sqlite3", "./events.db")
	checkErr(err)
	// create table of only valid actions, these should not be added or removed.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS actions (action TEXT PRIMARY KEY)")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO actions (action) VALUES (?)", "open")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO actions (action) VALUES (?)", "close")
	checkErr(err)
	// create table to store events, actions must be from actions table and times must be unique
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, time TEXT, action TEXT, FOREIGN KEY(action) REFERENCES actions(action), UNIQUE(time))")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO events (time, action) VALUES (?, ?)", "05:30", "open")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO events (time, action) VALUES (?, ?)", "17:30", "close")
	checkErr(err)
	// create table to store durations. These should be updated, not added or removed.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS durations (action TEXT PRIMARY KEY, duration INTEGER, FOREIGN KEY(action) REFERENCES actions(action))")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO durations (action, duration) VALUES (?, ?)", "open", 30)
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO durations (action, duration) VALUES (?, ?)", "close", 30)
	checkErr(err)
	return db
}
