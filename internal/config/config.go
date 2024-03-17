package config

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"github.com/nollidnosnhoj/kopalol/internal/queries"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
	"github.com/spf13/viper"
)

type Container struct {
	db      *sql.DB
	queries *queries.Queries
	storage storage.Storage
	logger  *slog.Logger
}

func NewContainer(context context.Context) (*Container, error) {
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}
	q := queries.New(db)
	s, err := NewStorageWithConfig(context)
	if err != nil {
		return nil, err
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Container{
		db:      db,
		queries: q,
		storage: s,
		logger:  l,
	}, nil
}

func (c *Container) Close() {
	c.db.Close()
}

func (c *Container) DB() *sql.DB {
	return c.db
}

func (c *Container) Queries() *queries.Queries {
	return c.queries
}

func (c *Container) Storage() storage.Storage {
	return c.storage
}

func (c *Container) Logger() *slog.Logger {
	return c.logger
}

func init() {
	viper.AutomaticEnv()
}
