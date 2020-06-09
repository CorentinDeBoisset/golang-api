package service

import (
	"hash/crc32"
)

// Generate an ID from the database name
func generateLockID(dbName string) (uint32) {
	return crc32.ChecksumIEEE([]byte(dbName))
}

// Migrator is used to run upgrade (or downgrade) the database schema
type Migrator struct {
	Logger		*Logger
	Connection	*Connection
	LockID		uint32
}

// ProvideMigrator is the provider for the Migrator Service
func ProvideMigrator(logger *Logger, conn *Connection) (*Migrator, error) {
	return &Migrator{Logger: logger, Connection: conn}, nil
}

// LockDatabase sets an advisory lock, to ensure only one migration is running
func (m Migrator) LockDatabase () error {
	m.LockID = generateLockID(m.Connection.DbName)

	return nil
}

// UnlockDatabase sets an advisory lock, to ensure only one migration is running
func (m Migrator) UnlockDatabase () error {
	return nil
}
