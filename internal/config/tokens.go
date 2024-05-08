package config

import (
	"time"
)

type Tokens struct {
	AccessExp  time.Duration `yaml:"access_exp" env-default:"2h"`
	RefreshExp time.Duration `yaml:"refresh_exp" env-default:"10h"`
	SecretKey  string        `yaml:"secret_key" env-required:"true"`
}
