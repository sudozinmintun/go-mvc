package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"go-pongo2-demo/internal/bootstrap"
	"go-pongo2-demo/internal/config"
	"go-pongo2-demo/internal/database"
	"go-pongo2-demo/internal/router"
	"go-pongo2-demo/internal/view"
)

func main() {
	config.LoadEnv()

	db, err := database.ConnectAndMigrate()
	if err != nil {
		log.Fatal(err)
	}

	deps := bootstrap.Init(db)

	e := echo.New()
	e.Renderer = view.NewPongoRenderer()

	router.Setup(e, deps)

	log.Println("Server running at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
