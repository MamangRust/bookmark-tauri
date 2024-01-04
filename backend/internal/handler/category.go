package handler

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initCategoryGroup(prefix string, r *chi.Mux) {

	r.Route(prefix, func(r chi.Router) {
		r.Get("/hello", h.handlerCategoryHello)
		r.Get("/", h.handlerCategoryAll)
		r.Get("/{id}", h.handlerCategoryById)
		r.Post("/create", h.handlerCategoryCreate)
		r.Put("/update/{id}", h.handlerCategoryUpdate)
		r.Delete("/delete/{id}", h.handlerCategoryDelete)
	})
}

func (h *Handler) handlerCategoryHello(w http.ResponseWriter, r *http.Request) {

	response.ResponseMessage(w, "Hello", nil, http.StatusOK)
}

func (h *Handler) handlerCategoryAll(w http.ResponseWriter, r *http.Request) {

	res, err := h.services.Category.GetAll()

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}

func (h *Handler) handlerCategoryById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, byidErr := h.services.Category.GetByID(int(Id))

	if byidErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, byidErr)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}

func (h *Handler) handlerCategoryCreate(w http.ResponseWriter, r *http.Request) {
	var category request.CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.Category.Create(category)

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}

func (h *Handler) handlerCategoryUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	var category request.UpdateCategoryRequest

	category.CategoryID = int(id)

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, updateErr := h.services.Category.Update(category)
	if updateErr != nil { // Check updateErr for an error
		response.ResponseError(w, http.StatusInternalServerError, updateErr)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)

}

func (h *Handler) handlerCategoryDelete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, deleteErr := h.services.Category.Delete(int(id))

	if deleteErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, deleteErr)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}
