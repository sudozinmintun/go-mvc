package bootstrap

import (
	"go-pongo2-demo/internal/app/http/controllers"
	"go-pongo2-demo/internal/app/repo"
	"go-pongo2-demo/internal/app/service"
)

type Container struct {
	CategoryController controllers.CategoryController
	AuthController     controllers.AuthController
}

func Init() *Container {
	// repos
	categoryRepo := repo.NewCategoryRepository()

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
