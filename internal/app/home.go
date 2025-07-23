package app

import (
	"log"
	"slices"

	"github.com/ezeoleaf/tblogs/internal"
	"github.com/gdamore/tcell/v2"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func (a *App) generateHomeList(listHome *tview.List) {
	appCfg := a.Config.App

	if len(appCfg.FollowingBlogs) == 0 {
		listHome.AddItem("Welcome! You're not following any blog", "Try follow a blog with Ctrl+S from the Blogs (Ctrl+B) page", ' ', nil)
	} else {
		listHome.Clear()

		for _, blogName := range a.Config.App.FollowingBlogs {
			ix := slices.Index(a.Config.App.FollowingBlogs, blogName)
			blog := a.Config.Blogs[ix]
			feed, err := a.getFeed(blog.Feed)
			if err != nil {
				listHome.AddItem("Failed to load feed", err.Error(), 0, nil)
				continue
			}
			for _, item := range feed.Items {
				postHash := internal.GetHash([]string{item.Title, item.Content, item.Link})
				isIn := false
				for _, post := range a.Config.App.SavedPosts {
					if post.Hash == postHash {
						isIn = true
						break
					}
				}
				r := emptyRune
				if isIn {
					r = savedRune
				}

				link := item.Link
				itemLink := link
				listHome.AddItem(item.Title, feed.Title+" - "+item.PublishedParsed.Format("2006-01-02 15:04"), r, func() {
					err := browser.OpenURL(itemLink)
					if err != nil {
						log.Printf("failed to open URL: %v", err)
					}
				})
			}
		}

		listHome.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			return event
		})
	}
}

func (a *App) homePage(nextSlide func()) (title string, content tview.Primitive) {
	listHome := a.viewsList["home"]
	if listHome == nil {
		listHome = getList()
		a.viewsList["home"] = listHome
		a.generateHomeList(listHome)
	}

	return homeSection, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listHome, 0, 1, true), 0, 1, true)
}
