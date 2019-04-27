package repository

import (
	"../database"
	"../model"
)

type UserRepository interface {
	GetUsers() (*model.User, error)
}

type UserRepositoryImpl struct {
	SqlHandler database.SqlHandler
}

func (repository *UserRepositoryImpl) GetUsers() (*model.User, error) {
	const query = `
		select 
			id,
			name
		from
			users
	`
	var user model.User
	err := repository.SqlHandler.Connect().QueryRow(query).Scan(&user.ID, &user.Name)
	return &user, err
}
