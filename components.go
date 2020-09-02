package main

import "github.com/rivo/tview"

const (
	emptyRune  = ' '
	savedRune  = 's'
	followRune = 'f'
)

var emptyFunc = func() {
	return
}

func getList() *tview.List {
	gList := tview.NewList()
	gList.SetBorderPadding(1, 1, 2, 2)

	return gList
}

func addItemToList(l *tview.List, title string, subtitle string, r rune, f func()) {

}

func updateItemList(l *tview.List, x int, title string, subtitle string, r rune, f func()) {
	l.RemoveItem(x)
	l.InsertItem(x, title, subtitle, r, f)
	l.SetCurrentItem(x)
}
