package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// import migrate source driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (d *Database) MigrateUp() error {
	return d.migrateUpWithSourceURL("file://deployments/migrations")
}

func (d *Database) migrateUpWithSourceURL(sourceURL string) error {
	sql, err := d.db.DB()
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(sql, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "mysql", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migration: %w", err)
	}

	log.Println("Migrations completed successfully")

	return nil
}
