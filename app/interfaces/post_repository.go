package interfaces

// A PostRepository belong to the inteface layer
type PostRepository struct {
	SQLHandler SQLHandler
}

// AddPost add a newly created resource of a Post.
func (ur *PostRepository) AddPost() (id int64, err error) {
	// NOTE: this is a transaction example.
	tx, err := ur.SQLHandler.Begin()
	if err != nil {
		return
	}

	const query = `
		insert into
			posts(id, user_id, body)
		values
			(?, ?, ?)
	`

	_, err = tx.Exec(query, 1, 1, "hoge")
	if err != nil {
		return
	}

	result, err := tx.Exec(query, 100, 1, "Hello, World. This is newly inserted.")
	if err != nil {
		return
	}

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
