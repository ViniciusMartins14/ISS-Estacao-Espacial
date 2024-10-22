package main

import (
    "iss-tracker/internal/api"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "log"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    api.SetupRoutes(e)

    log.Fatal(e.Start(":3000"))
}
