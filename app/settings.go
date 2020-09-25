package app

import (
	"strconv"

	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/rivo/tview"
)

func settingsPage(nextSlide func()) (title string, content tview.Primitive) {
	appCfg := cfg.GetAPPConfig()

	pages = tview.NewPages()

	settingPage := tview.NewFlex()

	noTextAccept := func(s string, r rune) bool {
		return false
	}

	f := tview.NewForm().
		AddInputField("Following blogs:", strconv.Itoa(len(appCfg.FollowingBlogs)), 5, noTextAccept, nil).
		AddInputField("Saved posts:", strconv.Itoa(len(appCfg.SavedPosts)), 5, noTextAccept, nil).
		AddButton("Reset", func() {
			pages.ShowPage(resetModalName)
		})
	f.SetBorder(false).SetTitle("Settings")

	settingPage.AddItem(f, 0, 1, true)

	pages.AddPage("settings", settingPage, true, true).AddPage(resetModalName, resetModal, true, false)

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
		pages.HidePage(resetModalName)
	}
	pages.HidePage(resetModalName)
}
