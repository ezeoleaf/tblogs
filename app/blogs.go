package app

import (
	"github.com/ezeoleaf/tblogs/api"
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/helpers"
	"github.com/ezeoleaf/tblogs/models"
	"github.com/gdamore/tcell/v2"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

var listBlogs *tview.List
var listPosts *tview.List
var blogs models.Blogs
var blogPage *tview.Flex

func generateBlogsList() {
	blogPage = tview.NewFlex()

	blogs = api.GetBlogs()

	listBlogs.ShowSecondaryText(false)
	listBlogs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlS {
			followBlogs()
			return nil
		} else if event.Key() == tcell.KeyCtrlF {
			pages.ShowPage(blogsModalName)
			//TODO: Search blogs
		}
		return event
	})

	for _, blog := range blogs.Blogs {
		appCfg := cfg.GetAPPConfig()
		r := emptyRune
		isIn, _ := helpers.IsIn(blog.ID, appCfg.FollowingBlogs)
		if isIn {
			r = followRune
		}
		listBlogs.AddItem(blog.Name, blog.Company, r, emptyFunc)
	}

	listPosts = getList()
	listPosts.SetDoneFunc(func() {
		app.SetFocus(listBlogs)
	})
	listBlogs.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
		listPosts.Clear()
		blogID := blogs.Blogs[x].ID
		posts := api.GetPostsByBlog(blogID)
		appCfg := cfg.GetAPPConfig()
		for _, post := range posts.Posts {
			r := emptyRune
			isIn, _ := helpers.IsHash(post.Hash, appCfg.SavedPosts)
			if isIn {
				r = savedRune
			}
			listPosts.AddItem(post.Title, post.Published, r, emptyFunc)
		}

		listPosts.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
			post := posts.Posts[x]
			browser.OpenURL(post.Link)
		})

		listPosts.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlS {
				appCfg := cfg.GetAPPConfig()

				x := listPosts.GetCurrentItem()

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
				updateItemList(listPosts, x, post.Title, post.Published, r, emptyFunc)
				generateSavedPosts()
				return nil
			}
			return event
		})
		app.SetFocus(listPosts)
	})

	blogPage.AddItem(listBlogs, 0, 1, true).
		AddItem(listPosts, 0, 1, false)
}

func followBlogs() {
	appCfg := cfg.GetAPPConfig()

	x := listBlogs.GetCurrentItem()

	blog := blogs.Blogs[x]

	r := emptyRune
	isIn, ix := helpers.IsIn(blog.ID, appCfg.FollowingBlogs)
	if !isIn {
		r = followRune
		appCfg.FollowingBlogs = append(appCfg.FollowingBlogs, blog.ID)
	} else {
		appCfg.FollowingBlogs = append(appCfg.FollowingBlogs[:ix], appCfg.FollowingBlogs[ix+1:]...)
	}
	cfg.UpdateAppConfig(appCfg)

	updateItemList(listBlogs, x, blog.Name, blog.Company, r, emptyFunc)
	generateHomeList()
}

func blogsPage(nextSlide func()) (title string, content tview.Primitive) {
	pages = tview.NewPages()

	listBlogs = getList()

	generateBlogsList()

	pages.AddPage("blogs", blogPage, true, true)

	return blogsSection, pages
}

func searchBlogs() {
	//TODO: Add ability to search
}

func cancelSearchBlogs() {
	//TODO: Add ability to search
}

// Backup version
// func openBrowser(url string) {
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		err = exec.Command("xdg-open", url).Start()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	default:
// 		err = fmt.Errorf("unsupported platform")
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
