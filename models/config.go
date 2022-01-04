package models

import (
	"time"

	"github.com/ezeoleaf/tblogs/data"
)

// Config contains the APP configuration from config file
type Config struct {
	SavedPosts     []data.Post `yaml:"saved_posts"`
	FollowingBlogs []int       `yaml:"following_blogs"`
	FirstUse       bool        `yaml:"first_use"`
	LastLogin      time.Time   `yaml:"last_login"`
	CurrentLogin   time.Time   `yaml:"current_login"`
	FilteredWords  []string    `yaml:"filtered_words"`
}
