package main

//go:generate go-bindata-assetfs -o assets.go assets/...

import (
	"log"
	"net/http"
  "os"
)

var (
	port   string = ":4444"
	server *http.Server
)

func main() {
	server = &http.Server{Addr: port}
  hs, _ := os.Hostname()

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.FileServer(assetFS()))

	http.HandleFunc("/home", home)
	http.HandleFunc("/create", create)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/setClock", setClock)
	log.Println("listening at http://localhost" + port + "/home")
	log.Println("listening at http://" + hs + port + "/home")
	log.Fatal(server.ListenAndServe())
}

func exit(w http.ResponseWriter, r *http.Request) {
	server.Shutdown(nil)
}
