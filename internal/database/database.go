package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nollidnosnhoj/vgpx/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewDatabase(config *config.Config) (*bun.DB, error) {
	sqldb, err := sql.Open("pgx", config.DATABASE_URL)
	if err != nil {
		return nil, err
	}
	return bun.NewDB(sqldb, pgdialect.New()), nil
}
