package database

import (
	"database/sql"

	"github.com/nollidnosnhoj/kopalol/queries"
	_ "github.com/tursodatabase/go-libsql"
)

type Database struct {
	db      *sql.DB
	queries *queries.Queries
}

func New(url string) (*Database, error) {
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
