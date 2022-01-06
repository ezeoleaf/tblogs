package app

import (
	"github.com/rivo/tview"
)

func (a *Tblogs) initTwitter() *tview.List {

	twitterList := getList()

	if a.twitterClient == nil {
		twitterList.AddItem("No Twitter Client", "Try go to settings and add your Twitter keys", ' ', nil)
	} else {
		tweets, _ := a.twitterClient.GetTimeline()

		for _, t := range tweets {
			twitterList.AddItem(t.Text, "", ' ', nil)
		}
	}

	return twitterList
}

func (a *Tblogs) initTwitterPage(nextSlide func()) (title string, content tview.Primitive) {
	twitterList := a.initTwitter()

	return TwitterSection, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(twitterList, 0, 1, true), 0, 1, true)
}
