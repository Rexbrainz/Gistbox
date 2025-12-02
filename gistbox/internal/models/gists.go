package models

import (
	"database/sql"
	"time"
)

type Gist struct {
	ID 			int
	Title		string
	Content	string
	Created	time.Time
	Expires	time.Time
}

// Define a Gist Model type which wraps a sql.DB connection pool.
type GistModel struct {
	DB *sql.DB
}

// Insert a new gist into the database.
func (m *GistModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO gists (title, content, created, expires)
						VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Execute an insert statemet.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method to get the ID of our new inserted record in the gists table.
	id , err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get a specific gist based on its id.
func (m *GistModel) Get(id int) (Gist, error) {
	return Gist{}, nil
}

// Return the ten most recently created gists.
func (m *GistModel) Latest() ([]Gist, error) {
	return nil, nil
}
