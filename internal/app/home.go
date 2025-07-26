package app

import (
	"log"
	"sort"
	"sync"

	"github.com/ezeoleaf/tblogs/internal/config"
	"github.com/gdamore/tcell/v2"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func (a *App) findBlogByName(name string) (config.Blog, bool) {
	for _, blog := range a.Config.Blogs {
		if blog.Name == name {
			return blog, true
		}
	}
	return config.Blog{}, false
}

func (a *App) generateHomeList(listHome *tview.List) {
	listHome.Clear()
	listHome.AddItem("Loading posts...", "", 0, nil)

	postsCh := make(chan PostWithMeta)
	doneCh := make(chan struct{})
	var wg sync.WaitGroup

	for _, blogName := range a.Config.App.FollowingBlogs {
		blog, ok := a.findBlogByName(blogName)
		if !ok {
			continue
		}
		wg.Add(1)
		go func(blog config.Blog) {
			defer wg.Done()
			feed, err := a.getFeed(blog.Feed)
			if err != nil {
				return
			}
			for _, item := range feed.Items {
				if a.Config.App.LastLoginMode {
					if item.PublishedParsed.Before(a.Config.App.LastLogin) {
						continue
					}
				}
				postsCh <- PostWithMeta{Item: item, BlogName: blog.Name}
			}
		}(blog)
	}

	// Close postsCh when all goroutines are done
	go func() {
		wg.Wait()
		close(postsCh)
		close(doneCh)
	}()

	// Collect posts
	var allPosts []PostWithMeta
	go func() {
		for post := range postsCh {
			allPosts = append(allPosts, post)
		}
	}()

	// When done, sort and display
	go func() {
		<-doneCh
		sort.Slice(allPosts, func(i, j int) bool {
			t1, t2 := allPosts[i].Item.PublishedParsed, allPosts[j].Item.PublishedParsed
			if t1 == nil && t2 == nil {
				return allPosts[i].Item.Title > allPosts[j].Item.Title
			}
			if t1 == nil {
				return false
			}
			if t2 == nil {
				return true
			}
			return t1.After(*t2)
		})
		a.TApp.QueueUpdateDraw(func() {
			listHome.Clear()

			if a.Config.App.LastLoginMode {
				listHome.AddItem("=== LAST LOGIN MODE ===", "Only showing posts since last login", 0, nil)
			} else {
				listHome.AddItem("=== ALL POSTS MODE ===", "Showing all posts from followed blogs", 0, nil)
			}

			if len(allPosts) == 0 {
				listHome.AddItem("No posts found", "", 0, nil)
				return
			}
			a.currentHomePosts = make([]PostWithMeta, len(allPosts))
			for i, post := range allPosts {
				_, _, isIn := a.isPostSaved(post.Item)
				r := emptyRune
				if isIn {
					r = savedRune
				}
				link := post.Item.Link
				itemLink := link
				listHome.AddItem(post.Item.Title, post.BlogName+" - "+post.Item.PublishedParsed.Format("2006-01-02 15:04"), r, func() {
					err := browser.OpenURL(itemLink)
					if err != nil {
						log.Printf("failed to open URL: %v", err)
					}
				})
				a.currentHomePosts[i] = post
			}
		})
	}()

	// Set the input handler ONCE, outside the goroutine
	listHome.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlS {
			x := listHome.GetCurrentItem()
			if x >= 0 && x < len(a.currentHomePosts) {
				a.savePost(listHome, a.currentHomePosts[x].Item, false, true, a.currentHomePosts[x].BlogName)
			}
			return nil
		}

		if event.Key() == tcell.KeyCtrlL {
			a.Config.App.LastLoginMode = !a.Config.App.LastLoginMode
			err := config.SaveConfig(a.Config)
			if err != nil {
				log.Printf("failed to save config: %v", err)
			}
			a.generateHomeList(listHome)
		}

		return event
	})
}

func (a *App) homePage(nextSlide func()) (title string, content tview.Primitive) {
	listHome := a.viewsList["home"]
	if listHome == nil {
		listHome = getList()
		a.viewsList["home"] = listHome
		a.generateHomeList(listHome)
	}

	title = homeSection
	if a.Config.App.LastLoginMode {
		title = "Home -- Last Login Mode"
	}

	return title, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listHome, 0, 1, true), 0, 1, true)
}
