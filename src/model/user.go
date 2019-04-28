package model

type Users []User

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
