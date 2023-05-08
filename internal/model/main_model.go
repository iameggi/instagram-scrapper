package model

type Post struct {
    ID        string `json:"id"`
    Caption   string `json:"caption"`
    ImageURL  string `json:"imageUrl"`
    Timestamp int64  `json:"timestamp"`
}
