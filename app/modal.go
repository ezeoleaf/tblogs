package app

import "github.com/rivo/tview"

const (
	blogsModalName = "blogsModal"
	resetModalName = "resetModal"
)

var blogsModal = tview.NewModal().
	SetText("Resize the window to see how the grid layout adapts").
	AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
	pages.HidePage(blogsModalName)
})

var resetModal = tview.NewModal().
	SetText("Are you sure to reset the app?").
	AddButtons([]string{"Yes", "No"}).
	SetDoneFunc(resetApp)
