package handler

import (
	"bookmark-backend/internal/middlewares"
	"bookmark-backend/internal/service"
	"bookmark-backend/pkg/auth"
	"bookmark-backend/pkg/upload"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	services     *service.Service
	images       upload.ImageUpload
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Service, tokenManager auth.TokenManager, images upload.ImageUpload) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		images:       images,
	}
}

func (h *Handler) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(middlewares.MiddlewareCors)

	fileServer := http.FileServer(http.Dir("./static"))

	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	h.InitApi(router)

	return router

}

func (h *Handler) InitApi(r *chi.Mux) {
	r.Get("/", h.handlerHello)

	h.initAuthGroup("/auth", r)
	h.initCategoryGroup("/category", r)
	h.initPostsGroup("/posts", r)
}
