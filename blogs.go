package main

import (
	"github.com/ezeoleaf/tblogs/api"

	"github.com/rivo/tview"
)

func BlogPage(nextSlide func()) (title string, content tview.Primitive) {
	b := api.GetBlogs()
	list := tview.NewList()
	for _, blog := range b.Blogs {
		list.AddItem(blog.Name, blog.Company, 'x', nextSlide)
	}

	return "Blogs", Center(80, 10, list)
}
