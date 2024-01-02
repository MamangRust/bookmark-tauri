package handler

import (
	"bookmark-backend/domain/request"
	"bookmark-backend/service"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type FolderHandler struct {
	FolderService service.FolderService
}

func NewFolderHandler(folderService service.FolderService) *FolderHandler {
	return &FolderHandler{
		FolderService: folderService,
	}
}

func (h *FolderHandler) FindAllFolder(w http.ResponseWriter, r *http.Request) {
	folders, err := h.FolderService.FindAllFolder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(folders)
}

func (h *FolderHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	var folder request.FolderDataRequest
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.FolderService.CreateFolder(folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *FolderHandler) FindFolder(w http.ResponseWriter, r *http.Request) {
	folder := chi.URLParam(r, "folder")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filesData, err := h.FolderService.FindFolder(folder)
	if err != nil {
		if strings.Contains(err.Error(), "folder not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(filesData)
}

func (h *FolderHandler) UpdateFolder(w http.ResponseWriter, r *http.Request) {
	oldFolderName := chi.URLParam(r, "folder")
	var newFolder request.FolderDataRequest
	if err := json.NewDecoder(r.Body).Decode(&newFolder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.FolderService.UpdateFolder(oldFolderName, newFolder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *FolderHandler) DeleteFolder(w http.ResponseWriter, r *http.Request) {
	folderName := chi.URLParam(r, "folder")

	if err := h.FolderService.DeleteFolder(folderName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
