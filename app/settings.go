package app

import (
	"strconv"

	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/rivo/tview"
)

func settingsPage(nextSlide func()) (title string, content tview.Primitive) {
	appCfg := cfg.GetAPPConfig()

	settingPage := tview.NewFlex()

	noTextAccept := func(s string, r rune) bool {
		return false
	}

	f := tview.NewForm().
		AddInputField("Following blogs:", strconv.Itoa(len(appCfg.FollowingBlogs)), 5, noTextAccept, nil).
		AddInputField("Saved posts:", strconv.Itoa(len(appCfg.SavedPosts)), 5, noTextAccept, nil).
		AddButton("Reset", nextSlide)
	f.SetBorder(false).SetTitle("Settings")

	settingPage.AddItem(f, 0, 1, true)

	return settingsSection, settingPage
}
