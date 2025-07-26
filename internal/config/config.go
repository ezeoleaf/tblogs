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
	LastLoginMode  bool      `yaml:"last_login_mode"`
	LastLogin      time.Time `yaml:"last_login"`
	CurrentLogin   time.Time `yaml:"current_login"`
	FilteredWords  []string  `yaml:"filtered_words"`
	XCred          XCred     `yaml:"x_cred"`
}

type XCred struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	AccessToken  string `yaml:"access_token"`
	RefreshToken string `yaml:"refresh_token"`
	Username     string `yaml:"username"`
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

func getConfigPath() (string, error) {
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

// LoadConfig loads the YAML config from the given path (or default path if empty) and returns a Config struct
func LoadConfig() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			blogs, err := loadDefaultBlogs()
			if err != nil {
				return nil, err
			}

			// File does not exist, create default config
			cfg := &Config{
				App: AppConfig{
					SavedPosts:     []Post{},
					FollowingBlogs: []string{},
					FilteredWords:  []string{},
					XCred:          XCred{},
				},
				Blogs: blogs,
			}

			// Save the default config
			if err := SaveConfig(cfg); err != nil {
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

	// Sync blogs with defaults
	if err := cfg.syncBlogs(); err != nil {
		log.Printf("warning: failed to sync blogs: %v", err)
		// Don't fail the load, just log the warning
	}

	cfg.App.CurrentLogin = time.Now()

	return &cfg, nil
}

func loadDefaultBlogs() ([]Blog, error) {
	f, err := os.Open("internal/config/default-blogs.yml")
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

func (cfg *Config) syncBlogs() error {
	defaultBlogs, err := loadDefaultBlogs()
	if err != nil {
		return err
	}

	// Create a map of existing blogs by name for quick lookup
	existingBlogs := make(map[string]Blog)
	for _, blog := range cfg.Blogs {
		existingBlogs[blog.Name] = blog
	}

	// Add missing blogs from default
	for _, defaultBlog := range defaultBlogs {
		if _, exists := existingBlogs[defaultBlog.Name]; !exists {
			cfg.Blogs = append(cfg.Blogs, defaultBlog)
		}
	}

	return nil
}

// SaveConfig writes the config struct to the given YAML file path (or default path if empty)
func SaveConfig(cfg *Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
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
