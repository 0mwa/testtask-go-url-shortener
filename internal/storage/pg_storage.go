package storage

import (
	"database/sql"
	"fmt"
)

type pgStorage struct {
	db *sql.DB
}

func NewPgStorage(db *sql.DB) Storage {
	return &pgStorage{db: db}
}

func (s *pgStorage) Write(sURL string, oURL string) error {
	stmtInsUrl, err := s.db.Prepare("INSERT INTO links (original_url, short_url) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("cannot prepare statement")
	}

	_, err = stmtInsUrl.Exec(oURL, sURL)
	return err
}
func (s *pgStorage) Read(sURL string) (string, error) {
	stmtSelUrl, err := s.db.Prepare("SELECT original_url FROM links WHERE short_url = $1")
	if err != nil {
		return "", fmt.Errorf("cannot prepare statement")
	}

	var originalURL string

	err = stmtSelUrl.QueryRow(sURL).Scan(&originalURL)
	return originalURL, err
}
