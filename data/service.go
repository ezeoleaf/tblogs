package data

type Service struct {
	blogs Blogs
}

const dataFileName = "../../data/data.json"

func NewService() Service {
	ds := Service{}
	ds.populateBlogs()
	return ds
}
