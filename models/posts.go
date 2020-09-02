package models

import "time"

type Post struct {
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Blog        string     `json:"blog"`
	BlogID      int64      `json:"blog_id"`
	Published   string     `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
	Link        string     `json:"link"`
	Hash        string     `json:"hash"`
}
type PostCache struct {
	Posts       Posts
	DateUpdated time.Time
}

type Posts struct {
	Posts []Post `json:"posts"`
}

type PostRequest struct {
	Blogs []int `json:"blogs"`
}
