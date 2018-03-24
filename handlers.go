package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM events ORDER BY time")
	checkErr(err)
	events := make([]Event, 0)
	for rows.Next() {
		event := Event{}
		rows.Scan(&event.Id, &event.Time, &event.Action, &event.Minutes)
		events = append(events, event)
	}

	go scheduleActions(events)

	a, err := Asset("assets/home.html")
	checkErr(err)
	t, _ := template.New("home").Parse(string(a))

	data := struct {
		Events []Event
	}{
		events,
	}

	t.Execute(w, &data)
}

func create(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

  event := Event{}
  var err error
	event.Time = r.FormValue("time")
	event.Action = r.FormValue("action")
  event.Minutes, err = strconv.Atoi(r.FormValue("minutes"))
	_, err = db.Exec("INSERT INTO events (time, action, minutes) VALUES (?, ?, ?)", event.Time, event.Action, event.Minutes)
	checkErr(err)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	id := r.URL.Query()["id"][0]
	if id == "" {
		http.Error(w, "Please send ID", http.StatusBadRequest)
	}
	_, err := db.Exec("DELETE FROM events WHERE id = ?", id)
	checkErr(err)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func setClock(w http.ResponseWriter, r *http.Request) {
  t := r.FormValue("time")
  d := r.FormValue("date")

  setDate(t, d)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
