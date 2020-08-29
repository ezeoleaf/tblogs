package main

import (
	"fmt"

	"github.com/ezeoleaf/tblogs/api"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func Grid(nextSlide func()) (title string, content tview.Primitive) {
	pages := tview.NewPages()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewFrame(nil).
			SetBorders(0, 0, 0, 0, 0, 0).
			AddText(text, true, tview.AlignCenter, tcell.ColorWhite)
	}

	blogs := newPrimitive("Following Blogs")
	posts := newPrimitive("Posts")

	grid := tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 0, 0).
		SetBorders(true)
	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyDown {
			fmt.Println("Down key")
			return nil
		}
		// if modalShown {
		// 	nextSlide()
		// 	modalShown = false
		// } else {
		// 	pages.ShowPage("modal")
		// 	modalShown = true
		// }
		return nil
	})

	// Layout for screens narrower than 100 cells (blogs and side bar are hidden).
	grid.AddItem(blogs, 0, 0, 0, 0, 0, 0, true).
		AddItem(posts, 1, 0, 1, 3, 0, 0, false)

	b := api.GetBlogs()
	list := tview.NewList()
	list.AddItem("Following blogs", "", ' ', nextSlide)
	for _, blog := range b.Blogs {
		list.AddItem(blog.Name, blog.Company, 'x', nextSlide)
	}

	// Layout for screens wider than 100 cells.
	grid.AddItem(list, 0, 0, 10, 1, 0, 100, true).
		AddItem(posts, 0, 1, 10, 4, 0, 100, false)
	// AddItem(list, 0, 1, 10, 4, 0, 100, false)

	pages.AddPage("grid", grid, true, true)

	return "Grid", pages
}
