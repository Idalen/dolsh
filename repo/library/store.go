package repo

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

func New() (*Repository, error) {
	db, err := sql.Open("sqlite", "dolsh.db")
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:db,
	}, nil
}

func (r *Repository) GetTracks() error {
	return errors.New("Unimplemented") 
}
