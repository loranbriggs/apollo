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
    rows.Scan(&event.Id, &event.Time, &event.Action)
    events = append(events, event)
  }

  rows, err = db.Query("SELECT * FROM durations")
  checkErr(err)
  durations := make([]Duration, 0)
  for rows.Next() {
    duration := Duration{}
    rows.Scan(&duration.Action, &duration.Duration)
    durations = append(durations, duration)
  }

  go scheduleActions(events, durations)

  a, err := Asset("assets/home.html")
  checkErr(err)
  t, _ := template.New("home").Parse(string(a))

  data := struct {
    Events []Event
    Durations []Duration
  } {
    events,
    durations,
  }

  t.Execute(w, &data)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  defer db.Close()

  if r.Method != http.MethodPost {
    http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    return
  }

  event := Event{}
  event.Time = r.FormValue("time")
  event.Action = r.FormValue("action")
  mountReadWrite()
  _, err := db.Exec("INSERT INTO events (time, action) VALUES (?, ?)", event.Time, event.Action)
  checkErr(err)
  mountReadOnly()
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  defer db.Close()

  id := r.URL.Query()["id"][0]
  if id == "" {
    http.Error(w, "Please send ID", http.StatusBadRequest)
  }
  mountReadWrite()
  _, err := db.Exec("DELETE FROM events WHERE id = ?", id)
  checkErr(err)
  mountReadOnly()
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateDuration(w http.ResponseWriter, r *http.Request) {
  db := dbConn()
  defer db.Close()

  dur := Duration{}
  dur.Action = r.FormValue("action")
  dur.Duration, _ = strconv.Atoi(r.FormValue("duration"))
  mountReadWrite()
  _, err := db.Exec("UPDATE durations SET duration = ? WHERE action = ? ", dur.Duration, dur.Action)
  checkErr(err)
  mountReadOnly()
  http.Redirect(w, r, "/", http.StatusSeeOther)
}
