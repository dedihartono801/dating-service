package database

import (
	"fmt"

	"dating-service/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB
}

func InitPostgres(config *config.Config) *Database {
	dsn := getDataSourceName(config)
	database, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db := &Database{
		database,
	}

	return db
}

func getDataSourceName(config *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseName,
	)
}
