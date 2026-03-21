package gui

import (
	"context"
	"database/sql"
	"os"

	"fyne.io/fyne/v2/container"

	"github.com/Vikuuu/invoice_generator/internal/database"
)

type Config struct {
	Db                *sql.DB
	Cont              *container.Split
	Cwd               string
	Queries           *database.Queries
	Context           context.Context
	TypstBinPath      string
	ApplicationPath   string
	TypstTemplatePath []byte
}

func NewConfig() *Config {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Config{Cwd: cwd}
}
