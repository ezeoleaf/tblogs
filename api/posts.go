package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ezeoleaf/tblogs/cfg"
	"github.com/ezeoleaf/tblogs/models"
)

var posts map[int]models.PostCache

const defaultTime = 4

// GetPostsByBlog returns a list of Posts for a single blog
func GetPostsByBlog(blogID int) models.Posts {

	if len(posts[blogID].Posts.Posts) > 0 {
		d := time.Now()
		diff := d.Sub(posts[blogID].DateUpdated).Hours()

		if diff < defaultTime {
			return posts[blogID].Posts
		}
	}

	if len(posts) == 0 {
		posts = make(map[int]models.PostCache)
	}

	pr := models.PostRequest{Blogs: []int{blogID}}

	postsResp := fetchPosts(pr)

	pc := models.PostCache{Posts: postsResp, DateUpdated: time.Now()}

	posts[blogID] = pc

	return postsResp
}

// GetPosts returns a list of Posts for a list of Blogs using the Blogs ids
func GetPosts(blogs []int) models.Posts {
	pr := models.PostRequest{Blogs: blogs}

	postsResp := fetchPosts(pr)

	return postsResp
}

func fetchPosts(reqPost models.PostRequest) models.Posts {
	rJSON, err := json.Marshal(reqPost)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	cfgAPI := cfg.GetAPIConfig()

	payload := strings.NewReader(string(rJSON))

	req, err := http.NewRequest("GET", cfgAPI.Host+"/posts", payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("BLOGIO-KEY", cfgAPI.Key)
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

	posts := models.Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		panic(err)
	}
	return posts
}
