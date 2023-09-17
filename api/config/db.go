package config

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrations struct {
	Migrate *migrate.Migrate
}

func CreateMigrations(db *sql.DB, migrationPath string) (migrations Migrations, err error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, "mysql", driver)

	migrations = Migrations{Migrate: m}
	return
}

func (m *Migrations) MigrateUp(ignoreNoChange bool) (err error) {
	err = m.Migrate.Up()
	if ignoreNoChange && err == migrate.ErrNoChange {
		err = nil
	}
	return
}

func (m *Migrations) MigrateDown(ignoreNoChange bool) (err error) {
	err = m.Migrate.Down()
	if ignoreNoChange && err == migrate.ErrNoChange {
		err = nil
	}
	return
}

func InitDb(dbUrl string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dbUrl)
	return
}
