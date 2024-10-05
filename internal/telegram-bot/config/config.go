package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	PostgreSQL PostgreSQL `yaml:"postgres" env-required:"true"`
}

type PostgreSQL struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}
func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == ""  {
		log.Fatal("CFG: CONFIG_PATH environment variable is not set")
		return nil
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CFG: CONFIG_PATH \"%s\" does not exist", configPath)
		return nil
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("CFG: unable to load config file %s: %v", configPath, err)
		return nil
	}

	return &config
}
