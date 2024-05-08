package assets

import (
	"embed"
)

//go:embed migrations/*.sql
var Migrations embed.FS

//
//func main() {
//	var db *sql.DB
//	// setup database
//
//	goose.SetBaseFS(embedMigrations)
//
//	if err := goose.SetDialect("postgres"); err != nil {
//		panic(err)
//	}
//
//	if err := goose.Up(db, "migrations"); err != nil {
//		panic(err)
//	}
//
//	// run app
//}
