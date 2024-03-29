package container

import (
	"log/slog"

	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/db"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
)

type Container struct {
	Storage storage.Storage
	Db      *db.Database
	Logger  *slog.Logger
	Config  *config.Configuration
}
