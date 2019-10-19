package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/go-chi/chi"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, sock!")
}

func main() {

	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/ws", wsHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
