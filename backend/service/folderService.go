package service

import (
	"bookmark-backend/domain/request"
	"bookmark-backend/domain/response"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FolderService struct{}

func NewFolderService() *FolderService {
	return &FolderService{}
}

func (fs *FolderService) CreateFolder(folder request.FolderDataRequest) error {
	path := "bookmark/" + folder.Name
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	fmt.Printf("Folder '%s' berhasil dibuat\n", path)
	return nil
}

func (fs *FolderService) FindAllFolder() ([]response.FolderDataResponse, error) {
	path := "bookmark"

	folders, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var folderStructs []response.FolderDataResponse

	for _, folder := range folders {
		if folder.IsDir() {
			folderStructs = append(folderStructs, response.FolderDataResponse{Name: folder.Name()})
		}
	}

	return folderStructs, nil
}

func (fs *FolderService) FindFolder(folder string) ([]response.FileDataResponse, error) {
	path := "bookmark/" + folder

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filesData []response.FileDataResponse

	if len(files) == 0 {
		return nil, fmt.Errorf("folder '%s' tidak berisi file", folder)
	}

	for _, file := range files {
		filePath := path + "/" + file.Name()

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file '%s': %s\n", file.Name(), err)
			continue
		}

		// Menghilangkan karakter baris baru
		cleanContent := strings.ReplaceAll(string(fileContent), "\n", "")

		fileData := response.FileDataResponse{
			Title:   strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())),
			Content: cleanContent,
		}

		filesData = append(filesData, fileData)
	}

	return filesData, nil
}

func (fs *FolderService) UpdateFolder(oldFolder string, newFolder request.FolderDataRequest) error {
	oldPath := "bookmark/" + oldFolder
	newPath := "bookmark/" + newFolder.Name

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}
	fmt.Printf("Folder '%s' berhasil diubah menjadi '%s'\n", oldPath, newPath)
	return nil
}

func (fs *FolderService) DeleteFolder(folder string) error {
	path := "bookmark/" + folder

	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	fmt.Printf("Folder '%s' berhasil dihapus\n", path)
	return nil
}
