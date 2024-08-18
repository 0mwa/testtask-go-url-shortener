package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func NewPostgresClient() (*sql.DB, error) {
	file := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"))
	db, err := sql.Open("postgres", file)
	if err != nil {
		return nil, err
	}
	return db, nil
}
