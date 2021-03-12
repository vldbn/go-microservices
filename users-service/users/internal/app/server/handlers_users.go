package server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"users/internal/app/model/requests"
	"users/internal/app/model/responses"
)

func (s *Server) HandlerUserCreate(w http.ResponseWriter, r *http.Request) {
	req := requests.UserReq{}
	res := responses.UserRes{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	usr, err := s.services.User().Create(req.Username)

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.User = usr
	s.response(w, http.StatusCreated, res)
	return
}

func (s *Server) HandlerUserGetByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	res := responses.UserRes{}

	usr, err := s.services.User().GetByID(userID)

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.User = usr
	s.response(w, http.StatusOK, res)
	return
}

func (s *Server) HandlerUserGetList(w http.ResponseWriter, r *http.Request) {
	res := responses.UsersRes{}

	users, err := s.services.User().GetList()

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.Users = users
	s.response(w, http.StatusOK, res)
	return
}

func (s *Server) HandlerUserUpdate(w http.ResponseWriter, r *http.Request) {
	req := requests.UserReq{}
	res := responses.UserRes{}
	userID := chi.URLParam(r, "userID")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	usr, err := s.services.User().Update(userID, req.Username)

	if err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	res.User = usr
	s.response(w, http.StatusOK, res)
	return
}

func (s *Server) HandlerUserDelete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	res := responses.UserRes{}

	if err := s.services.User().Delete(userID); err != nil {
		res.Error = err.Error()
		s.response(w, http.StatusBadRequest, res)
		return
	}

	s.response(w, http.StatusNoContent, res)
	return
}
