package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ezeoleaf/tblogs/models"

	"github.com/ezeoleaf/tblogs/cfg"
)

var blogs models.Blogs

// GetBlogs returns a list of Blogs from the Blogio API
func GetBlogs() models.Blogs {

	if len(blogs.Blogs) > 0 {
		return blogs
	}

	client := &http.Client{}

	cfgAPI := cfg.GetAPIConfig()

	request, err := http.NewRequest("GET", cfgAPI.Host+"/blogs", nil)

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

	blogs = models.Blogs{}
	err = json.Unmarshal(body, &blogs)
	if err != nil {
		panic(err)
	}

	return blogs
}
