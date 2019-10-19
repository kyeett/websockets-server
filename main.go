package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

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
