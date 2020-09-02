package main

import (
	"fmt"
	"strconv"

	"github.com/ezeoleaf/tblogs/cfg"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Slide func(nextSlide func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

func main() {
	cfg.Setup()

	appCfg := cfg.GetAPPConfig()

	// The presentation slides.
	slides := []Slide{
		Home,
		Blogs,
	}

	if appCfg.FirstUse {
		slides = append([]Slide{Help}, slides...)
		appCfg.FirstUse = false
		cfg.UpdateAppConfig(appCfg)
	} else {
		slides = append(slides, Help)
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
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		fmt.Println(slide)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		pages.AddPage(title, primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}
	info.Highlight("Home")

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
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
