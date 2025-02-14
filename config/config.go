package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config structure to hold all configurations
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

var AppConfig Config

// LoadConfig loads configuration from config.yml
func LoadConfig() {
	
	// Allow setting config file path via env variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yml" // Default path
	}
	
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML file: %s", err)
	}

	// 🔥 Always use OS-defined PORT if available
	if envPort := os.Getenv("PORT"); envPort != "" {
		AppConfig.Server.Port = envPort
		log.Printf("Using PORT from environment: %s", envPort)
	} else {
		log.Printf("Using PORT from config.yml: %s", AppConfig.Server.Port)
	}

	log.Println("Configuration loaded successfully")
}
