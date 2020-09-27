package app

import (
	"strconv"

	"github.com/ezeoleaf/tblogs/cfg"
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

	settingsComponent.AddItem(formComponent, 0, 1, true)

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
