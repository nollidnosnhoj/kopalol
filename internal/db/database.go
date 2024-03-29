package db

import (
	"database/sql"
	"errors"

	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/queries"
	_ "github.com/tursodatabase/go-libsql"
)

type Database struct {
	db      *sql.DB
	queries *queries.Queries
}

func New(config *config.Configuration) (*Database, error) {
	url := config.DATABASE_URL
	if url == "" {
		return nil, errors.New("DATABASE_URL is required")
	}
	authToken := config.DATABASE_AUTH_TOKEN
	if authToken != "" {
		url = url + "?authToken=" + authToken
	}
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}
	queries := queries.New(db)
	return &Database{db, queries}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) Db() *sql.DB {
	return d.db
}

func (d *Database) Queries() *queries.Queries {
	return d.queries
}
