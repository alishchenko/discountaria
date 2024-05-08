package cli

import (
	"database/sql"
	"github.com/alishchenko/discountaria/internal/assets"
	"github.com/alishchenko/discountaria/internal/config"
	"github.com/pressly/goose/v3"
)

func MigrateUp(cfg config.Config) error {
	db, err := sql.Open("postgres", cfg.DB.Url)

	if err != nil {
		return err
	}
	goose.SetBaseFS(assets.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}

func MigrateDown(cfg config.Config) error {
	db, err := sql.Open("postgres", cfg.DB.Url)

	if err != nil {
		return err
	}
	goose.SetBaseFS(assets.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Down(db, "migrations"); err != nil {
		return err
	}

	return nil
}
