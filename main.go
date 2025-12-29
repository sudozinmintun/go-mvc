
package main

import (
    "log"
    "go-pongo2-demo/internal/database"
    "go-pongo2-demo/internal/router"
    "go-pongo2-demo/internal/view"
    "github.com/labstack/echo/v4"
)

func main() {
    database.ConnectAndMigrate()

    e := echo.New()
    e.Renderer = view.NewPongoRenderer()

    router.Setup(e)

    log.Println("Server running at http://localhost:8080")
    e.Logger.Fatal(e.Start(":8080"))
}
