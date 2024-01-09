package upload

import (
	"bookmark-backend/internal/domain/response"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HandleImageUpload(w http.ResponseWriter, r *http.Request) string {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "file too large")

	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid file")

	}
	defer file.Close()

	// Validate file extension
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := filepath.Ext(fileHeader.Filename)
	if !allowedExtensions[ext] {
		response.ResponseError(w, http.StatusBadRequest, "invalid file format")

	}

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}
	filePath := filepath.Join(uploadDir, fileHeader.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "could not create the file")

	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "could not write the file")

	}

	return filePath
}
