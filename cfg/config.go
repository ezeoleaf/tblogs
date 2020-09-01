package cfg

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	api = "API"
	app = "APP"
)

type APIConfig struct {
	Config struct {
		Host string `yaml:"url"`
		Key  string `yaml:"key"`
	} `yaml:"api"`
}

type APPConfig struct {
	Config struct {
		SavedBlogs     []int `yaml:"saved_blogs"`
		SavedPosts     []int `yaml:"saved_posts"`
		FollowingBlogs []int `yaml:"following_blogs"`
	} `yaml:"app"`
}

func parseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./cfg/config.yml", "./cfg/config.yml")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func GetAPPConfig() APPConfig {
	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	config := APPConfig{}

	file, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
		return APPConfig{}
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		log.Fatal(err)
		return APPConfig{}
	}
	return config
}

func GetAPIConfig() APIConfig {
	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	config := APIConfig{}

	file, err := os.Open(cfgPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		panic(err)
	}
	return config
}
