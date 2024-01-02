package service

import (
	"bookmark-backend/domain/request"
	"bookmark-backend/domain/response"
	"fmt"
	"os"
	"strings"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (fs *FileService) CreateFile(request request.CreateFileRequest) error {
	filePath := "bookmark/" + request.Folder + "/" + request.Title
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(request.Content)
	if err != nil {
		return err
	}

	fmt.Printf("File '%s' berhasil dibuat\n", filePath)
	return nil
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

// Penggunaan di dalam fungsi FindFile
func (fs *FileService) FindFile(request request.FileRequest) (response.FileDataResponse, error) {
	filePath := "bookmark/" + request.Folder + "/" + request.FileName

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return response.FileDataResponse{}, err
	}

	title, date, author, content := extractMetadata(string(fileContent))

	fileData := response.FileDataResponse{
		Title:   title,
		Date:    date,
		Author:  author,
		Content: content,
	}

	fmt.Printf("File '%s' ditemukan\n", filePath)
	return fileData, nil
}

func (fs *FileService) UpdateFile(request request.UpdateFileRequest) (response.FileDataResponse, error) {
	filePath := "bookmark/" + request.Folder + "/" + request.Title

	err := os.WriteFile(filePath, []byte(request.Content), 0644)
	if err != nil {
		return response.FileDataResponse{}, err
	}

	fmt.Printf("File '%s' berhasil diubah\n", filePath)
	return response.FileDataResponse{Title: request.Title, Content: request.Content}, nil
}

func (fs *FileService) DeleteFile(request request.FileRequest) error {
	filePath := "bookmark/" + request.Folder + "/" + request.FileName

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	fmt.Printf("File '%s' berhasil dihapus\n", filePath)
	return nil
}
