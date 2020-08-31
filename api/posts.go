package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Post struct {
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Blog        string     `json:"blog"`
	BlogID      int64      `json:"blog_id"`
	Published   string     `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
}

// Link        string     `json:"link"`
// Hash string `json:"hash"`
// Saved       bool       `json:"saved"`

var posts map[int]PostCache

const defaultTime = 4

// var PostsCache []PostsList

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

func GetPostsByBlog(blogID int) Posts {

	if len(posts[blogID].Posts.Posts) > 0 {
		d := time.Now()
		diff := d.Sub(posts[blogID].DateUpdated).Hours()

		if diff < defaultTime {
			return posts[blogID].Posts
		}
	}

	pr := PostRequest{Blogs: []int{blogID}}

	if len(posts) == 0 {
		posts = make(map[int]PostCache)
	}

	postsResp := GetPosts(pr)

	pc := PostCache{Posts: postsResp, DateUpdated: time.Now()}

	posts[blogID] = pc

	return postsResp
}

func GetPosts(reqPost PostRequest) Posts {
	rJSON, err := json.Marshal(reqPost)
	if err != nil {
		panic(err)
	}
	// fmt.Println(rJSON)
	client := &http.Client{}

	payload := strings.NewReader(string(rJSON))

	req, err := http.NewRequest("GET", "https://api.dev-blogs.tech/api/posts", payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("blogio-key", "LALA")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		panic(err)
	}
	return posts
}
