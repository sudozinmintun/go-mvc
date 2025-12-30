package bootstrap

import (
	"pmsys/internal/app/http/handlers"
	repo "pmsys/internal/app/repository"
	"pmsys/internal/app/service"

	"gorm.io/gorm"
)

type Container struct {
	CategoryController handlers.CategoryController
	AuthController     handlers.AuthController
}

func Init(db *gorm.DB) *Container {
	// repos
	categoryRepo := repo.NewCategoryRepository(db)
	userRepo := repo.NewUserRepository(db)

	// services
	categoryService := service.NewCategoryService(categoryRepo)
	authService := service.NewAuthService(userRepo)

	// controllers
	categoryController := handlers.NewCategoryController(categoryService)
	authController := handlers.NewAuthController(authService)

	return &Container{
		CategoryController: categoryController,
		AuthController:     authController,
	}
}
