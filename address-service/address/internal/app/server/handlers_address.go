package server

import (
	"address/internal/app/model/requests"
	"address/internal/app/model/responses"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Server) HandlerAddressGetList(w http.ResponseWriter, r *http.Request) {
	res := responses.AddressesRes{}

	addresses, err := s.services.Address().GetList()

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.Addresses = addresses
	s.response(w, http.StatusOK, res)
	return
}

func (s *Server) HandlerAddressGetByID(w http.ResponseWriter, r *http.Request) {
	res := responses.AddressRes{}
	id := chi.URLParam(r, "addressID")

	adr, err := s.services.Address().GetByID(id)

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.Address = adr
	s.response(w, http.StatusOK, res)
	return
}

func (s *Server) HandlerAddressUpdate(w http.ResponseWriter, r *http.Request) {
	res := responses.AddressRes{}
	req := requests.AddressReq{}
	id := chi.URLParam(r, "addressID")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	adr, err := s.services.Address().Update(id, req.Country, req.City, req.Address)

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.Address = adr
	s.response(w, http.StatusOK, res)
	return
}
