package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func dbConn() *sql.DB {
	db, err := sql.Open("sqlite3", "./events.db")
	checkErr(err)
	// create table of valid actions.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS actions (action TEXT PRIMARY KEY)")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO actions (action) VALUES (?)", "open")
	checkErr(err)
	_, err = db.Exec("INSERT OR IGNORE INTO actions (action) VALUES (?)", "close")
	checkErr(err)
	// create table to store events, actions must be from actions table and times must be unique
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, time TEXT, action TEXT, minutes INTEGER, FOREIGN KEY(action) REFERENCES actions(action), UNIQUE(time))")
	checkErr(err)
	//_, err = db.Exec("INSERT OR IGNORE INTO events (time, action, minutes) VALUES (?, ?, ?)", "05:30", "open", 5)
	//checkErr(err)
	//_, err = db.Exec("INSERT OR IGNORE INTO events (time, action, minutes) VALUES (?, ?, ?)", "17:30", "close", 7)
	//checkErr(err)
	return db
}
