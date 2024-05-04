package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created)
    VALUES(?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created FROM snippets
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var s Snippet
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created FROM snippets
    WHERE created <= UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
