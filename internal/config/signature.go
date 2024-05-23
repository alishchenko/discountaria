package config

type Signature struct {
	N    int    `yaml:"n" env-required:"true"`
	R    int    `yaml:"r" env-required:"true"`
	P    int    `yaml:"p" env-required:"true"`
	Salt string `yaml:"salt" env-required:"true"`
}
