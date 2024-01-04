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
		r.Post("/", h.handlerHello)
		r.Post("/login", h.handleLogin)
		r.Post("/register", h.handleRegister)
	})
}

func (h *Handler) handlerHello(w http.ResponseWriter, r *http.Request) {

	response.ResponseMessage(w, "Hello", nil, http.StatusOK)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var register request.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.Auth.Register(register)

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var login request.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.Auth.Login(login)

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseToken(w, "Success", res.Jwt, http.StatusOK)

}
