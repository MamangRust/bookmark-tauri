package handler

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/middlewares"
	"bookmark-backend/pkg/auth"
	"bookmark-backend/pkg/slugify"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initCategoryGroup(prefix string, r *chi.Mux) {

	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)

		r.Get("/hello", h.handlerCategoryHello)

		r.Get("/", h.handlerCategoryAll)
		r.Get("/{id}", h.handlerCategoryById)
		r.Post("/create", h.handlerCategoryCreate)
		r.Put("/update/{id}", h.handlerCategoryUpdate)
		r.Delete("/delete/{id}", h.handlerCategoryDelete)
	})
}

func (h *Handler) handlerCategoryHello(w http.ResponseWriter, r *http.Request) {

	userIDValueRaw := auth.GetContextUserId(r)

	fmt.Println("userIDValueRaw: ", userIDValueRaw)

	response.ResponseMessage(w, "Success", userIDValueRaw, http.StatusOK)

}

func (h *Handler) handlerCategoryAll(w http.ResponseWriter, r *http.Request) {

	res, err := h.services.Category.GetAll()

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}

func (h *Handler) handlerCategoryById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	res, byidErr := h.services.Category.GetByID(int(Id))

	if byidErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, byidErr.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}

func (h *Handler) handlerCategoryCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2 * 1024 * 1024)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "file too large")
		return
	}

	var category request.CreateCategoryRequest

	slug := slugify.Slugify(r.FormValue("name"))

	category.Name = slug
	category.Description = r.FormValue("description")

	filePath := h.images.HandleImageUpload(slug, w, r)

	if filePath == "" {
		response.ResponseError(w, http.StatusBadRequest, "invalid file")
		return
	}

	category.Image = filePath

	if err := category.Validate(); err != nil {
		response.HandleValidationErrors(w, err)

		return
	}

	res, err_create := h.services.Category.Create(category)

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err_create.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}

func (h *Handler) handlerCategoryUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2 * 1024 * 1024)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "file too large")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var category request.UpdateCategoryRequest

	category.CategoryID = int(id)
	slug := slugify.Slugify(r.FormValue("name"))

	category.Name = slug
	category.Description = r.FormValue("description")

	filePath := h.images.HandleImageUpload(slug, w, r)

	if filePath == "" {
		response.ResponseError(w, http.StatusBadRequest, "invalid file")
		return
	}

	category.Image = filePath

	if err := category.Validate(); err != nil {
		response.HandleValidationErrors(w, err)

		return
	}

	fmt.Println("category: ", category)

	res, updateErr := h.services.Category.Update(category)
	if updateErr != nil {
		response.ResponseError(w, http.StatusInternalServerError, updateErr.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)

}

func (h *Handler) handlerCategoryDelete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	res, deleteErr := h.services.Category.Delete(int(id))

	if deleteErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, deleteErr.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}
