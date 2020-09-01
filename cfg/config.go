package cfg

import (
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

const configPath = "./cfg/config.yml"

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
	FirstUse       bool  `yaml:"first_use"`
}

func parseFlags() (string, error) {
	// String that contains the configured configuration path
	// var configPath string

	// // Set up a CLI flag called "-config" to allow users
	// // to supply the configuration file
	// flag.StringVar(&configPath, "config", "./cfg/config.yml", "./cfg/config.yml")

	// // Actually parse the flags
	// flag.Parse()

	configPath := configPath
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

func updateConfig() {
	d, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(cfgPath, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateAppConfig(a APPConfig) {
	config.APP = a

	updateConfig()
}

func Setup() {
	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

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
