package app

import (
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

// var listBlogs *tview.List
// var listPosts *tview.List
// var blogs models.Blogs
// var blogPage *tview.Flex

func (a *Tblogs) getBlogs() *tview.Flex {
	blogPage := tview.NewFlex()
	listBlogs := tview.NewList()

	blogs := a.dataService.GetBlogs()

	listBlogs.ShowSecondaryText(false)
	// listBlogs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	if event.Key() == tcell.KeyCtrlS {
	// 		followBlogs()
	// 		return nil
	// 	}
	// 	return event
	// })

	for _, blog := range blogs.Blogs {
		// appCfg := cfg.GetAPPConfig()
		r := emptyRune
		// isIn, _ := helpers.IsIn(blog.ID, appCfg.FollowingBlogs)
		// if isIn {
		// 	r = followRune
		// }
		listBlogs.AddItem(blog.Title, blog.Company, r, emptyFunc)
	}

	listPosts := getList()

	listPosts.SetDoneFunc(func() {
		a.app.SetFocus(listBlogs)
	})

	listBlogs.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
		listPosts.Clear()
		blog := blogs.Blogs[x]
		posts := a.dataService.GetPosts(blog)
		// appCfg := cfg.GetConfig()
		for _, post := range posts {
			r := emptyRune
			// isIn, _ := helpers.IsHash(post.Hash, appCfg.SavedPosts)
			// if isIn {
			// 	r = savedRune
			// }
			listPosts.AddItem(post.Title, post.Published, r, emptyFunc)
		}

		listPosts.SetSelectedFunc(func(x int, s string, s1 string, r rune) {
			post := posts[x]
			browser.OpenURL(post.Link)
		})

		// listPosts.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// 	if event.Key() == tcell.KeyCtrlS {
		// 		appCfg := cfg.GetConfig()

		// 		x := listPosts.GetCurrentItem()

		// 		post := posts[x]

		// 		r := emptyRune
		// 		isIn, ix := helpers.IsHash(post.Hash, appCfg.SavedPosts)
		// 		if !isIn {
		// 			r = savedRune
		// 			appCfg.SavedPosts = append(appCfg.SavedPosts, post)
		// 		} else {
		// 			appCfg.SavedPosts = append(appCfg.SavedPosts[:ix], appCfg.SavedPosts[ix+1:]...)
		// 		}
		// 		cfg.UpdateAppConfig(appCfg)
		// 		updateItemList(listPosts, x, post.Title, post.Published, r, emptyFunc)
		// 		// generateSavedPosts()
		// 		return nil
		// 	}
		// 	return event
		// })
		a.app.SetFocus(listPosts)
	})

	blogPage.AddItem(listBlogs, 0, 1, true).
		AddItem(listPosts, 0, 1, false)

	return blogPage
}

// func followBlogs() {
// 	appCfg := cfg.GetAPPConfig()

// 	x := listBlogs.GetCurrentItem()

// 	blog := blogs.Blogs[x]

// 	r := emptyRune
// 	isIn, ix := helpers.IsIn(blog.ID, appCfg.FollowingBlogs)
// 	if !isIn {
// 		r = followRune
// 		appCfg.FollowingBlogs = append(appCfg.FollowingBlogs, blog.ID)
// 	} else {
// 		appCfg.FollowingBlogs = append(appCfg.FollowingBlogs[:ix], appCfg.FollowingBlogs[ix+1:]...)
// 	}
// 	cfg.UpdateAppConfig(appCfg)

// 	updateItemList(listBlogs, x, blog.Name, blog.Company, r, emptyFunc)
// 	generateHomeList()
// }

func (a Tblogs) setBlogsPage(nextSlide func()) (title string, content tview.Primitive) {
	pages := tview.NewPages()

	// blogPage = getList()

	blogPage := a.getBlogs()

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
