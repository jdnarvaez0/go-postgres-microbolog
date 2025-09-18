package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources
func (s *Server) Start() error {
	return nil
}

// Start the server
func (serv *Server) Close() error {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())

	return nil
}
