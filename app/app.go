package app

import (
	"fmt"
	"strconv"

	"github.com/ezeoleaf/tblogs/cfg"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type slide func(nextSlide func()) (title string, content tview.Primitive)

var app = tview.NewApplication()

// App contains the tview application and the layout for the display
type App struct {
	App    *tview.Application
	Layout *tview.Flex
}

// Setup returns an instance of the application
func Setup() App {

	cfg.Setup()

	appCfg := cfg.GetAPPConfig()

	// The presentation slides.
	slides := []slide{
		homePage,
		savedPostsPage,
		blogsPage,
		settingsPage,
	}

	highlight := "Home"

	if appCfg.FirstUse {
		slides = append([]slide{helpPage}, slides...)
		appCfg.FirstUse = false
		highlight = "Help"
		cfg.UpdateAppConfig(appCfg)
	} else {
		slides = append(slides, helpPage)
	}

	pages := tview.NewPages()

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	goToHelp := func() {
		info.Highlight("Help").
			ScrollToHighlight()
	}
	goToBlogs := func() {
		info.Highlight("Blogs").
			ScrollToHighlight()
	}
	goToHome := func() {
		info.Highlight("Home").
			ScrollToHighlight()
	}
	goToSavedPosts := func() {
		info.Highlight("Saved Posts").
			ScrollToHighlight()
	}
	goToSettings := func() {
		info.Highlight("Saved Posts").
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		pages.AddPage(title, primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%s"][darkcyan]%s[white][""]  `, index+1, title, title)
	}
	info.Highlight(highlight)

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlH {
			goToHelp()
			return nil
		} else if event.Key() == tcell.KeyCtrlB {
			goToBlogs()
			return nil
		} else if event.Key() == tcell.KeyCtrlT {
			goToHome()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			goToSavedPosts()
			return nil
		} else if event.Key() == tcell.KeyCtrlX {
			goToSettings()
			return nil
		}
		return event
	})

	a := App{App: app, Layout: layout}

	return a
}

// Start launches the app
func (a App) Start() {
	if err := a.App.SetRoot(a.Layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
