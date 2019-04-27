package interfaces

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A UserRepository belong to the inteface layer
type UserRepository struct {
	SQLHandler SQLHandler
}

// FindAll return a listing of the resource of users.
func (ur *UserRepository) FindAll() (users domain.Users, err error) {
	const query = `
		select
			id,
			name
		from
			users
	`
	rows, err := ur.SQLHandler.Query(query)

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

// FindByID return the specified resource of a user.
func (ur *UserRepository) FindByID(userID int) (user domain.User, err error) {
	const query = `
		select
			id,
			name
		from
			users
		where
			id = ?
	`
	row, err := ur.SQLHandler.Query(query, userID)

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
