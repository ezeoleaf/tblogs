package main

import (
	"github.com/ezeoleaf/tblogs/api"
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

	// basic := func() {
	// 	table.SetBorders(false).
	// 		SetSelectable(false, false).
	// 		SetSeparator(' ')
	// 	code.Clear()
	// 	fmt.Fprint(code, tableBasic)
	// }

	// separator := func() {
	// 	table.SetBorders(false).
	// 		SetSelectable(false, false).
	// 		SetSeparator(tview.Borders.Vertical)
	// 	code.Clear()
	// 	fmt.Fprint(code, tableSeparator)
	// }

	// borders := func() {
	// 	table.SetBorders(true).
	// 		SetSelectable(false, false)
	// 	code.Clear()
	// 	fmt.Fprint(code, tableBorders)
	// }

	// selectRow := func() {
	// 	table.SetBorders(false).
	// 		SetSelectable(true, false).
	// 		SetSeparator(' ')
	// 	code.Clear()
	// 	fmt.Fprint(code, tableSelectRow)
	// }

	// selectColumn := func() {
	// 	table.SetBorders(false).
	// 		SetSelectable(false, true).
	// 		SetSeparator(' ')
	// 	code.Clear()
	// 	fmt.Fprint(code, tableSelectColumn)
	// }

	// selectCell := func() {
	// 	table.SetBorders(false).
	// 		SetSelectable(true, true).
	// 		SetSeparator(' ')
	// 	code.Clear()
	// 	fmt.Fprint(code, tableSelectCell)
	// }

	// navigate := func() {
	// 	app.SetFocus(table)
	// 	table.SetDoneFunc(func(key tcell.Key) {
	// 		app.SetFocus(list)
	// 	}).SetSelectedFunc(func(row int, column int) {
	// 		app.SetFocus(list)
	// 	})
	// }

	b := api.GetBlogs()
	listBlogs.SetBorderPadding(1, 1, 2, 2)
	listBlogs.ShowSecondaryText(false)
	listBlogs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlF {
			x := listBlogs.GetCurrentItem()
			blog := b.Blogs[x]
			listBlogs.RemoveItem(x)
			listBlogs.InsertItem(x, blog.Name, blog.Company, 'f', func() {
				return
			})
			return nil
		}
		return event
	})
	for _, blog := range b.Blogs {
		listBlogs.AddItem(blog.Name, blog.Company, ' ', func() {
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
