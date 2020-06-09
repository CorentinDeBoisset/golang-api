package service

import (
	"fmt"
	"strings"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the postgreSQL driver
)

// Connection is the database client
type Connection struct {
	Logger		*Logger
	Database	*sqlx.DB
	DbName		string
}

// ProvideConnection returns a new database client
func ProvideConnection(logger *Logger) (*Connection, error) {
	dbName := strings.ReplaceAll(viper.GetString("storage.database"), "'", "\\'")
	connStr := fmt.Sprintf(
		"user='%s' password='%s' host='%s' port=%d dbname='%s' sslmode=disable",
		strings.ReplaceAll(viper.GetString("storage.username"), "'", "\\'"),
		strings.ReplaceAll(viper.GetString("storage.password"), "'", "\\'"),
		strings.ReplaceAll(viper.GetString("storage.host"), "'", "\\'"),
		viper.GetInt("storage.port"),
		dbName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, errors.WithMessage(err, "The connection to the database failed.")
	}

	// Test we can ping the database
	err = db.Ping()
	if err != nil {
		return nil, errors.WithMessage(err, "The database was reached but is unresponsive.")
	}

	// TODO debug-log the successful connection

	return &Connection{Logger: logger, DbName: dbName, Database: db}, nil
}
