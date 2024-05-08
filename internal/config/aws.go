package config

import "time"

type AWSConfig struct {
	Endpoint       string        `yaml:"endpoint"`
	AccessKeyID    string        `yaml:"access_key" env-required:"true"`
	SecretKeyID    string        `yaml:"secret_key" env-required:"true"`
	Bucket         string        `yaml:"bucket" env-required:"true"`
	Expiration     time.Duration `yaml:"expiration" env-required:"true"`
	SslDisable     bool          `yaml:"ssldisable" env-required:"true"`
	ForcePathStyle bool          `yaml:"force_path_style" env-required:"true"`
	Region         string        `yaml:"region" env-required:"true"`
}

type MimeTypes struct {
	AllowedMimeTypes []string `yaml:"allowed_mime_types" env-required:"true"`
}
