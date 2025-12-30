package bootstrap

import (
	"go-pongo2-demo/internal/app/http/controllers"
	repo "go-pongo2-demo/internal/app/repository"
	"go-pongo2-demo/internal/app/service"

	"gorm.io/gorm"
)

type Container struct {
	CategoryController controllers.CategoryController
	AuthController     controllers.AuthController
}

func Init(db *gorm.DB) *Container {
	// repos
	categoryRepo := repo.NewCategoryRepository(db)

	// services
	categoryService := service.NewCategoryService(categoryRepo)

	// controllers
	categoryController := controllers.NewCategoryController(categoryService)
	authController := controllers.AuthController{}

	return &Container{
		CategoryController: categoryController,
		AuthController:     authController,
	}
}
