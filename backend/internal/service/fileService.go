package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type fileService struct {
}

func NewFileService() *fileService {
	return &fileService{}
}

func (fs *fileService) CreateFileImage(file multipart.File, filePath string) (string, error) {

	newFile, err := os.Create(filePath)
	if err != nil {
		return "", errors.New("failed to create file")
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)

	if err != nil {
		return "", errors.New("failed to copy file")
	}

	return filePath, nil
}

func (fs *fileService) CreateFile(request request.CreateFileRequest) (*response.ServiceResponse, *response.ServiceError) {
	filePath := "static/bookmark/" + request.Folder + "/" + request.Title + ".md"

	fmt.Printf("Membuat file '%s'\n", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, &response.ServiceError{Err: err, Description: "Failed to create file"}
	}
	defer file.Close()

	_, err = file.WriteString(request.Content)
	if err != nil {
		return nil, &response.ServiceError{Err: err, Description: "Failed to write to file"}
	}

	fmt.Printf("File '%s' berhasil dibuat\n", filePath)
	return &response.ServiceResponse{
		Data:  "File created successfully " + filePath,
		Error: nil,
	}, nil
}

func extractMetadata(content string) (title, date, author, extractedContent string) {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "title:") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "title:"))
		} else if strings.HasPrefix(line, "date:") {
			date = strings.TrimSpace(strings.TrimPrefix(line, "date:"))
		} else if strings.HasPrefix(line, "author:") {
			author = strings.TrimSpace(strings.TrimPrefix(line, "author:"))
		} else {
			extractedContent += line + "\n"
		}
	}

	extractedContent = strings.TrimSpace(extractedContent)

	return title, date, author, extractedContent
}

func (fs *fileService) UpdateFile(request request.UpdateFileRequest) (*response.ServiceResponse, *response.ServiceError) {
	oldFilePath := "static/bookmark/" + request.Folder + "/" + request.OldTitle + ".md"
	newFilePath := "static/bookmark/" + request.Folder + "/" + request.NewTitle + ".md"

	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return nil, &response.ServiceError{Err: err, Description: "Failed to rename file"}
	}

	err = os.WriteFile(newFilePath, []byte(request.Content), 0644)
	if err != nil {
		// Jika penulisan file baru gagal, kita harus memulihkan nama file lama
		// karena file sudah terubah namanya sebelumnya
		os.Rename(newFilePath, oldFilePath)
		return nil, &response.ServiceError{Err: err, Description: "Failed to update file"}
	}

	fmt.Printf("File '%s' berhasil diubah menjadi '%s'\n", oldFilePath, newFilePath)

	return &response.ServiceResponse{
		Data:  "File updated successfully",
		Error: nil,
	}, nil
}

func (fs *fileService) FindFile(request request.FileRequest) (*response.ServiceResponse, *response.ServiceError) {
	filePath := "static/bookmark/" + request.Folder + "/" + request.FileName

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, &response.ServiceError{Err: err, Description: "File not found"}
	}

	title, date, author, content := extractMetadata(string(fileContent))

	fileData := response.FileDataResponse{
		Title:   title,
		Date:    date,
		Author:  author,
		Content: content,
	}

	fmt.Printf("File '%s' ditemukan\n", filePath)
	return &response.ServiceResponse{
		Data:  fileData,
		Error: nil,
	}, nil
}

func (fs *fileService) DeleteFile(filePath string) error {

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	fmt.Printf("File '%s' berhasil dihapus\n", filePath)
	return nil
}
