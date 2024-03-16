package database

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/nollidnosnhoj/vgpx/internal/config"
)

func Open(cfg *config.Config) (*sql.DB, error) {
	databaseUrl := cfg.DATABASE_URL
	if cfg.DATABASE_AUTH_TOKEN != "" {
		databaseUrl = fmt.Sprintf("%s?authToken=%s", databaseUrl, cfg.DATABASE_AUTH_TOKEN)
	}
	return sql.Open("libsql", databaseUrl)
}
