package mysql

import (
	"database/sql"
	"errors"

	"github.com/coderste/web-application/pkg/models"
)

// SnippetModel holds our sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// The SQL statement we want to run.
	// NB: We use ? to indicate a placeholder parameter
	statement := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the statement
	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use LastInsertId() to get the ID of the data
	// we just inserted into the database
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	// id above is returned as an int64 so we cast it to an int
	return int(id), nil
}

// Get will return a specific method based on the given ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	snippet := &models.Snippet{}
	err := m.DB.QueryRow("SELECT ...", id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return snippet, nil
}

// Latest will return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
