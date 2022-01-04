package app

import (
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/data"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type slide func(nextSlide func()) (title string, content tview.Primitive)

// var app = tview.NewApplication()

// App contains the tview application and the layout for the display
type Tblogs struct {
	app         *tview.Application
	layout      *tview.Flex
	dataService data.Service
}

func NewApp(dataService data.Service) Tblogs {
	cfg.Setup()

	a := Tblogs{dataService: dataService}

	a.setLayout()

	app := tview.NewApplication()

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch ek := event.Key(); ek {
		case tcell.KeyCtrlH:
			goToSection(helpSection)
		case tcell.KeyCtrlB:
			goToSection(blogsSection)
		case tcell.KeyCtrlT:
			goToSection(homeSection)
		case tcell.KeyCtrlP:
			goToSection(savedPostsSection)
		case tcell.KeyCtrlX:
			goToSection(settingsSection)
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
