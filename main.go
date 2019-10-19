package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := s.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	ws.WriteMessage(websocket.TextMessage, []byte("Welcome, sock"))
}

type server struct {
	websocket.Upgrader
}

func main() {
	var s server

	r := chi.NewRouter()
	r.Get("/", s.homeHandler)
	r.Get("/ws", s.wsHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
