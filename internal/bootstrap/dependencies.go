package bootstrap

import (
	"go-pongo2-demo/internal/app/http/handlers"
	repo "go-pongo2-demo/internal/app/repository"
	"go-pongo2-demo/internal/app/service"

	"gorm.io/gorm"
)

type Container struct {
	CategoryController handlers.CategoryController
	AuthController     handlers.AuthController
}

func Init(db *gorm.DB) *Container {
	// repos
	categoryRepo := repo.NewCategoryRepository(db)

	// services
	categoryService := service.NewCategoryService(categoryRepo)

	// controllers
	categoryController := handlers.NewCategoryController(categoryService)
	authController := handlers.AuthController{}

	return &Container{
		CategoryController: categoryController,
		AuthController:     authController,
	}
}
