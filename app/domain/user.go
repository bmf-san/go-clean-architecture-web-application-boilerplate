package domain

// A Users belong to the domain layer.
type Users []User

// A User belong to the domain layer.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
