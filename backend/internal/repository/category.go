package repository

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/models"

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

	checkCategories := db.Debug().Find(&categories)

	if checkCategories.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &categories, nil
}

func (r *categoryRepository) FindCategoryByID(categoryID int) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategory := db.Debug().Where("category_id = ?", categoryID).Find(&category)

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

func (r *categoryRepository) CreateCategory(request request.CreateCategoryRequest) (*models.Category, error) {
	var newCategory models.Category

	db := r.db.Model(&newCategory)

	newCategory.Name = request.Name
	newCategory.Image = request.Image
	newCategory.Description = request.Description

	checkCategoryName := db.Debug().Where("name = ?", newCategory.Name).Find(&newCategory)

	if checkCategoryName.RowsAffected > 0 {
		return nil, gorm.ErrDuplicatedKey
	}
	addCategory := db.Debug().Create(&newCategory).Commit()

	if addCategory.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &newCategory, nil
}

func (r *categoryRepository) UpdateCategory(request request.UpdateCategoryRequest) (*models.Category, error) {
	var newCategory models.Category

	db := r.db.Model(&newCategory)

	checkCategory := db.Debug().Where("category_id = ?", request.CategoryID).Find(&newCategory)

	if checkCategory.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	newCategory.Name = request.Name
	newCategory.Image = request.Image
	newCategory.Description = request.Description

	updateCategory := db.Debug().Updates(&newCategory)

	if updateCategory.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &newCategory, nil

}

func (r *categoryRepository) DeleteCategory(id int) error {
	var deletecategory models.Category

	db := r.db.Model(&deletecategory)

	checkCategory := db.Debug().Where("category_id = ?", id).First(&deletecategory)

	if checkCategory.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	deleteCategory := db.Debug().Delete(&deletecategory)

	if deleteCategory.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
