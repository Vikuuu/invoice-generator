package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "modernc.org/sqlite"
)

var topWindow fyne.Window

type config struct {
	db   *sql.DB
	cont *container.Split
}

func main() {
	a := app.NewWithID("io.invoice.go")
	makeTray(a)
	logLifeCycle(a)
	w := a.NewWindow("Invoice Generator")
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()
	w.Resize(fyne.NewSize(640, 460))

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(pwd, "db"), 0o755)
	if err != nil {
		panic(err)
	}

	db, err := newDatabase("")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cfg := &config{db: db}

	sidebar := cfg.sidebar(a, w)
	company := cfg.addNewCompany(a, w)
	cfg.cont = container.NewHSplit(sidebar, company)
	cfg.cont.Offset = 0.2
	w.SetContent(cfg.cont)
	w.ShowAndRun()
}

func logLifeCycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)

	file := fyne.NewMenu("File", newItem)

	main := fyne.NewMainMenu(file)
	return main
}

func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func (c *config) mainPage(a fyne.App, w fyne.Window) {
	sb := c.sidebar(a, w)
	form := c.addNewCompany(a, w)

	content := container.New(layout.NewHBoxLayout(), sb, form)

	w.SetContent(content)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}

var (
	sidebarData = map[string]func(*config, fyne.App, fyne.Window) *widget.Form{
		"add company": func(c *config, a fyne.App, w fyne.Window) *widget.Form {
			return c.addNewCompany(a, w)
		},
		"add payment method": func(c *config, a fyne.App, w fyne.Window) *widget.Form {
			return c.addNewPaymentMethod(a, w)
		},
	}

	sidebarKey = []string{
		"add company",
		"add payment method",
	}
)

func (c *config) sidebar(a fyne.App, w fyne.Window) *widget.List {
	list := widget.NewList(
		func() int { return len(sidebarKey) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Sidebar")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(sidebarKey[i])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		key := sidebarKey[id]
		if fn, ok := sidebarData[key]; ok {
			log.Println("Called: ", key)
			con := fn(c, a, w)
			c.cont.Trailing = con
			c.cont.Refresh()
		}
	}

	return list
}

func newDatabase(dbPath string) (*sql.DB, error) {
	if dbPath == "" {
		dbPath = "./db/my.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	var sqliteVersion string
	err = db.QueryRow("SELECT sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		return nil, err
	}

	log.Println("DB connected successfully, version: ", sqliteVersion)
	return db, nil
}
