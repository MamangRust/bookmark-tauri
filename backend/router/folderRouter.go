package router

import (
	"bookmark-backend/handler"
	"bookmark-backend/service"

	"github.com/go-chi/chi/v5"
)

func NewFolderRouter(prefix string, router *chi.Mux) {
	folderService := service.NewFolderService()

	folderHandler := handler.NewFolderHandler(*folderService)

	router.Route(prefix, func(r chi.Router) {
		r.Get("/", folderHandler.FindAllFolder)
		r.Get("/find/{folder}", folderHandler.FindFolder)
		r.Post("/", folderHandler.CreateFolder)
		r.Put("/{folder}", folderHandler.UpdateFolder)
		r.Delete("/{folder}", folderHandler.DeleteFolder)
	})

}
