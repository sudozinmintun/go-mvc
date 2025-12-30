package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"pmsys/internal/bootstrap"
	"pmsys/internal/config"
	"pmsys/internal/database"
	"pmsys/internal/router"
	"pmsys/internal/view"
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
