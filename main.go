package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Slide func(nextSlide func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

func main() {
	// The presentation slides.
	slides := []Slide{
		Home,
		Blogs,
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

	goToHome := func() {
		info.Highlight("0").
			ScrollToHighlight()
	}
	goToBlogs := func() {
		info.Highlight("1").
			ScrollToHighlight()
	}
	goToGrid := func() {
		info.Highlight("2").
			ScrollToHighlight()
	}
	goToTable := func() {
		info.Highlight("2").
			ScrollToHighlight()
	}
	// Create the pages for all slides.
	// previousSlide := func() {
	// 	slide, _ := strconv.Atoi(info.GetHighlights()[0])
	// 	slide = (slide - 1 + len(slides)) % len(slides)
	// 	fmt.Println(slide)
	// 	info.Highlight(strconv.Itoa(slide)).
	// 		ScrollToHighlight()
	// }
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}
	info.Highlight("0")

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// if event.Key() == tcell.KeyEnter {
		// 	return nil
		// }
		if event.Key() == tcell.KeyCtrlH {
			goToHome()
			return nil
		} else if event.Key() == tcell.KeyCtrlB {
			goToBlogs()
			return nil
		} else if event.Key() == tcell.KeyCtrlG {
			goToGrid()
			goToTable()
			return nil
		}
		return event
	})

	// app.SetMouseCapture(func(event *tcell.EventMouse, mouseAction tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
	// 	if mouseAction == tview.MouseLeftClick {
	// 		return nil, mouseAction
	// 	}
	// 	return event, mouseAction
	// })

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
