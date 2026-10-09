package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func (s *Store) AddList(name string) error {
	_, err := s.db.Exec(
		"INSERT INTO lists (name) VALUES (?)",
		name,
	)
	return err
}

func (s Store) GetListID(name string) (int, error) {
	var id int
	err := s.db.QueryRow(
		"SELECT id FROM lists WHERE name = ?",
		name,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s Store) GetListEntries(id int) ([]string, error) {
	var content []string

	rows, err := s.db.Query(
		`SELECT id, name, data
		 FROM entries
		 WHERE list_id = ?
		 ORDER BY id ASC`,
		id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var name string
		var data string

		if err := rows.Scan(&id, &name, &data); err != nil {
			return nil, err
		}

		content = append(content, fmt.Sprintf("%d %s %s", id, name, data))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func (s Store) GetLists() ([]string, error) {
	var content []string

	rows, err := s.db.Query(
		`SELECT id, name
		 FROM lists
		 ORDER BY id ASC`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		content = append(content, fmt.Sprintf("%d %s", id, name))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func (s *Store) AddEntry(list_id int, name, data string) error {
	_, err := s.db.Exec(
		"INSERT INTO entries (list_id, name, data) VALUES (?, ?, ?)",
		list_id, name, data,
	)
	return err
}

func (s *Store) Close() error {
	return s.db.Close()
}

func NewStore(path string) (*Store, error) {
	var zero Store

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return &zero, err
	}

	if err := db.Ping(); err != nil {
		return &zero, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS lists (
			id   INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS entries (
			id      INTEGER PRIMARY KEY,
			list_id INTEGER NOT NULL REFERENCES lists(id),
			name    TEXT NOT NULL,
			data    TEXT NOT NULL DEFAULT '{}'
		);
	`)
	if err != nil {
		return &zero, nil
	}

	return &Store{db: db}, nil
}
