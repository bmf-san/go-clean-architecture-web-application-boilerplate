package repository

import (
	"log"

	"../database"
	"../model"
)

type UserRepository interface {
	GetUser() (*model.User, error)
	GetUsers() (*model.Users, error)
}

type UserRepositoryImpl struct {
	SqlHandler database.SqlHandler
}

func (repository *UserRepositoryImpl) GetUser() (*model.User, error) {
	// TODO: want to get a id from query params
	var userID int
	userID = 1

	const query = `
		select
			id,
			name
		from
			users
		where
			id = ?
	`
	rows, err := repository.SqlHandler.Query(query, userID)

	defer rows.Close()

	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &user, err
}

func (repository *UserRepositoryImpl) GetUsers() (*model.Users, error) {
	rows, err := repository.SqlHandler.Query(`
		select 
			id,
			name
		from
			users
	`)

	defer rows.Close()

	var users model.Users
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatal(err)

		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &users, err
}
