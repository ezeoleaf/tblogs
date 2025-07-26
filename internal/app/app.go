package app

import (
	"log"
	"sync"
	"time"

	"github.com/ezeoleaf/tblogs/internal/config"
	"github.com/gdamore/tcell/v2"
	"github.com/mmcdole/gofeed"
	"github.com/rivo/tview"
)

// App contains the tview application, the layout for the display, and the loaded config
type App struct {
	TApp             *tview.Application
	TLayout          *tview.Flex
	Config           *config.Config
	Cache            *Cache
	viewsList        map[string]*tview.List
	currentBlogs     []config.Blog
	currentHomePosts []PostWithMeta
}

type Cache struct {
	sync.RWMutex
	data map[string]Feed
	ttl  time.Duration
}

type Feed struct {
	Post      *gofeed.Feed
	Timestamp time.Time
}

// NewApp returns an instance of the application, initialized with the provided config
func NewApp(cfg *config.Config) *App {
	app := &App{
		TApp:   tview.NewApplication(),
		Config: cfg,
		Cache: &Cache{
			data: make(map[string]Feed),
			ttl:  1 * time.Hour,
		},
		viewsList:    make(map[string]*tview.List),
		currentBlogs: cfg.Blogs,
	}

	pages, info := app.getPagesInfo()

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	app.TApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch ek := event.Key(); ek {
		case tcell.KeyCtrlH:
			app.goToSection(helpSection, info)
		case tcell.KeyCtrlB:
			app.goToSection(blogsSection, info)
			if app.viewsList["posts"] != nil {
				app.viewsList["posts"].Clear()
			}
		case tcell.KeyCtrlT:
			app.goToSection(homeSection, info)
		case tcell.KeyCtrlP:
			app.goToSection(savedPostsSection, info)
		}
		return event
	})

	app.TLayout = layout

	return app
}

func (a *App) Run() {
	if err := a.TApp.SetRoot(a.TLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	defer func() {
		a.Config.App.LastLogin = a.Config.App.CurrentLogin
		err := config.SaveConfig(a.Config)
		if err != nil {
			log.Printf("failed to save config: %v", err)
		}
	}()
}
