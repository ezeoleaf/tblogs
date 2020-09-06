package models

// Blog represents a blog from Blogio API
type Blog struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Feed    string `json:"feed"`
}

// Blogs is a list of Blog
type Blogs struct {
	Blogs []Blog `json:"blogs"`
}
