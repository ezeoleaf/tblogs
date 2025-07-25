package app

import "github.com/mmcdole/gofeed"

type PostWithMeta struct {
	Item     *gofeed.Item
	BlogName string
}
