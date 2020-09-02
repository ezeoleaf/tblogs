package models

type Blog struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Feed    string `json:"feed"`
}

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}
