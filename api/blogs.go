package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ezeoleaf/tblogs/cfg"
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

	cfgAPI := cfg.GetConfig().API

	request, err := http.NewRequest("GET", cfgAPI.Host, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Add("BLOGIO-KEY", cfgAPI.Key)

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
