package config

import (
	"flag"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC     GRPCConfig     `yaml:"grpc"`
	Logger   LoggerConfig   `yaml:"logger"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type GRPCConfig struct {
	Port         int           `yaml:"port"`
	TimeOut      time.Duration `yaml:"timeout"`
	JWTSecretKey string        `yaml:"jwtSecretKey"`
	JWTTimeOut   time.Duration `yaml:"jwtTimeOut"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	DbName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"sslmode"`
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
