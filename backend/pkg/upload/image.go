package upload

import (
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/service"
	"fmt"
	"net/http"
	"path/filepath"
)

type ImageUpload interface {
	HandleImageUpload(category string, w http.ResponseWriter, r *http.Request) string
}

type Image struct {
	fileService   service.FileService
	folderService service.FolderService
}

func NewImage(fileService service.FileService, folderService service.FolderService) *Image {
	return &Image{
		fileService:   fileService,
		folderService: folderService,
	}
}

func (i *Image) HandleImageUpload(category string, w http.ResponseWriter, r *http.Request) string {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "file too large")
		return ""
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid file")
		return ""
	}
	defer file.Close()

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := filepath.Ext(fileHeader.Filename)
	if !allowedExtensions[ext] {
		response.ResponseError(w, http.StatusBadRequest, "invalid file format")
		return ""
	}

	path_folder, err := i.folderService.CreateFolder(category)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "could not create the folder")
		return ""
	}

	fmt.Println("path_folder: ", path_folder)

	filePath := filepath.Join(path_folder, fileHeader.Filename)

	fmt.Println("filePath: ", filePath)

	filePath_file, err := i.fileService.CreateFileImage(file, filePath)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "could not create the file")
		return ""
	}

	return filePath_file
}
