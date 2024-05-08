package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	ConfigPath           string
	DB                   DB                `yaml:"db" env-required:"true"`
	Env                  Env               `yaml:"env" env-required:"true"`
	HTTPServer           HTTPServer        `yaml:"http_server" env-required:"true"`
	Tokens               Tokens            `yaml:"tokens" env-required:"true"`
	OAuth2FacebookConfig OAuth2Config      `yaml:"oauth2_facebook" env-required:"true"`
	OAuth2GoogleConfig   OAuth2Config      `yaml:"oauth2_google" env-required:"true"`
	OAuth2LinkedinConfig OAuth2Config      `yaml:"oauth2_linkedin" env-required:"true"`
	OAuth2StateConfig    OAuth2StateConfig `yaml:"oauth2_state" env-required:"true"`

	AWSConfig AWSConfig `yaml:"aws" env-required:"true"`
	MimeTypes MimeTypes `yaml:"mime_types" env-required:"true"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	cfg.ConfigPath = configPath
	return cfg
}
