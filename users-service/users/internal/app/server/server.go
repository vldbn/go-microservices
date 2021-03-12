package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"users/internal/app/core"
	"users/internal/app/service"
)

type Server struct {
	router   *chi.Mux
	services service.Services
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) response(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			return
		}
	}
}

func (s *Server) configRoutes() {
	s.router.Group(func(r chi.Router) {
		r.Route("/api/users", func(r chi.Router) {
			r.Get("/", s.HandlerUserGetList)
			r.Post("/", s.HandlerUserCreate)
		})

		r.Route("/api/users/{userID}", func(r chi.Router) {
			r.Get("/", s.HandlerUserGetByID)
			r.Patch("/", s.HandlerUserUpdate)
			r.Delete("/", s.HandlerUserDelete)
		})
	})
}

func NewServer(services service.Services) *Server {
	return &Server{
		router:   chi.NewRouter(),
		services: services,
	}
}

func StartServer(config *core.Config, services service.Services) error {
	log.Println("Starting user service at port ", config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	srv := NewServer(services)
	srv.router.Use(middleware.Logger)
	srv.configRoutes()
	return http.ListenAndServe(port, srv)
}
