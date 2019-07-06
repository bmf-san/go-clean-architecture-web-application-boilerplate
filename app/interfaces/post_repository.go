package interfaces

import "github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"

// A PostRepository belong to the inteface layer
type PostRepository struct {
	SQLHandler SQLHandler
}

// FindAll is returns all entities.
func (pr *PostRepository) FindAll() (posts domain.Posts, err error) {
	const query = `
	SELECT
		id,
		user_id,
		body
	FROM
		posts
	`

	rows, err := pr.SQLHandler.Query(query)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var userID int
		var body string
		if err = rows.Scan(&id, &userID, &body); err != nil {
			return
		}
		post := domain.Post{
			ID:     id,
			UserID: userID,
			Body:   body,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// Save is saves the given entity.
func (pr *PostRepository) Save(p domain.Post) (id int64, err error) {
	// NOTE: this is a transaction example.
	tx, err := pr.SQLHandler.Begin()
	if err != nil {
		return
	}

	const query = `
		INSERT INTO
			posts(user_id, body)
		VALUES
			(?, ?)
	`

	result, err := tx.Exec(query, p.UserID, p.Body)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, nil
	}

	return
}

// DeleteByID is deletes the entity identified by the given id.
func (pr *PostRepository) DeleteByID(postID int) (err error) {
	const query = `
		DELETE
		FROM
			posts
		WHERE
			id = ?
	`

	_, err = pr.SQLHandler.Exec(query, postID)
	if err != nil {
		return
	}

	return
}
