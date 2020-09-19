package models

import (
	"time"
)

// Config contains both API and APP configuration from config file
type Config struct {
	API APIConfig `yaml:"api"`
	APP APPConfig `yaml:"app"`
}

// APIConfig contains only the API configuration from config file
type APIConfig struct {
	Host string `yaml:"url"`
	Key  string `yaml:"key"`
}

// APPConfig contains only the APP configuration from config file
type APPConfig struct {
	SavedPosts     []Post    `yaml:"saved_posts"`
	FollowingBlogs []int     `yaml:"following_blogs"`
	FirstUse       bool      `yaml:"first_use"`
	LastLogin      time.Time `yaml:"last_login"`
	CurrentLogin   time.Time `yaml:"current_login"`
}
