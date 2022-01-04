package data

import (
	"encoding/json"
	"io/ioutil"
)

type Blog struct {
	Title   string `json:"title"`
	Company string `json:"company"`
	URL     string `json:"url"`
}

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}

func (s *Service) populateBlogs() {
	file, _ := ioutil.ReadFile(dataFileName)

	_ = json.Unmarshal([]byte(file), &s.blogs)

	// TODO: Add validation for when file could not be open or unmarshall
}

func (s *Service) GetBlogs() Blogs {
	return s.blogs
}
