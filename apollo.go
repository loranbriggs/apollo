package main

//go:generate go-bindata -o assets.go assets/

import (
  "log"
	"net/http"
)

func main() {
  http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", home)
	http.HandleFunc("/createEvent", createEvent)
	http.HandleFunc("/deleteEvent", deleteEvent)
	http.HandleFunc("/updateDuration", updateDuration)
	log.Println("listening at http://localhost:4444/")
	log.Fatal(http.ListenAndServe(":4444", nil))
}
