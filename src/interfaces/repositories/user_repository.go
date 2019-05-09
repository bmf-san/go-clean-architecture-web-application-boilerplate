package repositories

import (
	"../../domain"
	"../database"
)

type UserRepository struct {
	SqlHandler database.SqlHandler
}

func (repository *UserRepository) FindAll() (users domain.Users, err error) {
	const query = `
		select 
			id,
			name
		from
			users
	`
	rows, err := repository.SqlHandler.Query(query)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			return
		}
		user := domain.User{
			ID:   id,
			Name: name,
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (repository *UserRepository) FindById(userId int) (user domain.User, err error) {
	const query = `
		select
			id,
			name
		from
			users
		where
			id = ?
	`
	row, err := repository.SqlHandler.Query(query, userId)

	defer row.Close()

	if err != nil {
		return
	}

	var id int
	var name string
	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	user.ID = id
	user.Name = name

	return
}
