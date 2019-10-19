package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New connection")

	ws, err := s.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		ws.Close()
		fmt.Println("Connection closed")
	}()

	ws.WriteMessage(websocket.TextMessage, []byte("Welcome, sock"))
}

type server struct {
	websocket.Upgrader
}

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("PORT not set")
	}

	var s server

	r := chi.NewRouter()
	r.Get("/", s.homeHandler)
	r.Get("/ws", s.wsHandler)

	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
