package app

import (
	"log"
	"slices"
	"strings"
	"time"

	"github.com/ezeoleaf/tblogs/internal"
	"github.com/ezeoleaf/tblogs/internal/config"
	"github.com/gdamore/tcell/v2"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func (a *App) generateBlogsList() (blogPage *tview.Flex, listBlogs *tview.List, listPosts *tview.List) {
	blogPage = tview.NewFlex()

	listBlogs = tview.NewList()
	listBlogs.ShowSecondaryText(false)

	if a.viewsList["posts"] != nil {
		a.viewsList["posts"].Clear()
	}

	listBlogs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlS {
			a.followBlogs(listBlogs)
			return nil
		} else if event.Key() == tcell.KeyCtrlF {
			a.showBlogSearchModal(listBlogs)
			return nil
		} else if event.Key() == tcell.KeyEsc {
			a.filterBlogsList(listBlogs, "")
			a.TApp.SetRoot(a.TLayout, true)
		}
		return event
	})

	listPosts = tview.NewList()
	listPosts.SetDoneFunc(func() {
		a.TApp.SetFocus(listBlogs)
	})

	a.viewsList["posts"] = listPosts

	a.filterBlogsList(listBlogs, "")

	blogPage.AddItem(listBlogs, 0, 1, true).
		AddItem(listPosts, 0, 1, false)

	return blogPage, listBlogs, listPosts
}

func (a *App) showBlogSearchModal(listBlogs *tview.List) {
	var input *tview.InputField
	input = tview.NewInputField().
		SetLabel("Search blog: ").
		SetFieldWidth(30).
		SetDoneFunc(func(key tcell.Key) {
			switch key {
			case tcell.KeyEnter:
				term := input.GetText()
				a.filterBlogsList(listBlogs, term)
				a.TApp.SetRoot(a.TLayout, true)
			case tcell.KeyEsc:
				a.filterBlogsList(listBlogs, "")
				a.TApp.SetRoot(a.TLayout, true)
			}
		})

	modal := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(input, 3, 1, true)

	a.TApp.SetRoot(modal, true).SetFocus(input)
}

// Filter blogs in the list by name
func (a *App) filterBlogsList(listBlogs *tview.List, term string) {
	listBlogs.Clear()
	listPosts := a.viewsList["posts"]

	found := false
	a.currentBlogs = []config.Blog{}
	for _, blog := range a.Config.Blogs {
		if term == "" || strings.Contains(strings.ToLower(blog.Name), strings.ToLower(term)) {
			found = true
			isIn := slices.Contains(a.Config.App.FollowingBlogs, blog.Name)

			r := emptyRune
			if isIn {
				r = followRune
			}

			if term != "" {
				listBlogs.AddItem("=== SEARCHING MODE === "+term, "", 0, nil)
			}

			listBlogs.AddItem(blog.Name, blog.Feed, r, func() {
				listPosts.Clear()
				listPosts.AddItem("Loading...", blog.Feed, 0, nil)
				go func() {
					feed, err := a.getFeed(blog.Feed)
					a.TApp.QueueUpdateDraw(func() {
						listPosts.Clear()
						if err != nil {
							listPosts.AddItem("Failed to load feed", err.Error(), 0, nil)
							return
						}
						for _, item := range feed.Items {
							title := item.Title
							published := ""
							if item.PublishedParsed != nil {
								published = item.PublishedParsed.Format("2006-01-02 15:04")
							}
							link := item.Link
							itemLink := link

							_, _, isIn := a.isPostSaved(item)
							r := emptyRune
							if isIn {
								r = savedRune
							}

							listPosts.AddItem(title, published, r, func() {
								err := browser.OpenURL(itemLink)
								if err != nil {
									log.Printf("failed to open URL: %v", err)
								}
							})
						}
						if len(feed.Items) == 0 {
							listPosts.AddItem("No posts found", "", 0, nil)
						}
						listPosts.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
							if event.Key() == tcell.KeyCtrlS {
								x := listPosts.GetCurrentItem()
								post := feed.Items[x]
								a.savePost(listPosts, post, true, false, blog.Name)
								return nil
							}
							return event
						})
					})
				}()
				a.TApp.SetFocus(listPosts)
			})
			a.currentBlogs = append(a.currentBlogs, blog)
		}
	}

	if !found {
		listBlogs.AddItem("No blogs found", "", 0, nil)
	}
}

func (a *App) getFeed(url string) (*gofeed.Feed, error) {
	a.Cache.RLock()
	feed, ok := a.Cache.data[url]
	a.Cache.RUnlock()
	if ok && time.Since(feed.Timestamp) < a.Cache.ttl {
		return feed.Post, nil
	}

	fp := gofeed.NewParser()
	feedRsp, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	a.Cache.Lock()
	a.Cache.data[url] = Feed{Post: feedRsp, Timestamp: time.Now()}
	a.Cache.Unlock()

	return feedRsp, nil
}

func (a *App) isPostSaved(post *gofeed.Item) (string, int, bool) {
	postHash := internal.GetHash([]string{post.Title, post.Content, post.Link})

	for i, post := range a.Config.App.SavedPosts {
		if post.Hash == postHash {
			return postHash, i, true
		}
	}

	return postHash, -1, false
}

func (a *App) savePost(listPosts *tview.List, post *gofeed.Item, refreshHome, includeBlogName bool, blogName string) {
	x := listPosts.GetCurrentItem()

	postHash, ix, isIn := a.isPostSaved(post)

	r := emptyRune
	if !isIn {
		r = savedRune
		a.Config.App.SavedPosts = append(a.Config.App.SavedPosts, config.Post{
			Title:     post.Title,
			Published: post.PublishedParsed,
			Link:      post.Link,
			Hash:      postHash,
		})
	} else {
		a.Config.App.SavedPosts = append(a.Config.App.SavedPosts[:ix], a.Config.App.SavedPosts[ix+1:]...)
	}

	err := config.SaveConfig(a.Config)
	if err != nil {
		log.Printf("failed to save config: %v", err)
		return
	}

	published := ""
	if post.PublishedParsed != nil {
		published = post.PublishedParsed.Format("2006-01-02 15:04")
	}

	updateItemList(listPosts, x, post.Title, blogName+" - "+published, r, func() {
		err := browser.OpenURL(post.Link)
		if err != nil {
			log.Printf("failed to open URL: %v", err)
		}
	})

	if refreshHome {
		a.generateHomeList(a.viewsList["home"])
	}
	a.generateSavedPosts(a.viewsList["savedPosts"])
}

func (a *App) followBlogs(listBlogs *tview.List) {
	x := listBlogs.GetCurrentItem()
	if x < 0 || x >= len(a.currentBlogs) {
		return
	}
	blog := a.currentBlogs[x]

	isIn := false
	ix := -1
	for i, id := range a.Config.App.FollowingBlogs {
		if id == blog.Name {
			isIn = true
			ix = i
			break
		}
	}

	r := emptyRune
	if !isIn {
		r = followRune
		a.Config.App.FollowingBlogs = append(a.Config.App.FollowingBlogs, blog.Name)
	} else {
		a.Config.App.FollowingBlogs = append(a.Config.App.FollowingBlogs[:ix], a.Config.App.FollowingBlogs[ix+1:]...)
	}

	err := config.SaveConfig(a.Config)
	if err != nil {
		log.Printf("failed to save config: %v", err)
	}

	updateItemList(listBlogs, x, blog.Name, blog.Feed, r, func() {})
	a.generateHomeList(a.viewsList["home"])
}

func (a *App) blogsPage(nextSlide func()) (title string, content tview.Primitive) {
	pages := tview.NewPages()
	blogPage, _, _ := a.generateBlogsList()
	pages.AddPage("blogs", blogPage, true, true)
	return "Blogs", pages
}
