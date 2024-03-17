package database

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

func New(url string) (*sql.DB, error) {
	return sql.Open("libsql", url)
}
