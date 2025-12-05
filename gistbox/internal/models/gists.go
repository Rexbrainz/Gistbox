package models

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, title, content, created, expires FROM gists
						WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Retrieve the row from the database, returns a sql.Row object
	row := m.DB.QueryRow(stmt, id)

	var g Gist

	// .Scan() is used to copy the values from each field in the sql.Row object
	// It takes pointers to each field, the number of arguments must be same as 
	// the number of columns returned by statement.
	err := row.Scan(&g.ID, &g.Title, &g.Content, &g.Created, &g.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Gist{}, ErrNoRecord
		} else {
			return Gist{}, err
		}
	}
	return g, nil
}

// Return the ten most recently created gists.
func (m *GistModel) Latest() ([]Gist, error) {
	stmt := `SELECT id, title, content, created, expires FROM gists
						WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// Critical to close the rows, otherwise the connection to the underlying
	// database remains open and if something goes wrong, the connections in the
	// pool could be used up.
	defer rows.Close()

	var gists []Gist

	for rows.Next() {
		var g Gist

		err = rows.Scan(&g.ID, &g.Title, &g.Content, &g.Created, &g.Expires)
		if err != nil {
			return nil, err
		}
		gists = append(gists, g)
	}

	// Do not assume that the iteration was completed successfully with out any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gists, nil
}
