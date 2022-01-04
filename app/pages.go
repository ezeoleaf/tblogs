package app

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

var info *tview.TextView

// var pages *tview.Pages

const (
	homeSection       = "Home"
	helpSection       = "Help"
	blogsSection      = "Blogs"
	savedPostsSection = "Saved Posts"
	settingsSection   = "Settings"
)

func (a *Tblogs) setLayout() {
	// appCfg := cfg.GetAPPConfig()
	// The presentation slides.
	slides := []slide{
		// homePage,
		// savedPostsPage,
		a.setBlogsPage,
		// settingsPage,
		helpPage,
	}

	pages := tview.NewPages()

	// The bottom row has some info on where we are.
	info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

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
	info.Highlight(blogsSection)

	// Create the main layout.
	a.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

}

func goToSection(section string) {
	info.Highlight(section).
		ScrollToHighlight()
}
