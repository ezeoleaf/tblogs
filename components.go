package main

import "github.com/rivo/tview"

func getList() *tview.List {
	gList := tview.NewList()
	gList.SetBorderPadding(1, 1, 2, 2)

	return gList
}
