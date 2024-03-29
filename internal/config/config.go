package config

import (
	"context"
	"log/slog"
	"os"

	database "github.com/nollidnosnhoj/kopalol/internal/db"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
	"github.com/spf13/viper"
)

// Container that contains all the dependencies for the application based on configuration
type Container struct {
	db      *database.Database
	storage storage.Storage
	logger  *slog.Logger
}

func NewContainer(context context.Context) (*Container, error) {
	db, err := NewDatabaseWithConfig()
	if err != nil {
		return nil, err
	}
	s, err := NewStorageWithConfig(context)
	if err != nil {
		return nil, err
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Container{
		db:      db,
		storage: s,
		logger:  l,
	}, nil
}

func (c *Container) Close() {
	c.db.Close()
}

func (c *Container) Database() *database.Database {
	return c.db
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
