package handler

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initAuthGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Get("/", h.handlerHello)
		r.Post("/login", h.handleLogin)
		r.Post("/register", h.handleRegister)
	})
}
func (h *Handler) handlerHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Auth"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var register request.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := register.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "validation failed: "+err.Error())
		return
	}

	res, err := h.services.Auth.Register(register)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var login request.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := login.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "validation failed: "+err.Error())
		return
	}

	res, err := h.services.Auth.Login(login)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err.Description)
		return
	}

	response.ResponseToken(w, "Success", res.Jwt, http.StatusOK)
}
