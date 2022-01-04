package data

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/mmcdole/gofeed"
)

const (
	separator = `,`
	rss       = "rss"
	atom      = "atom"
)

type Post struct {
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Link        string     `json:"link"`
	Hash        string     `json:"hash"`
	Saved       bool       `json:"saved"`
	Blog        string     `json:"blog"`
	BlogID      int64      `json:"blog_id"`
	Published   string     `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

func (s Service) GetPosts(blog Blog) []Post {

	feeds, err := s.retrieveFeed(blog.URL)
	if err != nil {
		return nil
	}

	posts := []Post{}

	for _, item := range feeds.Items {
		if feeds.FeedType == rss {
			item.Content = item.Description
		}
		timeParsed := ""
		if item.PublishedParsed != nil {
			timeParsed = item.PublishedParsed.Format("January 2, 2006")
		}
		hash := s.hash([]string{item.Title, item.Content})

		p := Post{
			Title:       item.Title,
			Content:     item.Content,
			Link:        item.Link,
			Hash:        hash,
			Blog:        blog.Title,
			Published:   timeParsed,
			PublishedAt: item.PublishedParsed,
		}

		posts = append(posts, p)
	}

	return posts
}

func (s Service) hash(textToHash []string) string {

	toHash := ""

	for _, text := range textToHash {
		toHash += text
	}

	hashed := md5.Sum([]byte(toHash))
	hash := hex.EncodeToString(hashed[:])

	return hash
}

func (s Service) retrieveFeed(blogURI string) (*gofeed.Feed, error) {
	if blogURI == "" {
		return nil, errors.New("no feed URI provided")
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(blogURI)

	if err != nil {
		return nil, err
	}

	return feed, err
}
