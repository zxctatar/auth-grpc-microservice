package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC   GRPCConfig   `yaml:"grpc"`
	Logger LoggerConfig `yaml:"logger"`
}

type GRPCConfig struct {
	Port         int    `yaml:"port"`
	JWTSecretKey string `yaml:"jwtSecretKey"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()

	if configPath == "" {
		panic("empty config path")
	}

	if _, err := os.Stat(configPath); err != nil {
		panic("config file does not exist")
	}

	data, err := os.ReadFile(configPath)

	if err != nil {
		panic("failed to read config file")
	}

	var cfg *Config

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		panic("failed to parse config file")
	}

	return cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "input config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}