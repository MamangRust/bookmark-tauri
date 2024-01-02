package repository

import (
	"bookmark-backend/internal/models"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAllCategory() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindCategoryByID(categoryID uint) (models.Category, error) {
	var category models.Category
	if err := r.db.Where("category_id = ?", categoryID).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) FindCategoryByName(name string) (models.Category, error) {
	var category models.Category
	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) CreateCategory(category models.Category) error {
	if err := r.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) UpdateCategory(category models.Category) error {
	if err := r.db.Model(&category).Updates(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(category models.Category) error {
	if err := r.db.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
