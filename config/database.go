package config

import (
	"errors"
	"fmt"

	database "github.com/nollidnosnhoj/kopalol/db"
	_ "github.com/tursodatabase/go-libsql"

	"github.com/spf13/viper"
)

func NewDatabaseWithConfig() (*database.Database, error) {
	databaseUrl := viper.GetString("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("DATABASE_URL is required")
	}
	authToken := viper.GetString("DATABASE_AUTH_TOKEN")
	if authToken != "" {
		databaseUrl = fmt.Sprintf("%s?authToken=%s", databaseUrl, authToken)
	}
	return database.New(databaseUrl)
}
