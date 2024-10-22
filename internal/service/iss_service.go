package service

import (
    "encoding/json"
    "net/http"
    "iss-tracker/internal/models"
)

func FetchISSLocation() (*models.ISSLocation, error) {
    resp, err := http.Get("http://api.open-notify.org/iss-now.json")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var location models.ISSLocation
    if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
        return nil, err
    }

    return &location, nil
}
