package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	WorldDataDirectory string `yaml:"data_dir"`
	WorldSaveDirectory string `yaml:"save_dir"`
	WorldSaveTarget    string `yaml:"save_target"`
	WorldWebDomain     string `yaml:"web_domain"`
}

var Cfg *Config

// LoadConfig - load config from file if configFile != ""
func LoadConfig(configFile string) {
	var err error
	Cfg, err = getConfigFromFile(configFile)
	if err != nil {
		Cfg = getConfigFromEnv()
	}
}

// getConfigFromFile - configFile example: "config.yaml"
func getConfigFromFile(configFile string) (config *Config, err error) {
	var file *os.File
	file, err = os.Open(configFile)
	if err != nil {
		return &Config{}, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return &Config{}, err
	}
	return config, nil
}

func getConfigFromEnv() *Config {
	return &Config{
		WorldDataDirectory: os.Getenv("WORLD_DATA_DIRECTORY"),
		WorldSaveDirectory: os.Getenv("WORLD_SAVE_DIRECTORY"),
		WorldSaveTarget:    os.Getenv("WORLD_SAVE_TARGET"),
		WorldWebDomain:     os.Getenv("WORLD_WEB_DOMAIN"),
	}
}
