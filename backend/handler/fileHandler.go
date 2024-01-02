package handler

import (
	"bookmark-backend/domain/request"
	"bookmark-backend/service"
	"encoding/json"
	"net/http"
)

type FileHandler struct {
	FileService service.FileService
}

func NewFileHandler(fileService service.FileService) *FileHandler {

	return &FileHandler{
		FileService: fileService,
	}
}

func (fh *FileHandler) CreateFile(w http.ResponseWriter, r *http.Request) {
	var createFileReq request.CreateFileRequest
	if err := json.NewDecoder(r.Body).Decode(&createFileReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := fh.FileService.CreateFile(createFileReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (fh *FileHandler) FindFile(w http.ResponseWriter, r *http.Request) {
	var fileReq request.FileRequest
	if err := json.NewDecoder(r.Body).Decode(&fileReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fileData, err := fh.FileService.FindFile(fileReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fileData)
}

func (fh *FileHandler) UpdateFile(w http.ResponseWriter, r *http.Request) {
	var updateFileReq request.UpdateFileRequest
	if err := json.NewDecoder(r.Body).Decode(&updateFileReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fileData, err := fh.FileService.UpdateFile(updateFileReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fileData)
}

func (fh *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	var fileReq request.FileRequest
	if err := json.NewDecoder(r.Body).Decode(&fileReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := fh.FileService.DeleteFile(fileReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
