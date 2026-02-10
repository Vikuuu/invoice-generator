package main

import (
	"database/sql"
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed internal/migrations/*.sql
var embedMigrations embed.FS

func migrate(db *sql.DB) {
	// setup database
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "internal/migrations"); err != nil {
		panic(err)
	}
	log.Println("Up Migration Done")
}

func setUpDatabase(cwd string) *sql.DB {
	err := os.MkdirAll(filepath.Join(cwd, "db"), 0o755)
	if err != nil {
		panic(err)
	}

	db := newDatabase("")
	migrate(db)

	return db
}

func newDatabase(dbPath string) *sql.DB {
	if dbPath == "" {
		dbPath = "./db/my.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	var sqliteVersion string
	err = db.QueryRow("SELECT sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		panic(err)
	}

	log.Println("DB connected successfully, version: ", sqliteVersion)
	return db
}
