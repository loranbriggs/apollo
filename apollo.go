package main

//go:generate go-bindata-assetfs -o assets.go assets/...

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
  http.Handle("/", http.FileServer(assetFS()))

	http.HandleFunc("/home", home)
	http.HandleFunc("/createEvent", createEvent)
	http.HandleFunc("/deleteEvent", deleteEvent)
	http.HandleFunc("/updateDuration", updateDuration)
	log.Println("listening at http://localhost:4444/home")
	log.Fatal(http.ListenAndServe(":4444", nil))
}
