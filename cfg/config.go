package cfg

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ezeoleaf/tblogs/models"
	"gopkg.in/yaml.v2"
)

var config models.Config

const configPath = "./cfg/config.yml"

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

func GetAPPConfig() models.APPConfig {
	return config.APP
}

func GetAPIConfig() models.APIConfig {
	return config.API
}

func GetConfig() models.Config {
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

func UpdateAppConfig(a models.APPConfig) {
	config.APP = a

	updateConfig()
}

func setNewFile() (string, error) {
	from, err := os.Open("./cfg/config.example.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

	return parseFlags()
}

func Setup() {
	cfgPath, err := parseFlags()
	if err != nil {
		log.Println(err)
	}

	if cfgPath == "" {
		cfgPath, err = setNewFile()
		if err != nil {
			log.Fatal(err)
		}
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
