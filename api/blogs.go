package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Blog struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Feed    string `json:"feed"`
}
type Blogs struct {
	Blogs []Blog `json:"blogs"`
}

var blogs Blogs

func GetBlogs() Blogs {

	if len(blogs.Blogs) > 0 {
		return blogs
	}

	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://api.dev-blogs.tech/api/blogs", nil)

	if err != nil {
		panic(err)
	}

	request.Header.Add("BLOGIO-KEY", "LALA")

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	blogs = Blogs{}
	err = json.Unmarshal(body, &blogs)
	if err != nil {
		panic(err)
	}

	return blogs
}
