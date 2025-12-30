package service

import (
	"errors"

	dto "go-pongo2-demo/internal/app/http/requests"
	"go-pongo2-demo/internal/app/models"
	"go-pongo2-demo/internal/app/repo"

	"github.com/go-playground/validator/v10"
)

// Public interface (used by controllers)
type CategoryService interface {
	GetAll() ([]models.Category, error)
	Create(input dto.CreateCategoryDTO) error
}

// private implementation
type categoryService struct {
	repo      repo.CategoryRepository
	validator *validator.Validate
}

// constructor returns the interface
func NewCategoryService(r repo.CategoryRepository) CategoryService {
	return &categoryService{
		repo:      r,
		validator: validator.New(),
	}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) Create(input dto.CreateCategoryDTO) error {
	if err := s.validator.Struct(input); err != nil {
		// You can improve this later to return field-specific errors
		return errors.New("name is required and must be at least 2 characters")
	}

	return s.repo.Create(input.Name)
}
