package config

import (
	"gopkg.in/yaml.v2"

	"fmt"
	"os"
	"time"
)

type Config struct {
	Env          string       `yaml:"env"`
	DBType       string       `yaml:"db"`
	MySQLStorage MySQLStorage `yaml:"mysql`
	Redis        Redis        `yaml:"redis`
	HttpServer   HttpServer   `yaml:"httpServer"`
}

type MySQLStorage struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	DBName   string `yaml:"DBName"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"password"`
}

type Redis struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type HttpServer struct {
	Port         string        `yaml:"port"`
	Host         string        `yaml:"host"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, fmt.Errorf("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)

	}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
