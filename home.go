package main

import (
	"github.com/ezeoleaf/tblogs/api"
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/helpers"
	"github.com/gdamore/tcell"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func Home(nextSlide func()) (title string, content tview.Primitive) {
	listHome := getList()

	appCfg := cfg.GetAPPConfig()

	if len(appCfg.FollowingBlogs) == 0 {
		listHome.AddItem("You're not following blogs", "Try Ctrl+B", ' ', nil)
	} else {
		listHome.Clear()

		posts := api.GetPosts(appCfg.FollowingBlogs)

		for _, post := range posts.Posts {
			r := ' '
			isIn, _ := helpers.IsHash(post.Hash, appCfg.SavedPosts)
			if isIn {
				r = 's'
			}
			listHome.AddItem(post.Title, post.Blog+" - "+post.Published, r, func() {
				return
			})
		}

		listHome.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
			post := posts.Posts[x]
			browser.OpenURL(post.Link)
		})

		listHome.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlS {
				appCfg := cfg.GetAPPConfig()

				x := listHome.GetCurrentItem()

				post := posts.Posts[x]

				r := ' '
				isIn, ix := helpers.IsHash(post.Hash, appCfg.SavedPosts)
				if !isIn {
					r = 's'
					appCfg.SavedPosts = append(appCfg.SavedPosts, post)
					cfg.UpdateAppConfig(appCfg)
				} else {
					appCfg.SavedPosts = append(appCfg.SavedPosts[:ix], appCfg.SavedPosts[ix+1:]...)
					cfg.UpdateAppConfig(appCfg)
				}

				listHome.RemoveItem(x)
				listHome.InsertItem(x, post.Title, post.Blog+" - "+post.Published, r, func() {
					return
				})
				listHome.SetCurrentItem(x)
				generateSavedPosts()
				return nil
			}
			return event
		})
	}

	return "Home", tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listHome, 0, 1, true), 0, 1, true)
}
