package app

import (
	"log"

	"github.com/ezeoleaf/tblogs/internal/config"
	"github.com/gdamore/tcell/v2"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func (a *App) generateSavedPosts(listSavedPosts *tview.List) {
	appCfg := a.Config.App
	listSavedPosts.Clear()

	if len(appCfg.SavedPosts) == 0 {
		listSavedPosts.AddItem("You don't have saved posts", "Try save them using Ctrl+S from the Blogs (Ctrl+B) or the Home (Ctrl+T) pages", ' ', nil)
	} else {

		posts := appCfg.SavedPosts

		for _, post := range posts {
			published := ""
			if post.Published != nil {
				published = post.Published.Format("2006-01-02 15:04")
			}
			link := post.Link
			itemLink := link // capture for closure
			listSavedPosts.AddItem(post.Title, published, savedRune, func() {
				err := browser.OpenURL(itemLink)
				if err != nil {
					log.Printf("failed to open URL: %v", err)
				}
			})
		}

		listSavedPosts.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlD {
				x := listSavedPosts.GetCurrentItem()

				if x < 0 || x >= len(a.Config.App.SavedPosts) {
					return nil
				}
				// // Remove the post at index x
				a.Config.App.SavedPosts = append(a.Config.App.SavedPosts[:x], a.Config.App.SavedPosts[x+1:]...)
				// Persist the config change
				err := config.SaveConfig(a.Config)
				if err != nil {
					log.Printf("failed to save config: %v", err)
					return nil
				}
				// // Refresh the saved posts list
				a.generateSavedPosts(listSavedPosts)
				a.generateHomeList(a.viewsList["home"])
				return nil
			}
			return event
		})
	}
}

func (a *App) savedPostsPage(nextSlide func()) (title string, content tview.Primitive) {
	listSavedPosts := a.viewsList["savedPosts"]
	if listSavedPosts == nil {
		listSavedPosts = getList()
		a.viewsList["savedPosts"] = listSavedPosts
		a.generateSavedPosts(listSavedPosts)
	}

	return savedPostsSection, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listSavedPosts, 0, 1, true), 0, 1, true)
}
