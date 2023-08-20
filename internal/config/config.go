package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:9129"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
	//unsafe
	User     string `yaml:"user" env-required:"true" env:"HTTP_SERVER_USER"`
	Password string `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() Config {
	configPath, err := filepath.Abs("../../config/conf.yaml")
	if err != nil {
		log.Fatalf("config path is not set: %s", err)
	}

	// check config file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cant read config: %s", err)
	}

	return config
}
