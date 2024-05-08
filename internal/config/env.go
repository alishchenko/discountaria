package config

type Env struct {
	Level string `yaml:"level" env-required:"true"`
}
