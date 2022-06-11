package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rtnhnd255/osau_biblio_back/app/storage"
)

type Server struct {
	config  *Config
	storage *storage.Storage
	router  *mux.Router
}

func NewServer(cfg *Config, stor *storage.Storage) *Server {
	return &Server{
		config:  cfg,
		storage: stor,
		router:  mux.NewRouter(),
	}
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/", s.homeHandler())
}

func (s *Server) Run() error {
	log.Println("Starting server at ", s.config.Addr, ":", s.config.Port)
	return http.ListenAndServe(s.config.Addr+":"+s.config.Port, s.router)
}
