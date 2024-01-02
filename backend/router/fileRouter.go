package router

import (
	"bookmark-backend/handler"
	"bookmark-backend/service"

	"github.com/go-chi/chi/v5"
)

func NewFileRouter(prefix string, router *chi.Mux) {
	fileService := service.NewFileService()
	fileHandler := handler.NewFileHandler(*fileService)

	router.Route(prefix, func(r chi.Router) {
		r.Post("/create", fileHandler.CreateFile)
		r.Post("/find", fileHandler.FindFile)
		r.Put("/update", fileHandler.UpdateFile)
		r.Delete("/delete", fileHandler.DeleteFile)
	})
}
