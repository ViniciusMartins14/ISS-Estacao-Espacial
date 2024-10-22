package models

type ISSLocation struct {
    Message     string `json:"message"`
    ISSPosition struct {
        Latitude  string `json:"latitude"`
        Longitude string `json:"longitude"`
    } `json:"iss_position"`
    Timestamp int64 `json:"timestamp"`
}
