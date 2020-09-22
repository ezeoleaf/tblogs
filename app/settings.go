package app

import (
	"github.com/rivo/tview"
)

func settingsPage(nextSlide func()) (title string, content tview.Primitive) {

	settingPage := tview.NewFlex()

	f := tview.NewForm().
		AddInputField("First name:", "", 20, nil, nil).
		AddInputField("Last name:", "", 20, nil, nil).
		AddDropDown("Role:", []string{"Engineer", "Manager", "Administration"}, 0, nil).
		AddCheckbox("On vacation:", false, nil).
		AddPasswordField("Password:", "", 10, '*', nil).
		AddButton("Reset", nextSlide)
	f.SetBorder(false).SetTitle("Settings")

	settingPage.AddItem(f, 0, 1, true)
	// return "Forms", Code(f, 36, 15, form)
	// lines := strings.Split(logo, "\n")
	// logoWidth := 0
	// logoHeight := len(lines)
	// for _, line := range lines {
	// 	if len(line) > logoWidth {
	// 		logoWidth = len(line)
	// 	}
	// }
	// logoBox := tview.NewTextView().
	// 	SetTextColor(tcell.ColorGreen).
	// 	SetDoneFunc(func(key tcell.Key) {
	// 		nextSlide()
	// 	})
	// fmt.Fprint(logoBox, logo)

	// // Create a frame for the subtitle and navigation infos.
	// frame := tview.NewFrame(tview.NewBox()).
	// 	SetBorders(0, 0, 0, 0, 0, 0).
	// 	AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
	// 	AddText("", true, tview.AlignCenter, tcell.ColorWhite).
	// 	AddText(quote, true, tview.AlignCenter, tcell.ColorWhite).
	// 	AddText("", true, tview.AlignCenter, tcell.ColorWhite)

	// for _, s := range shortcuts {
	// 	frame.AddText(s, true, tview.AlignCenter, tcell.ColorTeal)
	// }

	// // Create a Flex layout that centers the logo and subtitle.
	// flex := tview.NewFlex().
	// 	SetDirection(tview.FlexRow).
	// 	AddItem(tview.NewBox(), 0, 1, false).
	// 	AddItem(tview.NewFlex().
	// 		AddItem(tview.NewBox(), 0, 1, false).
	// 		AddItem(logoBox, logoWidth, 1, true).
	// 		AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
	// 	AddItem(frame, 0, 10, false)

	return settingsSection, settingPage
}
