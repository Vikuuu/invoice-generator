package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/Vikuuu/invoice_generator/assets"
	"github.com/Vikuuu/invoice_generator/internal/database"
	gui "github.com/Vikuuu/invoice_generator/internal/gui"
	"github.com/Vikuuu/invoice_generator/internal/utils"
)

var (
	topWindow fyne.Window
	logger    *slog.Logger
)

var ErrUnsupportedPlatform = errors.New("unsupported platform")

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	a := app.NewWithID("io.invoice.go")
	logLifeCycle(a)
	w := a.NewWindow("Invoice Generator")
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()
	w.Resize(fyne.NewSize(640, 460))

	cfg := gui.NewConfig()

	// Create the folder to store applications data
	path, err := createAppDir()
	if err != nil {
		slog.Error("Setup: App Folder", "msg", err)
	}
	cfg.ApplicationPath = path
	cfg.TypstTemplatePath = embedTypstTemplate

	err = createAppAssetsDir()
	if err != nil {
		slog.Error("Setup: App assets folder", "msg", err)
	}

	typstBinPath, err := assets.Init()
	if err != nil {
		slog.Error("Setup: Typst", "msg", err)
	}
	// Clean up the temp for the Typst binaries.
	defer assets.Cleanup()
	cfg.TypstBinPath = typstBinPath

	db := setUpDatabase(cfg.ApplicationPath)
	queries := database.New(db)

	cfg.Db, cfg.Queries = db, queries

	// TODO: Use proper context
	cfg.Context = context.Background()

	w.SetCloseIntercept(func() {
		slog.Info("Window closing, cleaning up...")
		db.Close()
		w.Close()
		a.Quit()
	})

	cfg.GreatingPage(a, w)
	w.Show()
	a.Run()
}

func logLifeCycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		slog.Info("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		slog.Info("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		slog.Info("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		slog.Info("Lifecycle: Exited Foreground")
	})
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)

	file := fyne.NewMenu("File", newItem)

	about := fyne.NewMenuItem("About", nil)
	help := fyne.NewMenu("Help", about)

	main := fyne.NewMainMenu(file, help)
	return main
}

func createAppDir() (string, error) {
	var basePath string
	var err error
	switch runtime.GOOS {
	case "linux":
		basePath, err = utils.GetLinuxAppPath()
	case "windows":
		basePath, err = utils.GetWinAppPath()
	case "darwin":
		basePath, err = utils.GetDarwinAppPath()
	default:
		return "", ErrUnsupportedPlatform
	}

	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return "", err
	}
	return basePath, nil
}

func createAppAssetsDir() error {
	path, err := utils.GetAssetsAppPath()
	if err != nil {
		return err
	}

	if err = os.MkdirAll(path, 0o755); err != nil {
		return err
	}
	return nil
}
