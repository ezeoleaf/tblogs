package app

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

const (
	homeSection       = "Home"
	helpSection       = "Help"
	blogsSection      = "Blogs"
	savedPostsSection = "Saved Posts"
	xSection          = "X"
)

func (a *App) getPagesInfo() (*tview.Pages, *tview.TextView) {
	slides := []func(func()) (string, tview.Primitive){
		a.homePage,
		a.savedPostsPage,
		a.blogsPage,
		a.xPage,
		a.helpPage,
	}

	pages := tview.NewPages()

	info := tview.NewTextView().
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
		_, _ = fmt.Fprintf(info, `%d ["%s"][darkcyan]%s[white][""]  `, index+1, title, title)
	}
	info.Highlight(homeSection)

	return pages, info
}

func (a *App) goToSection(section string, info *tview.TextView) {
	info.Highlight(section).
		ScrollToHighlight()
}
