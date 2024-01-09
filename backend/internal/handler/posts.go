package handler

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/middlewares"
	"bookmark-backend/pkg/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initPostsGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)
		r.Get("/hello", h.handlerPostsHello)

		r.Get("/", h.handlerPostsAll)
		r.Get("/{id}", h.handlerPostsById)
		r.Post("/create", h.handlerPostsCreate)
		r.Put("/update/{id}", h.handlerPostsUpdate)
		r.Delete("/delete/{id}", h.handlerPostsDelete)
	})
}

func (h *Handler) handlerPostsHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Posts"))
}

func (h *Handler) handlerPostsAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.Post.FindAllPosts()

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)

}

func (h *Handler) handlerPostsById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	res, byidErr := h.services.Post.FindPostByID(int(id))

	if byidErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, byidErr.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}

func (h *Handler) handlerPostsCreate(w http.ResponseWriter, r *http.Request) {

	userId := auth.GetContextUserId(r)

	var post request.CreatePostRequest

	userIdInt, err := strconv.Atoi(userId)

	fmt.Println("userIdInt: ", userIdInt)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {

		fmt.Println("err: ", err.Error())
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := post.Validate(); err != nil {
		response.HandleValidationErrors(w, err)

		return
	}

	res, err_create := h.services.Post.Create(userIdInt, post)

	if err != nil {

		response.ResponseError(w, http.StatusInternalServerError, err_create.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)

}

func (h *Handler) handlerPostsUpdate(w http.ResponseWriter, r *http.Request) {

	userId := auth.GetContextUserId(r)

	userIdInt, err := strconv.Atoi(userId)

	fmt.Println("userIdInt: ", userIdInt)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var post request.UpdatePostRequest

	post.PostID = int(id)

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := post.Validate(); err != nil {
		response.HandleValidationErrors(w, err)

		return
	}

	res, updateErr := h.services.Post.Update(userIdInt, post)

	if updateErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, updateErr.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)

}

func (h *Handler) handlerPostsDelete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {

		response.ResponseError(w, http.StatusBadRequest, "invalid id")
		return
	}

	res, deleteErr := h.services.Post.Delete(int(id))

	if deleteErr != nil {

		response.ResponseError(w, http.StatusInternalServerError, deleteErr.Description)
		return
	}

	response.ResponseMessage(w, "Success", res.Data, http.StatusOK)
}
