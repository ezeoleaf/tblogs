package main

import (
	"github.com/ezeoleaf/tblogs/api"
	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/helpers"
	"github.com/gdamore/tcell"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

func Blogs(nextSlide func()) (title string, content tview.Primitive) {

	code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

	listBlogs := tview.NewList()

	b := api.GetBlogs()

	listBlogs.SetBorderPadding(1, 1, 2, 2)
	listBlogs.ShowSecondaryText(false)
	listBlogs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlF {
			appCfg := cfg.GetAPPConfig()

			x := listBlogs.GetCurrentItem()

			blog := b.Blogs[x]

			r := ' '
			isIn, ix := helpers.IsIn(blog.ID, appCfg.FollowingBlogs)
			if !isIn {
				r = 'f'
				appCfg.FollowingBlogs = append(appCfg.FollowingBlogs, blog.ID)
				cfg.UpdateAppConfig(appCfg)
			} else {
				appCfg.FollowingBlogs = append(appCfg.FollowingBlogs[:ix], appCfg.FollowingBlogs[ix+1:]...)
				cfg.UpdateAppConfig(appCfg)
			}

			listBlogs.RemoveItem(x)
			listBlogs.InsertItem(x, blog.Name, blog.Company, r, func() {
				return
			})
			return nil
		}
		return event
	})
	for _, blog := range b.Blogs {
		appCfg := cfg.GetAPPConfig()
		r := ' '
		isIn, _ := helpers.IsIn(blog.ID, appCfg.FollowingBlogs)
		if isIn {
			r = 'f'
		}
		listBlogs.AddItem(blog.Name, blog.Company, r, func() {
			return
		})
	}

	listPosts := tview.NewList()
	listPosts.SetBorderPadding(1, 1, 2, 2)
	listPosts.SetDoneFunc(func() {
		app.SetFocus(listBlogs)
	})

	listBlogs.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
		listPosts.Clear()
		blogID := b.Blogs[x].ID
		posts := api.GetPostsByBlog(blogID)
		for _, post := range posts.Posts {
			listPosts.AddItem(post.Title, post.Published, '-', func() {
				return
			})
		}

		listPosts.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
			post := posts.Posts[x]
			browser.OpenURL(post.Link)
		})
		app.SetFocus(listPosts)
	})

	return "Blogs", tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listBlogs, 0, 1, true), 0, 1, true).
		AddItem(listPosts, 100, 1, false)
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
