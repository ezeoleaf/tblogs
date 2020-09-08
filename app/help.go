package app

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const logo = `
 ___________ _                 
|_   _| ___ \ |                
  | | | |_/ / | ___   __ _ ___ 
  | | | ___ \ |/ _ \ / _  / __|
  | | | |_/ / | (_) | (_| \__ \
  \_/ \____/|_|\___/ \__, |___/
                      __/ |    
                     |___/     
`

const (
	subtitle = `tblogs - Read development blogs from the terminal`
	quote    = `All your blogs are belongs to us`
	mouse    = `(or use your mouse)`
)

var shortcuts = []string{"Ctrl-S: Save Blog or Post", "Ctrl-F: Search Blogs or Posts", "Ctrl-R: Reload posts in Home page",
	"Ctrl-T: Move to Home page", "Ctrl-B: Move to Blogs page",
	"Ctrl-H: Move to Help page", "Ctrl-P: Move to Saved Posts page",
	"Ctrl-C: Close TBlogs", "Enter: Search posts for blog or open browser with post",
	"Esc: Returns from posts results to blogs list"}

func helpPage(nextSlide func()) (title string, content tview.Primitive) {
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(quote, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite)

	for _, s := range shortcuts {
		frame.AddText(s, true, tview.AlignCenter, tcell.ColorTeal)
	}

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Help", flex
}
