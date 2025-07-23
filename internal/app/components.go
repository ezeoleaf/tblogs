package app

import (
	"github.com/rivo/tview"
)

const (
	emptyRune  = ' '
	savedRune  = 's'
	followRune = 'f'
	noOpenRune = 'o'
)

func getList() *tview.List {
	gList := tview.NewList()
	gList.SetBorderPadding(1, 1, 2, 2)

	return gList
}

func updateItemList(l *tview.List, x int, title string, subtitle string, r rune, f func()) {
	l.RemoveItem(x)
	l.InsertItem(x, title, subtitle, r, f)
	l.SetCurrentItem(x)
}
