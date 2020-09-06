package models

import "time"

// Post represents a post of a blogs from Blogio API
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

// PostCache contains a list of fetched Posts saved in memory for a certain amount of time
type PostCache struct {
	Posts       Posts
	DateUpdated time.Time
}

// Posts is a list of Post
type Posts struct {
	Posts []Post `json:"posts"`
}

// PostRequest represents the request sent to Blogio API
type PostRequest struct {
	Blogs []int `json:"blogs"`
}
