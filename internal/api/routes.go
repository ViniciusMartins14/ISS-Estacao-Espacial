package api

import (
    "github.com/labstack/echo/v4"
    "iss-tracker/internal/api/handlers"
)

func SetupRoutes(e *echo.Echo) {
    e.GET("/location", handlers.GetISSLocation)
}
