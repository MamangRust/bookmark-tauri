package handler

import (
	"bookmark-backend/internal/middlewares"
	"bookmark-backend/internal/service"
	"bookmark-backend/pkg/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	services *service.Service

	tokenManager auth.TokenManager
}

func NewHandler(services *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(middlewares.MiddlewareCors)

	h.InitApi(router)

	return router

}

func (h *Handler) InitApi(r *chi.Mux) {
	h.initAuthGroup("/auth", r)
	h.initCategoryGroup("/category", r)
	h.initPostsGroup("/posts", r)
}
