package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppType int

const (
	Pub AppType = iota
	Sub
	Other
)

// Конфигурация сервиса
type Config struct {
	DB   PostgresConfig      `yaml:"database"`
	NS   NatsStreamingConfig `yaml:"nats_streaming"`
	Http HttpConfig          `yaml:"http"`
}

// Конфигурация базы данных
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// Конфигурация http приложения
type HttpConfig struct {
	Port int `yaml:"port"`
}

// Конфигурация брокера сообщений
type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	ChanName  string `yaml:"channel_name"`
}

func NewConfig(appType string) (*Config, error) {
	switch getAppType(appType) {

	case Pub:
		configYAML, err := os.ReadFile("pub_config.yaml")
		if err != nil {
			return nil, fmt.Errorf("error while reading config file, %w", err)
		}
		config := &Config{}

		err = yaml.Unmarshal(configYAML, &config)
		if err != nil {
			return nil, fmt.Errorf("error while reading config file, %w", err)
		}
		return config, nil

	case Sub:
		configYAML, err := os.ReadFile("config.yaml")
		if err != nil {
			return nil, fmt.Errorf("error while reading config file, %w", err)
		}
		config := &Config{}

		err = yaml.Unmarshal(configYAML, &config)
		if err != nil {
			return nil, fmt.Errorf("error while reading config file, %w", err)
		}
		return config, nil

	default:
		return nil, fmt.Errorf("error storage type")
	}
}

func getAppType(appType string) AppType {

	if appType == "pub" {
		return Pub
	}
	if appType == "sub" {
		return Sub
	}
	return Other
}
