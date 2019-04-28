package domain

// A Posts belong to the domain layer.
type Posts []Post

// A Post belong to the domain layer.
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Body   string `json:"body"`
}
