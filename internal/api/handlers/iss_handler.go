package handlers

import (
    "net/http"
    "iss-tracker/internal/service"
    "github.com/labstack/echo/v4"
)

func GetISSLocation(c echo.Context) error {
    location, err := service.FetchISSLocation()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch ISS location"})
    }
    return c.JSON(http.StatusOK, location)
}
