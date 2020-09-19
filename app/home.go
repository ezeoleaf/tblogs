package app

import (
	"github.com/ezeoleaf/tblogs/api"
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/helpers"
	"github.com/gdamore/tcell"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

var listHome *tview.List
var lastMode bool

func generateHomeList() {
	appCfg := cfg.GetAPPConfig()

	if len(appCfg.FollowingBlogs) == 0 {
		listHome.AddItem("You're not following blogs", "Try follow a blog with Ctrl+S from the Blogs (Ctrl+B) page", ' ', nil)
	} else {
		listHome.Clear()

		posts := api.GetPosts(appCfg.FollowingBlogs)

		for _, post := range posts.Posts {
			if lastMode {
				if post.PublishedAt == nil {
					continue
				}
				if post.PublishedAt.Before(appCfg.LastLogin) {
					continue
				}
			}
			r := emptyRune
			isIn, _ := helpers.IsHash(post.Hash, appCfg.SavedPosts)
			if isIn {
				r = savedRune
			}
			listHome.AddItem(post.Title, post.Blog+" - "+post.Published, r, emptyFunc)
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

				r := emptyRune
				isIn, ix := helpers.IsHash(post.Hash, appCfg.SavedPosts)
				if !isIn {
					r = savedRune
					appCfg.SavedPosts = append(appCfg.SavedPosts, post)
				} else {
					appCfg.SavedPosts = append(appCfg.SavedPosts[:ix], appCfg.SavedPosts[ix+1:]...)
				}
				cfg.UpdateAppConfig(appCfg)

				updateItemList(listHome, x, post.Title, post.Blog+" - "+post.Published, r, emptyFunc)
				generateSavedPosts()
				return nil
			} else if event.Key() == tcell.KeyCtrlR {
				listHome.Clear()
				generateHomeList()
			} else if event.Key() == tcell.KeyCtrlF {
				//TODO: Search posts
			} else if event.Key() == tcell.KeyCtrlL { //Only displays the last ones
				lastMode = !lastMode
				generateHomeList()
			}
			return event
		})
	}
}

func homePage(nextSlide func()) (title string, content tview.Primitive) {
	listHome = getList()

	generateHomeList()

	return "Home", tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listHome, 0, 1, true), 0, 1, true)
}
