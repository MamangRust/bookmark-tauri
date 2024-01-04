package handler

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initPostsGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Get("/hello", h.handlerPostsHello)
		r.Get("/", h.handlerPostsAll)
		r.Get("/{id}", h.handlerPostsById)
		r.Post("/create", h.handlerPostsCreate)
		r.Put("/update/{id}", h.handlerPostsUpdate)
		r.Delete("/delete/{id}", h.handlerPostsDelete)
	})
}

func (h *Handler) handlerPostsHello(w http.ResponseWriter, r *http.Request) {
	response.ResponseMessage(w, "Hello", nil, http.StatusOK)
}

func (h *Handler) handlerPostsAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.Post.FindAllPosts()

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)

}

func (h *Handler) handlerPostsById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, byidErr := h.services.Post.FindPostByID(int(id))

	if byidErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, byidErr)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}

func (h *Handler) handlerPostsCreate(w http.ResponseWriter, r *http.Request) {

	var post request.CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.Post.Create(post)

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)

}

func (h *Handler) handlerPostsUpdate(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	var post request.UpdatePostRequest

	post.PostID = int(id)

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, updateErr := h.services.Post.Update(post)

	if updateErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, updateErr)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)

}

func (h *Handler) handlerPostsDelete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, deleteErr := h.services.Post.Delete(int(id))

	if deleteErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, deleteErr)
		return
	}

	response.ResponseMessage(w, "Success", res, http.StatusOK)
}
