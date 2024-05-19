package main

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

func applyMigrations(connStr string, migrationFS embed.FS) error {
	var db *sql.DB
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	goose.SetBaseFS(migrationFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}
	return nil
}
