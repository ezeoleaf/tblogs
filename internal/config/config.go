package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	SavedPosts     []Post    `yaml:"saved_posts"`
	FollowingBlogs []string  `yaml:"following_blogs"`
	FirstUse       bool      `yaml:"first_use"`
	LastLogin      time.Time `yaml:"last_login"`
	CurrentLogin   time.Time `yaml:"current_login"`
	FilteredWords  []string  `yaml:"filtered_words"`
}

type Post struct {
	Title     string     `yaml:"title"`
	Published *time.Time `yaml:"published"`
	Link      string     `yaml:"link"`
	Hash      string     `yaml:"hash"`
}

type Blog struct {
	Name string `yaml:"name"`
	Feed string `yaml:"feed"`
}

type Config struct {
	App   AppConfig `yaml:"app"`
	Blogs []Blog    `yaml:"blogs"`
}

// GetConfigPath returns the OS-appropriate config file path for tblogs
const defaultConfigFile = "data.yml"

func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	tblogsDir := filepath.Join(configDir, "tblogs")
	if err := os.MkdirAll(tblogsDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(tblogsDir, defaultConfigFile), nil
}

func loadDefaultBlogsYAML(path string) ([]Blog, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	var blogs []Blog

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

// LoadConfig loads the YAML config from the given path (or default path if empty) and returns a Config struct
func LoadConfig(path string) (*Config, error) {
	if path == "" {
		var err error
		path, err = GetConfigPath()
		if err != nil {
			return nil, err
		}
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			blogs, err := loadDefaultBlogsYAML("internal/config/default-blogs.yml")
			if err != nil {
				return nil, err
			}

			// File does not exist, create default config
			cfg := &Config{
				App: AppConfig{
					SavedPosts:     []Post{},
					FollowingBlogs: []string{},
					FirstUse:       true,
					FilteredWords:  []string{},
				},
				Blogs: blogs,
			}

			// Save the default config
			if err := SaveConfig(cfg, path); err != nil {
				return nil, err
			}

			return cfg, nil
		}

		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	var cfg Config

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// SaveConfig writes the config struct to the given YAML file path (or default path if empty)
func SaveConfig(cfg *Config, path string) error {
	if path == "" {
		var err error

		path, err = GetConfigPath()
		if err != nil {
			return err
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	encoder := yaml.NewEncoder(f)
	defer func() {
		if err := encoder.Close(); err != nil {
			log.Printf("failed to close encoder: %v", err)
		}
	}()

	return encoder.Encode(cfg)
}
