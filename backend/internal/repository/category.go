package repository

import (
	"bookmark-backend/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAllCategory() (*[]models.Category, error) {
	var categories []models.Category

	db := r.db.Model(&categories)

	checkCategories := db.Debug().Preload("Posts").Find(&categories)

	if checkCategories.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &categories, nil
}

func (r *categoryRepository) FindCategoryByID(categoryID int) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategory := db.Debug().Preload("Posts").Where("id = ?", categoryID).Find(&category)

	if checkCategory.RowsAffected < 1 {
		return &category, gorm.ErrRecordNotFound
	}

	return &category, nil
}

func (r *categoryRepository) FindCategoryByName(name string) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategory := db.Debug().Where("name = ?", name).Find(&category)

	if checkCategory.RowsAffected < 1 {
		return &category, gorm.ErrRecordNotFound
	}

	return &category, nil
}

func (r *categoryRepository) CreateCategory(request models.Category) (*models.Category, error) {

	db := r.db.Model(&request)

	checkCategoryName := db.Debug().Where("name = ?", request.Name).Find(&request)

	if checkCategoryName.RowsAffected > 0 {
		return nil, gorm.ErrDuplicatedKey
	}
	addCategory := db.Debug().Create(&request).Commit()

	if addCategory.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &request, nil
}

func (r *categoryRepository) UpdateCategory(request models.Category) (*models.Category, error) {
	var existingCategory models.Category
	if err := r.db.First(&existingCategory, "id = ?", request.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	existingCategory.Name = request.Name
	existingCategory.Image = request.Image
	existingCategory.Description = request.Description

	if err := r.db.Save(&existingCategory).Error; err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return &existingCategory, nil
}

func (r *categoryRepository) DeleteCategory(id int) error {
	var deletecategory models.Category

	db := r.db.Model(&deletecategory)

	checkCategory := db.Debug().Where("id = ?", id).First(&deletecategory)

	if checkCategory.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	deleteCategory := db.Debug().Delete(&deletecategory)

	if deleteCategory.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
