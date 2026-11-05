package library

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

type Store struct {
	DB *sql.DB
}

func Open(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "hithlain.db")

	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	conn.SetMaxOpenConns(1)

	if _, err := conn.Exec("PRAGMA foreign_keys = ON"); err != nil {
		conn.Close()
		return nil, err
	}

	store := &Store{DB: conn}

	if err := store.Init(); err != nil {
		conn.Close()
		return nil, err
	}

	return store, nil
}

func (s *Store) Close() error {
	return s.DB.Close()
}
