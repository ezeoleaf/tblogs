package app

import (
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/data"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type slide func(nextSlide func()) (title string, content tview.Primitive)

// var app = tview.NewApplication()

type appData struct {
}

// App contains the tview application and the layout for the display
type Tblogs struct {
	app         *tview.Application
	layout      *tview.Flex
	dataService data.Service
	data        appData
}

func NewApp(dataService data.Service) Tblogs {
	cfg.Setup()

	a := Tblogs{dataService: dataService, data: appData{}}

	a.setLayout()

	app := tview.NewApplication()

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch ek := event.Key(); ek {
		case tcell.KeyCtrlA:
			goToSection(AboutSection)
		case tcell.KeyCtrlB:
			goToSection(BlogsSection)
		case tcell.KeyCtrlT:
			goToSection(TwitterSection)
			// goToSection(HomeSection)
		case tcell.KeyCtrlP:
			goToSection(SavedPostsSection)
		case tcell.KeyCtrlX:
			goToSection(SettingsSection)
		}
		return event
	})

	a.app = app

	return a
}

// Start launches the app
func (a Tblogs) Start() {
	if err := a.app.SetRoot(a.layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
