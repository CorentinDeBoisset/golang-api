package service

import (
	"fmt"
	"regexp"
	"hash/crc32"
	"github.com/corentindeboisset/golang-api/app/migration"
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

// RunMigrations execute all the missing migrations on the server
func (m Migrator) RunMigrations() error {
	// Run everything in a transaction. In case of error, we can roll it back
	tx, err := m.Connection.Database.Begin()
	if err != nil {
		// Connection could not be started
		return err
	}

	// First check if the database db_migrations exists
	res := tx.QueryRow(`SELECT EXISTS(
		SELECT *
		FROM information_schema.tables
		WHERE
			table_schema = 'public' AND
			table_name = 'db_migrations'
	)`)

	var migTablePresent bool
	err = res.Scan(&migTablePresent)
	if err != nil {
		// result was invalid
		tx.Rollback()
		return err
	}

	alreadyRunMigrations := make(map[string]bool)
	if !migTablePresent {
		_, err = tx.Query(`
			CREATE TABLE db_migrations (version VARCHAR(50) NOT NULL, executed_at TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY(version))
		`)
		if err != nil {
			// could not create db_migration table
			tx.Rollback()
			return err
		}
	} else {
		versionRows, err := tx.Query(`
			SELECT version FROM db_migrations
		`)
		if err != nil {
			// could not fetch the list of executed migrations
			tx.Rollback()
			return err
		}
		for versionRows.Next() {
			var version string
			err = versionRows.Scan(&version)
			if err != nil {
				// A version number could not be parsed
				tx.Rollback()
				return err
			}

			alreadyRunMigrations[version] = true
		}
	}

	availableMigrations, err := m.checkAvailableMigrations()
	if err != nil {
		tx.Rollback()
		return err
	}

	var migrationsToRun []string
	for version := range availableMigrations {
		if _, ok := alreadyRunMigrations[version]; !ok {
			migrationsToRun = append(migrationsToRun, version)
		}
	}
	for version := range alreadyRunMigrations {
		if _, ok := availableMigrations[version]; !ok {
			// Warn there is a present migration with no corresponding file
		}
	}

	for _, version := range migrationsToRun {
		migrationByteContent, err := migration.Asset(fmt.Sprintf("%s_up.sql", version))
		if err != nil {
			tx.Rollback()
			return err
		}
		migrationContent := string(migrationByteContent)

		_, err = tx.Query(migrationContent)
		if err != nil {
			// There was an error running the migration
			tx.Rollback()
			return err
		}
		_, err = tx.Query(`INSERT INTO db_migrations (version) VALUES ($1)`, version)
		if err != nil {
			// There was an error running the migration
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (m Migrator) checkAvailableMigrations() (map[string]bool, error) {
	// list all present migrations, check if there is both up() and down() every time
	migrationVersionRegexp, err := regexp.Compile("^(\\d+)_(up|down)\\.sql$")
	if err != nil {
		// could not compile regex
		return nil, fmt.Errorf("Migration version regexp is invalid")
	}

	upMigrations := make(map[string]bool)
	downMigrations := make(map[string]bool)
	for _, sqlName := range migration.AssetNames() {
		submatches := migrationVersionRegexp.FindStringSubmatch(sqlName)
		if submatches == nil {
			// There was a non-compliant migration
			return nil, fmt.Errorf("Migration \"%s\" does not match the expected format", sqlName)
		}
		if submatches[2] == "up" {
			upMigrations[submatches[1]] = true
		} else {
			downMigrations[submatches[1]] = true
		}
	}
	for k := range upMigrations {
		if _, ok := downMigrations[k]; !ok {
			// there is an error in the migrations
			return nil, fmt.Errorf("Up migration \"%s\" has no matching down migration", k)
		}
	}
	for k := range downMigrations {
		if _, ok := upMigrations[k]; !ok {
			// there is an error in the migrations
			return nil, fmt.Errorf("Down migration \"%s\" has no matching up migration", k)
		}
	}

	return upMigrations, nil
}
