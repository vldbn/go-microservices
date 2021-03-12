package server

import (
	"address/internal/app/core"
	"address/internal/app/server/broker"
	"address/internal/app/service"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

// Server
type Server struct {
	router   *chi.Mux
	services service.Services
	ch       *amqp.Channel
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
		r.Route("/api/addresses", func(r chi.Router) {
			r.Get("/", s.HandlerAddressGetList)
		})
		r.Route("/api/addresses/{addressID}", func(r chi.Router) {
			r.Get("/", s.HandlerAddressGetByID)
			r.Patch("/", s.HandlerAddressUpdate)
		})
	})
}

func (s *Server) consume() {
	msgCreate, err := broker.MakeMessageChannel(s.ch, "create")
	if err != nil {
		log.Println(err)
	}

	msgUpdate, err := broker.MakeMessageChannel(s.ch, "update")
	if err != nil {
		log.Println(err)
	}

	msgDelete, err := broker.MakeMessageChannel(s.ch, "delete")
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case msgCreate := <-msgCreate:
			s.services.Address().Create(msgCreate.Body)
		case msgUpd := <-msgUpdate:
			s.services.Address().UpdateUsername(msgUpd.Body)
		case msgDel := <-msgDelete:
			s.services.Address().Delete(msgDel.Body)
		}

	}
}

// NewServer server constructor
func NewServer(services service.Services, ch *amqp.Channel) *Server {
	return &Server{
		router:   chi.NewRouter(),
		services: services,
		ch:       ch,
	}
}

// StartServer starts server
func StartServer(config *core.Config, services service.Services, ch *amqp.Channel) error {
	log.Println("Starting user service at port ", config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	srv := NewServer(services, ch)

	go srv.consume()

	srv.router.Use(middleware.Logger)
	srv.configRoutes()
	return http.ListenAndServe(port, srv)
}
