package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Addr       string `env:"URLSHORTENER_ADDR"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASS"`
	DBAddr     string `env:"DB_ADDR"`
	DBName     string `env:"DB_NAME"`
}

func Get() (config Config, err error) {
	err = envconfig.Process("", &config)

	return config, err
}
