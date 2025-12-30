package repo

import (
	"go-pongo2-demo/internal/app/models"
	"go-pongo2-demo/internal/database"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	Create(name string) error
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Create(name string) error {
	return database.DB.Create(&models.Category{Name: name}).Error
}
