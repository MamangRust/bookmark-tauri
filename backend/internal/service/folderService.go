package service

import (
	"bookmark-backend/internal/domain/response"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type folderService struct{}

func NewFolderService() *folderService {
	return &folderService{}
}

func (fs *folderService) CreateFolder(name string) (string, error) {
	bookmarkPath := "static/bookmark"
	if _, err := os.Stat(bookmarkPath); os.IsNotExist(err) {
		err := os.Mkdir(bookmarkPath, 0755)
		if err != nil {
			return "", err
		}
	}

	path := filepath.Join(bookmarkPath, name)
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		return "", err
	}
	fmt.Printf("Folder '%s' berhasil dibuat\n", path)

	return path, nil
}
func (fs *folderService) FindAllFolder() ([]response.FolderDataResponse, error) {
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

func (fs *folderService) CheckIfFolderExists(folderName string) (bool, error) {
	_, err := os.Stat("bookmark/" + folderName)
	if err == nil {
		return true, nil // Folder ditemukan
	}
	if os.IsNotExist(err) {
		return false, nil // Folder tidak ditemukan
	}
	return false, err // Terjadi error lainnya
}

func (fs *folderService) CheckAndUpdateFolder(oldFolder string, newFolder string) error {
	oldPath := "bookmark/" + oldFolder
	newPath := "bookmark/" + newFolder

	_, err := os.Stat(newPath)
	if err == nil {
		return fmt.Errorf("folder '%s' already exists", newPath)
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Printf("Error renaming folder: %v\n", err)
		return err
	}

	fmt.Printf("Folder '%s' successfully renamed to '%s'\n", oldPath, newPath)
	return nil
}

func (fs *folderService) DeleteFolder(folder string) error {
	path := "static/bookmark/" + folder

	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	fmt.Printf("Folder '%s' berhasil dihapus\n", path)
	return nil
}
