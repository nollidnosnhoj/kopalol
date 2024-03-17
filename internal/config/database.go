package config

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/spf13/viper"
)

func NewDatabaseConfig() (*sql.DB, error) {
	databaseUrl := viper.GetString("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("DATABASE_URL is required")
	}
	authToken := viper.GetString("DATABASE_AUTH_TOKEN")
	if authToken != "" {
		databaseUrl = fmt.Sprintf("%s?authToken=%s", databaseUrl, authToken)
	}
	return sql.Open("libsql", databaseUrl)
}
