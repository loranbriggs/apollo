package main

type Event struct {
	Id     int64
	Time   string
	Action string
	Minutes int // duration of event in minutes
}
