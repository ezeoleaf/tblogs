package models

type Config struct {
	API APIConfig `yaml:"api"`
	APP APPConfig `yaml:"app"`
}

type APIConfig struct {
	Host string `yaml:"url"`
	Key  string `yaml:"key"`
}

type APPConfig struct {
	SavedPosts     []Post `yaml:"saved_posts"`
	FollowingBlogs []int  `yaml:"following_blogs"`
	FirstUse       bool   `yaml:"first_use"`
}
