package cfg

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	api = "API"
	app = "APP"
)

var config Config

type Config struct {
	API APIConfig `yaml:"api"`
	APP APPConfig `yaml:"app"`
}

type APIConfig struct {
	Host string `yaml:"url"`
	Key  string `yaml:"key"`
}

type APPConfig struct {
	SavedBlogs     []int `yaml:"saved_blogs"`
	SavedPosts     []int `yaml:"saved_posts"`
	FollowingBlogs []int `yaml:"following_blogs"`
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
	return config.APP
}

func GetAPIConfig() APIConfig {
	return config.API
}

func GetConfig() Config {
	return config
}

func UpdateConfig(c Config) {
	d, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("changed.yaml", d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Setup() {
	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	config = Config{}

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
}
