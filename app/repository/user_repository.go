package repository

import (
	"database/sql"

	"../model"
)

func UserRepositoryIndex(db *sql.DB) (*model.UserModel, error) {
	const query = `
		select 
			id,
			name
		from
			users
	`
	var user model.UserModel
	err := db.QueryRow(query).Scan(&user.ID, &user.Name)
	return &user, err
}
