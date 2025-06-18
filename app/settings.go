package app

import (
	"strconv"

	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var settingsComponent *tview.Flex
var formComponent *tview.Form

func generateSettingsPage() {
	appCfg := cfg.GetAPPConfig()

	settingsComponent = tview.NewFlex()

	noTextAccept := func(s string, r rune) bool {
		return false
	}

	formComponent.Clear(true)

	formComponent.AddInputField("Following blogs:", strconv.Itoa(len(appCfg.FollowingBlogs)), 5, noTextAccept, nil).
		AddInputField("Saved posts:", strconv.Itoa(len(appCfg.SavedPosts)), 5, noTextAccept, nil).
		AddButton("Reset", func() {
			pages.ShowPage(resetModalName)
		}).
		SetBorder(false).SetTitle("Settings")

	table := tview.NewTable().
		SetFixed(1, 1)

	// Header
	for column, cell := range []string{"ID", "Word", "Posts filtered"} {
		color := tcell.ColorWhite
		align := tview.AlignCenter
		tableCell := tview.NewTableCell(cell).
			SetTextColor(color).
			SetAlign(align).
			SetSelectable(false)
		if column >= 1 && column <= 3 {
			tableCell.SetExpansion(1)
		}
		table.SetCell(0, column, tableCell)
	}

	table.SetBorder(false).SetTitle("Filtered words")

	code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

	settingsComponent.AddItem(tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(formComponent, 10, 1, true).
		AddItem(table, 20, 0, false), 0, 1, true).
		AddItem(code, 56, 1, false)
	// AddItem(formComponent, 0, 1, true).
	// 	AddItem(formComponent, 0, 1, false)

}

func settingsPage(nextSlide func()) (title string, content tview.Primitive) {
	pages = tview.NewPages()
	formComponent = tview.NewForm()
	generateSettingsPage()

	pages.AddPage("settings", settingsComponent, true, true).AddPage(resetModalName, resetModal, true, false)

	return settingsSection, pages
}

func resetApp(buttonIndex int, buttonLabel string) {
	if buttonLabel == "Yes" {
		cfg.ResetAPPConfig()
		listHome.Clear()
		listBlogs.Clear()
		listPosts.Clear()
		generateHomeList()
		generateBlogsList()
		generateSettingsPage()
		pages.HidePage(resetModalName)
	}
	pages.HidePage(resetModalName)
}
