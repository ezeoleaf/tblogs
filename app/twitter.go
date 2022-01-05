package app

import (
	"github.com/rivo/tview"
)

func (a *Tblogs) initTwitter() *tview.List {

	twitterList := getList()

	return twitterList
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

func (a *Tblogs) initTwitterPage(nextSlide func()) (title string, content tview.Primitive) {
	twitterList := a.initTwitter()
	// generateHomeList()

	return TwitterSection, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(twitterList, 0, 1, true), 0, 1, true)
}
